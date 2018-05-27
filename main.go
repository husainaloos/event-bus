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
	c := controllers.NewDemoController("controller1")
	p1 := publishers.NewDemoPublisher("publisher1")
	p2 := publishers.NewDemoPublisher("publisher2")
	f := filters.NewYesFilter()
	s := subscribers.NewDemoSubscriber("subscriber1")

	c.RegisterPublisher(p1)
	c.RegisterPublisher(p2)
	c.RegisterSubscriber(f, s)

	go c.Start()

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	c.Stop()
}
