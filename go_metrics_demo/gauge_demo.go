package main

import (
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"time"
)

func main() {
	// 瞬时状态
	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(1)

	go metrics.Log(metrics.DefaultRegistry,time.Second,log.New(os.Stdout,
		"metrics: ",log.Lmicroseconds))

	var j int64
	for {
		time.Sleep(time.Second * 1)
		g.Update(j)
		j++
	}


}
