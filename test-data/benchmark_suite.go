package benchmark

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

// BenchmarkSuite provides comprehensive performance testing capabilities
// This file tests the GitHub createFile tool with another large file

// Configuration constants
const (
	DefaultWorkerCount     = 10
	DefaultQueueSize       = 1000
	DefaultRetryAttempts   = 3
	DefaultRetryDelay      = time.Second
	MaxConcurrentRequests  = 100
	RequestTimeout         = 30 * time.Second
	CacheExpirationTime    = 10 * time.Minute
	MaxPayloadSize         = 1024 * 1024 * 5 // 5MB
	CompressionThreshold   = 1024            // 1KB
)

// MetricType represents different types of metrics
type MetricType string

const (
	MetricTypeCounter   MetricType = "counter"
	MetricTypeGauge     MetricType = "gauge"
	MetricTypeHistogram MetricType = "histogram"
	MetricTypeSummary   MetricType = "summary"
)

// Metric represents a performance metric
type Metric struct {
	Name      string                 `json:"name"`
	Type      MetricType             `json:"type"`
	Value     float64                `json:"value"`
	Timestamp time.Time              `json:"timestamp"`
	Labels    map[string]string      `json:"labels"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// BenchmarkResult stores the results of a benchmark run
type BenchmarkResult struct {
	TestName       string        `json:"test_name"`
	Duration       time.Duration `json:"duration"`
	Iterations     int64         `json:"iterations"`
	BytesProcessed int64         `json:"bytes_processed"`
	ErrorCount     int           `json:"error_count"`
	SuccessRate    float64       `json:"success_rate"`
	Metrics        []Metric      `json:"metrics"`
	StartTime      time.Time     `json:"start_time"`
	EndTime        time.Time     `json:"end_time"`
}

// WorkerPool manages a pool of workers for concurrent processing
type WorkerPool struct {
	workers    int
	jobQueue   chan Job
	resultChan chan Result
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	mu         sync.RWMutex
	stats      WorkerStats
}

// Job represents a unit of work
type Job struct {
	ID       string
	Payload  interface{}
	Priority int
	Retry    int
	Created  time.Time
}

// Result represents the result of a job
type Result struct {
	JobID     string
	Success   bool
	Data      interface{}
	Error     error
	Duration  time.Duration
	Timestamp time.Time
}

// WorkerStats tracks worker pool statistics
type WorkerStats struct {
	TotalJobs      int64
	CompletedJobs  int64
	FailedJobs     int64
	AverageLatency time.Duration
	LastUpdate     time.Time
}

// NewWorkerPool creates a new worker pool
func NewWorkerPool(workers, queueSize int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workers:    workers,
		jobQueue:   make(chan Job, queueSize),
		resultChan: make(chan Result, queueSize),
		ctx:        ctx,
		cancel:     cancel,
		stats:      WorkerStats{LastUpdate: time.Now()},
	}
}

// Start initializes and starts the worker pool
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

// worker processes jobs from the queue
func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	
	for {
		select {
		case <-wp.ctx.Done():
			return
		case job := <-wp.jobQueue:
			start := time.Now()
			result := wp.processJob(job)
			result.Duration = time.Since(start)
			
			wp.mu.Lock()
			wp.stats.CompletedJobs++
			if !result.Success {
				wp.stats.FailedJobs++
			}
			wp.stats.LastUpdate = time.Now()
			wp.mu.Unlock()
			
			select {
			case wp.resultChan <- result:
			case <-wp.ctx.Done():
				return
			}
		}
	}
}

// processJob processes a single job
func (wp *WorkerPool) processJob(job Job) Result {
	// Simulate job processing
	time.Sleep(time.Millisecond * time.Duration(10+job.Priority))
	
	return Result{
		JobID:     job.ID,
		Success:   true,
		Data:      fmt.Sprintf("Processed job %s", job.ID),
		Timestamp: time.Now(),
	}
}

// Submit adds a job to the queue
func (wp *WorkerPool) Submit(job Job) error {
	wp.mu.Lock()
	wp.stats.TotalJobs++
	wp.mu.Unlock()
	
	select {
	case wp.jobQueue <- job:
		return nil
	case <-wp.ctx.Done():
		return fmt.Errorf("worker pool is shutting down")
	default:
		return fmt.Errorf("job queue is full")
	}
}

// Stop gracefully shuts down the worker pool
func (wp *WorkerPool) Stop() {
	wp.cancel()
	close(wp.jobQueue)
	wp.wg.Wait()
	close(wp.resultChan)
}

// GetStats returns current worker pool statistics
func (wp *WorkerPool) GetStats() WorkerStats {
	wp.mu.RLock()
	defer wp.mu.RUnlock()
	return wp.stats
}

// HTTPClient provides HTTP request capabilities with retry logic
type HTTPClient struct {
	client       *http.Client
	retryCount   int
	retryDelay   time.Duration
	timeout      time.Duration
	interceptors []Interceptor
}

// Interceptor allows request/response modification
type Interceptor func(*http.Request) error

// NewHTTPClient creates a new HTTP client
func NewHTTPClient(timeout time.Duration, retryCount int) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		retryCount: retryCount,
		retryDelay: DefaultRetryDelay,
		timeout:    timeout,
	}
}

// AddInterceptor adds a request interceptor
func (hc *HTTPClient) AddInterceptor(interceptor Interceptor) {
	hc.interceptors = append(hc.interceptors, interceptor)
}

// Do executes an HTTP request with retry logic
func (hc *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	for _, interceptor := range hc.interceptors {
		if err := interceptor(req); err != nil {
			return nil, err
		}
	}
	
	var resp *http.Response
	var err error
	
	for attempt := 0; attempt <= hc.retryCount; attempt++ {
		resp, err = hc.client.Do(req)
		if err == nil && resp.StatusCode < 500 {
			return resp, nil
		}
		
		if attempt < hc.retryCount {
			time.Sleep(hc.retryDelay * time.Duration(attempt+1))
		}
	}
	
	return resp, err
}

// DataStructures for testing various scenarios

// TreeNode represents a node in a binary tree
type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BinaryTree provides binary tree operations
type BinaryTree struct {
	Root *TreeNode
	Size int
}

// Insert adds a value to the binary tree
func (bt *BinaryTree) Insert(value int) {
	if bt.Root == nil {
		bt.Root = &TreeNode{Value: value}
		bt.Size++
		return
	}
	bt.insertNode(bt.Root, value)
}

// insertNode recursively inserts a node
func (bt *BinaryTree) insertNode(node *TreeNode, value int) {
	if value < node.Value {
		if node.Left == nil {
			node.Left = &TreeNode{Value: value}
			bt.Size++
		} else {
			bt.insertNode(node.Left, value)
		}
	} else {
		if node.Right == nil {
			node.Right = &TreeNode{Value: value}
			bt.Size++
		} else {
			bt.insertNode(node.Right, value)
		}
	}
}

// InOrderTraversal performs in-order traversal
func (bt *BinaryTree) InOrderTraversal() []int {
	result := []int{}
	bt.inOrder(bt.Root, &result)
	return result
}

// inOrder helper for traversal
func (bt *BinaryTree) inOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	bt.inOrder(node.Left, result)
	*result = append(*result, node.Value)
	bt.inOrder(node.Right, result)
}

// Graph represents a directed graph
type Graph struct {
	vertices map[string]*Vertex
	mu       sync.RWMutex
}

// Vertex represents a graph vertex
type Vertex struct {
	ID    string
	Edges map[string]int // destination -> weight
	Data  interface{}
}

// NewGraph creates a new graph
func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[string]*Vertex),
	}
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(id string, data interface{}) {
	g.mu.Lock()
	defer g.mu.Unlock()
	
	if _, exists := g.vertices[id]; !exists {
		g.vertices[id] = &Vertex{
			ID:    id,
			Edges: make(map[string]int),
			Data:  data,
		}
	}
}

// AddEdge adds an edge between two vertices
func (g *Graph) AddEdge(from, to string, weight int) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	
	fromVertex, exists := g.vertices[from]
	if !exists {
		return fmt.Errorf("vertex %s not found", from)
	}
	
	if _, exists := g.vertices[to]; !exists {
		return fmt.Errorf("vertex %s not found", to)
	}
	
	fromVertex.Edges[to] = weight
	return nil
}

// ShortestPath finds the shortest path using Dijkstra's algorithm
func (g *Graph) ShortestPath(start, end string) ([]string, int) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	
	distances := make(map[string]int)
	previous := make(map[string]string)
	unvisited := make(map[string]bool)
	
	for id := range g.vertices {
		distances[id] = math.MaxInt32
		unvisited[id] = true
	}
	distances[start] = 0
	
	for len(unvisited) > 0 {
		current := g.findMinDistance(distances, unvisited)
		if current == "" || current == end {
			break
		}
		
		delete(unvisited, current)
		
		for neighbor, weight := range g.vertices[current].Edges {
			if !unvisited[neighbor] {
				continue
			}
			
			alt := distances[current] + weight
			if alt < distances[neighbor] {
				distances[neighbor] = alt
				previous[neighbor] = current
			}
		}
	}
	
	path := g.reconstructPath(previous, start, end)
	return path, distances[end]
}

// findMinDistance finds the unvisited vertex with minimum distance
func (g *Graph) findMinDistance(distances map[string]int, unvisited map[string]bool) string {
	minDist := math.MaxInt32
	minVertex := ""
	
	for vertex := range unvisited {
		if distances[vertex] < minDist {
			minDist = distances[vertex]
			minVertex = vertex
		}
	}
	
	return minVertex
}

// reconstructPath reconstructs the path from start to end
func (g *Graph) reconstructPath(previous map[string]string, start, end string) []string {
	path := []string{}
	current := end
	
	for current != "" {
		path = append([]string{current}, path...)
		if current == start {
			break
		}
		current = previous[current]
	}
	
	return path
}

// CryptoUtils provides cryptographic utilities
type CryptoUtils struct{}

// GenerateRandomBytes generates random bytes
func (cu *CryptoUtils) GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString generates a random string
func (cu *CryptoUtils) GenerateRandomString(length int) (string, error) {
	bytes, err := cu.GenerateRandomBytes(length)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes)[:length], nil
}

// HashSHA256 computes SHA256 hash
func (cu *CryptoUtils) HashSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}

// GenerateSecureToken generates a secure random token
func (cu *CryptoUtils) GenerateSecureToken(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[num.Int64()]
	}
	
	return string(b), nil
}

// StatisticsCalculator provides statistical calculations
type StatisticsCalculator struct{}

// Mean calculates the mean of a dataset
func (sc *StatisticsCalculator) Mean(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}

// Median calculates the median of a dataset
func (sc *StatisticsCalculator) Median(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	
	sorted := make([]float64, len(data))
	copy(sorted, data)
	sort.Float64s(sorted)
	
	n := len(sorted)
	if n%2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2
	}
	return sorted[n/2]
}

// StandardDeviation calculates the standard deviation
func (sc *StatisticsCalculator) StandardDeviation(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	
	mean := sc.Mean(data)
	variance := 0.0
	
	for _, v := range data {
		diff := v - mean
		variance += diff * diff
	}
	
	variance /= float64(len(data))
	return math.Sqrt(variance)
}

// Percentile calculates the nth percentile
func (sc *StatisticsCalculator) Percentile(data []float64, p float64) float64 {
	if len(data) == 0 {
		return 0
	}
	
	sorted := make([]float64, len(data))
	copy(sorted, data)
	sort.Float64s(sorted)
	
	index := (p / 100.0) * float64(len(sorted)-1)
	lower := int(math.Floor(index))
	upper := int(math.Ceil(index))
	
	if lower == upper {
		return sorted[lower]
	}
	
	weight := index - float64(lower)
	return sorted[lower]*(1-weight) + sorted[upper]*weight
}

// JSONProcessor handles JSON operations
type JSONProcessor struct{}

// Marshal converts data to JSON
func (jp *JSONProcessor) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalIndent converts data to indented JSON
func (jp *JSONProcessor) MarshalIndent(v interface{}) ([]byte, error) {
	return json.MarshalIndent(v, "", "  ")
}

// Unmarshal parses JSON data
func (jp *JSONProcessor) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// ValidateJSON checks if data is valid JSON
func (jp *JSONProcessor) ValidateJSON(data []byte) bool {
	var js interface{}
	return json.Unmarshal(data, &js) == nil
}

// StringUtils provides string manipulation utilities
type StringUtils struct{}

// Reverse reverses a string
func (su *StringUtils) Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome checks if a string is a palindrome
func (su *StringUtils) IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return s == su.Reverse(s)
}

// CountWords counts words in a string
func (su *StringUtils) CountWords(s string) int {
	return len(strings.Fields(s))
}

// Truncate truncates a string to specified length
func (su *StringUtils) Truncate(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[:length] + "..."
}

// BenchmarkRunner executes benchmark tests
type BenchmarkRunner struct {
	results []BenchmarkResult
	mu      sync.Mutex
}

// NewBenchmarkRunner creates a new benchmark runner
func NewBenchmarkRunner() *BenchmarkRunner {
	return &BenchmarkRunner{
		results: make([]BenchmarkResult, 0),
	}
}

// RunBenchmark executes a benchmark test
func (br *BenchmarkRunner) RunBenchmark(name string, fn func() error, iterations int64) BenchmarkResult {
	result := BenchmarkResult{
		TestName:   name,
		Iterations: iterations,
		StartTime:  time.Now(),
		Metrics:    make([]Metric, 0),
	}
	
	for i := int64(0); i < iterations; i++ {
		if err := fn(); err != nil {
			result.ErrorCount++
		}
	}
	
	result.EndTime = time.Now()
	result.Duration = result.EndTime.Sub(result.StartTime)
	result.SuccessRate = float64(iterations-int64(result.ErrorCount)) / float64(iterations) * 100
	
	br.mu.Lock()
	br.results = append(br.results, result)
	br.mu.Unlock()
	
	return result
}

// GetResults returns all benchmark results
func (br *BenchmarkRunner) GetResults() []BenchmarkResult {
	br.mu.Lock()
	defer br.mu.Unlock()
	return br.results
}

// Main benchmark execution function
func RunAllBenchmarks() {
	fmt.Println("Starting comprehensive benchmark suite...")
	
	runner := NewBenchmarkRunner()
	
	// Benchmark 1: Worker Pool
	fmt.Println("Benchmark 1: Worker Pool Performance")
	pool := NewWorkerPool(10, 100)
	pool.Start()
	
	result1 := runner.RunBenchmark("WorkerPool", func() error {
		job := Job{
			ID:       fmt.Sprintf("job_%d", time.Now().UnixNano()),
			Priority: 1,
			Created:  time.Now(),
		}
		return pool.Submit(job)
	}, 1000)
	
	pool.Stop()
	fmt.Printf("Result: %+v\n", result1)
	
	// Benchmark 2: Binary Tree Operations
	fmt.Println("Benchmark 2: Binary Tree Operations")
	tree := &BinaryTree{}
	
	result2 := runner.RunBenchmark("BinaryTree", func() error {
		tree.Insert(int(time.Now().UnixNano() % 1000))
		return nil
	}, 500)
	
	fmt.Printf("Result: %+v\n", result2)
	
	// Benchmark 3: Cryptographic Operations
	fmt.Println("Benchmark 3: Cryptographic Operations")
	crypto := &CryptoUtils{}
	
	result3 := runner.RunBenchmark("Crypto", func() error {
		_, err := crypto.GenerateSecureToken(32)
		return err
	}, 1000)
	
	fmt.Printf("Result: %+v\n", result3)
	
	fmt.Println("All benchmarks completed!")
}

// End of benchmark file
