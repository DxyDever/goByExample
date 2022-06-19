package main

import (
	"fmt"
	"time"
)

/*
工作池
*/

func worker1(id int, jobs <-chan int, result chan<- int) {

	for job := range jobs {
		fmt.Println("worker", id, "start job", job)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", jobs)
		result <- job * 2
	}

}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	result := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker1(w, jobs, result)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-result
	}
}
