package main

import (
	"bufio"
	"os"

	"github.com/husainaloos/event-bus/controllers"
	"github.com/husainaloos/event-bus/filters"
	"github.com/husainaloos/event-bus/publishers"
	"github.com/husainaloos/event-bus/subscribers"
)

func main() {
	c := controllers.NewDefaultController("controller1")
	p1 := publishers.NewTimedPublisher("publisher1")
	p2 := publishers.NewTimedPublisher("publisher2")
	f := filters.NewAlwaysAllowFilter()
	s := subscribers.NewWriterSubscriber("subscriber1", os.Stdout)

	c.RegisterPublisher(p1)
	c.RegisterPublisher(p2)
	c.RegisterSubscriber(f, s)

	go c.Start()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	c.Stop()
}
