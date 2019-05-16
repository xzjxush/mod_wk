package main

import (
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"time"
)

func main(){

	m := metrics.NewMeter()
	metrics.Register("quux", m)
	m.Mark(1)


	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))


	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		j++
		m.Mark(j)
	}
}
