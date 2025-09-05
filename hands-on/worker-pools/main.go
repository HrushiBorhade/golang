package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Job %d started by worker %d \n", j ,id)
		time.Sleep(time.Second)
		fmt.Printf("Job %d ended by worker %d \n", j, id)
		results <- j*2
	}
}


func main() {
	const numJobs = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w:=1; w<=3; w++ {
		go worker(w,jobs,results)
	}

	for j:=1; j<=numJobs; j++{
		jobs <- j
	}

	close(jobs)

	for a:=1; a<=numJobs; a++ {
		fmt.Println("result: ", <-results)
	}
}
