package main

import (
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"time"
)

func main()  {
	c := metrics.NewCounter()
	metrics.Register("foo", c)

	go metrics.Log(metrics.DefaultRegistry,time.Second,log.New(os.Stdout,
		"metrics: ",log.Lmicroseconds))
	time.Sleep(time.Second)

	for {
		c.Inc(2)
		time.Sleep(time.Second * 1)
	}
}