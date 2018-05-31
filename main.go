package main

import (
	"bufio"
	"os"

	"github.com/husainaloos/event-bus/controller"
	"github.com/husainaloos/event-bus/filter"
	"github.com/husainaloos/event-bus/publisher"
	"github.com/husainaloos/event-bus/subscriber"
)

func main() {
	c := controller.NewDefaultController("controller1")
	p1 := publisher.NewTimedPublisher("publisher1")
	p2 := publisher.NewTimedPublisher("publisher2")
	f := filter.NewAlwaysAllowFilter()
	s := subscriber.NewWriterSubscriber("subscriber1", os.Stdout)

	c.AddPublisher(p1)
	c.AddPublisher(p2)
	c.AddSubscriber(f, s)

	go c.Start()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	c.Stop()
}
