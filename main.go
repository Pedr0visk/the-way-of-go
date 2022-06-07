package main

import (
	"educative/adapters/queue"
	"fmt"
)

func main() {
	queue := make(chan *queue.Message)

	for i := 0; i < 10; i++ {
		queue <- queue.NewMessage("key1", fmt.Sprintf("message %v", i))
	}
	close(queue)

}
