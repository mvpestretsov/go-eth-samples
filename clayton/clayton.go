package main

import (
	"log"
)

type Pipe interface {
	Process(in chan int) chan int
}
type Sq struct{}

func (sq Sq) Process(in chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			log.Printf("squaring %d", i)
			out <- i * i
		}
		close(out)
	}()
	return out
}

type Add struct{}

func (add Add) Process(in chan int) chan int {
	out := make(chan int)
	go func() {
		for i := range in {
			log.Printf("multiply 2 %d", i)
			out <- i + i
		}
		close(out)
	}()
	return out
}

func NewPipeline(pipes ...Pipe) Pipeline {
	head := make(chan int)
	var next_chan chan int
	for _, pipe := range pipes {
		if next_chan == nil {
			next_chan = pipe.Process(head)
		} else {
			next_chan = pipe.Process(next_chan)
		}
	}
	return Pipeline{head: head, tail: next_chan}
}

type Pipeline struct {
	head chan int
	tail chan int
}

func (p *Pipeline) Enqueue(item int) {
	p.head <- item
}

func (p *Pipeline) Dequeue(handler func(int)) {
	for i := range p.tail {
		handler(i)
	}
}

func (p *Pipeline) Close() {
	close(p.head)
}

func main() {

	pipeline := NewPipeline(Sq{}, Add{})
	go func() {
		for i := 0; i < 3; i++ {
			log.Printf("Sending: %v", i)
			pipeline.Enqueue(i)
		}
		log.Println("Closing Pipeline.")
		pipeline.Close()
	}()
	pipeline.Dequeue(func(i int) {
		log.Printf("Received: %v", i)
	})
}
