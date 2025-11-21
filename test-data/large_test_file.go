package testdata

import (
	"fmt"
	"math/rand"
	"time"
)

// LargeTestFile is a comprehensive test file to validate GitHub createFile functionality
// This file contains extensive code, comments, and data structures to test large file handling

// Constants for testing
const (
	MaxIterations = 10000
	BufferSize    = 8192
	DefaultTimeout = 30 * time.Second
	APIVersion    = "v1.0.0"
	ServiceName   = "TestService"
)

// User represents a user entity in the system
type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsActive  bool      `json:"is_active"`
	Roles     []string  `json:"roles"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// Product represents a product in the inventory
type Product struct {
	ID          int64   `json:"id"`
	SKU         string  `json:"sku"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category    string  `json:"category"`
	Tags        []string `json:"tags"`
}

// Order represents a customer order
type Order struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Products    []Product `json:"products"`
	TotalAmount float64   `json:"total_amount"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at,omitempty"`
}

// DataProcessor handles large-scale data processing operations
type DataProcessor struct {
	batchSize int
	workers   int
	timeout   time.Duration
}

// NewDataProcessor creates a new DataProcessor instance
func NewDataProcessor(batchSize, workers int, timeout time.Duration) *DataProcessor {
	return &DataProcessor{
		batchSize: batchSize,
		workers:   workers,
		timeout:   timeout,
	}
}

// ProcessBatch processes a batch of data items
func (dp *DataProcessor) ProcessBatch(items []interface{}) error {
	fmt.Printf("Processing batch of %d items with %d workers\n", len(items), dp.workers)
	
	// Simulate processing
	for i, item := range items {
		if i%100 == 0 {
			fmt.Printf("Processed %d/%d items\n", i, len(items))
		}
		_ = item // Process item
	}
	
	return nil
}

// GenerateLargeDataset creates a large dataset for testing
func GenerateLargeDataset(size int) []User {
	users := make([]User, size)
	rand.Seed(time.Now().UnixNano())
	
	for i := 0; i < size; i++ {
		users[i] = User{
			ID:        int64(i + 1),
			Username:  fmt.Sprintf("user_%d", i+1),
			Email:     fmt.Sprintf("user%d@example.com", i+1),
			FirstName: fmt.Sprintf("FirstName%d", i+1),
			LastName:  fmt.Sprintf("LastName%d", i+1),
			CreatedAt: time.Now().Add(-time.Duration(rand.Intn(365*24)) * time.Hour),
			UpdatedAt: time.Now(),
			IsActive:  rand.Intn(2) == 1,
			Roles:     []string{"user", "member"},
			Metadata:  map[string]interface{}{"level": rand.Intn(10)},
		}
	}
	
	return users
}

// PerformanceTest runs various performance tests
func PerformanceTest() {
	fmt.Println("Starting performance tests...")
	
	// Test 1: Memory allocation
	fmt.Println("Test 1: Memory allocation test")
	data := make([]byte, 1024*1024*10) // 10MB
	for i := range data {
		data[i] = byte(i % 256)
	}
	
	// Test 2: CPU intensive operations
	fmt.Println("Test 2: CPU intensive operations")
	result := 0
	for i := 0; i < MaxIterations; i++ {
		result += i * i
	}
	fmt.Printf("Result: %d\n", result)
	
	// Test 3: String operations
	fmt.Println("Test 3: String operations")
	var builder strings.Builder
	for i := 0; i < 1000; i++ {
		builder.WriteString(fmt.Sprintf("Line %d: This is a test string with some content\n", i))
	}
	
	// Test 4: Map operations
	fmt.Println("Test 4: Map operations")
	testMap := make(map[string]int)
	for i := 0; i < 10000; i++ {
		testMap[fmt.Sprintf("key_%d", i)] = i
	}
}

// ComplexAlgorithm demonstrates a complex algorithm implementation
func ComplexAlgorithm(input []int) []int {
	if len(input) == 0 {
		return []int{}
	}
	
	// Bubble sort for demonstration
	result := make([]int, len(input))
	copy(result, input)
	
	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	
	return result
}

// DataValidator validates data structures
type DataValidator struct {
	rules map[string]func(interface{}) bool
}

// NewDataValidator creates a new validator
func NewDataValidator() *DataValidator {
	return &DataValidator{
		rules: make(map[string]func(interface{}) bool),
	}
}

// AddRule adds a validation rule
func (dv *DataValidator) AddRule(name string, rule func(interface{}) bool) {
	dv.rules[name] = rule
}

// Validate validates data against all rules
func (dv *DataValidator) Validate(data interface{}) bool {
	for name, rule := range dv.rules {
		if !rule(data) {
			fmt.Printf("Validation failed for rule: %s\n", name)
			return false
		}
	}
	return true
}

// CacheManager manages in-memory cache
type CacheManager struct {
	cache map[string]interface{}
	ttl   time.Duration
}

// NewCacheManager creates a new cache manager
func NewCacheManager(ttl time.Duration) *CacheManager {
	return &CacheManager{
		cache: make(map[string]interface{}),
		ttl:   ttl,
	}
}

// Set adds an item to cache
func (cm *CacheManager) Set(key string, value interface{}) {
	cm.cache[key] = value
}

// Get retrieves an item from cache
func (cm *CacheManager) Get(key string) (interface{}, bool) {
	value, exists := cm.cache[key]
	return value, exists
}

// Clear clears the cache
func (cm *CacheManager) Clear() {
	cm.cache = make(map[string]interface{})
}

// Additional test data and functions to increase file size
var (
	TestData1 = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit",
		"Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua",
		"Ut enim ad minim veniam, quis nostrud exercitation ullamco",
		"Duis aute irure dolor in reprehenderit in voluptate velit",
		"Excepteur sint occaecat cupidatat non proident, sunt in culpa",
	}
	
	TestData2 = map[string]int{
		"alpha": 1, "beta": 2, "gamma": 3, "delta": 4, "epsilon": 5,
		"zeta": 6, "eta": 7, "theta": 8, "iota": 9, "kappa": 10,
	}
	
	TestData3 = []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987}
)

// HelperFunction1 demonstrates helper functionality
func HelperFunction1(a, b int) int {
	return a + b
}

// HelperFunction2 demonstrates more helper functionality
func HelperFunction2(s string) string {
	return fmt.Sprintf("Processed: %s", s)
}

// HelperFunction3 demonstrates complex helper functionality
func HelperFunction3(data []interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i, item := range data {
		result[fmt.Sprintf("item_%d", i)] = item
	}
	return result
}

// Main function for testing
func main() {
	fmt.Println("Large Test File - GitHub CreateFile Tool Test")
	fmt.Println("==============================================")
	
	// Initialize components
	processor := NewDataProcessor(100, 4, DefaultTimeout)
	validator := NewDataValidator()
	cache := NewCacheManager(5 * time.Minute)
	
	// Generate test data
	users := GenerateLargeDataset(1000)
	fmt.Printf("Generated %d users\n", len(users))
	
	// Process data
	items := make([]interface{}, len(users))
	for i, user := range users {
		items[i] = user
	}
	
	if err := processor.ProcessBatch(items); err != nil {
		fmt.Printf("Error processing batch: %v\n", err)
	}
	
	// Run performance tests
	PerformanceTest()
	
	// Test complex algorithm
	testInput := []int{64, 34, 25, 12, 22, 11, 90}
	sorted := ComplexAlgorithm(testInput)
	fmt.Printf("Sorted: %v\n", sorted)
	
	// Cache operations
	cache.Set("test_key", "test_value")
	if value, exists := cache.Get("test_key"); exists {
		fmt.Printf("Cache hit: %v\n", value)
	}
	
	fmt.Println("Test completed successfully!")
}

// Additional structures and functions to reach larger file size

// Logger interface for logging operations
type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Warn(msg string)
}

// ConsoleLogger implements Logger for console output
type ConsoleLogger struct {
	prefix string
}

// NewConsoleLogger creates a new console logger
func NewConsoleLogger(prefix string) *ConsoleLogger {
	return &ConsoleLogger{prefix: prefix}
}

// Info logs info message
func (cl *ConsoleLogger) Info(msg string) {
	fmt.Printf("[INFO] %s: %s\n", cl.prefix, msg)
}

// Error logs error message
func (cl *ConsoleLogger) Error(msg string) {
	fmt.Printf("[ERROR] %s: %s\n", cl.prefix, msg)
}

// Debug logs debug message
func (cl *ConsoleLogger) Debug(msg string) {
	fmt.Printf("[DEBUG] %s: %s\n", cl.prefix, msg)
}

// Warn logs warning message
func (cl *ConsoleLogger) Warn(msg string) {
	fmt.Printf("[WARN] %s: %s\n", cl.prefix, msg)
}

// Repository pattern implementation
type Repository interface {
	Create(entity interface{}) error
	Read(id int64) (interface{}, error)
	Update(entity interface{}) error
	Delete(id int64) error
	List() ([]interface{}, error)
}

// UserRepository implements Repository for User entities
type UserRepository struct {
	data map[int64]User
}

// NewUserRepository creates a new user repository
func NewUserRepository() *UserRepository {
	return &UserRepository{
		data: make(map[int64]User),
	}
}

// Create adds a new user
func (ur *UserRepository) Create(entity interface{}) error {
	user, ok := entity.(User)
	if !ok {
		return fmt.Errorf("invalid entity type")
	}
	ur.data[user.ID] = user
	return nil
}

// Read retrieves a user by ID
func (ur *UserRepository) Read(id int64) (interface{}, error) {
	user, exists := ur.data[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// Update updates an existing user
func (ur *UserRepository) Update(entity interface{}) error {
	user, ok := entity.(User)
	if !ok {
		return fmt.Errorf("invalid entity type")
	}
	ur.data[user.ID] = user
	return nil
}

// Delete removes a user
func (ur *UserRepository) Delete(id int64) error {
	delete(ur.data, id)
	return nil
}

// List returns all users
func (ur *UserRepository) List() ([]interface{}, error) {
	result := make([]interface{}, 0, len(ur.data))
	for _, user := range ur.data {
		result = append(result, user)
	}
	return result, nil
}

// End of large test file
