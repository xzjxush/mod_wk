package main

import (
	"github.com/rcrowley/go-metrics"
	"github.com/vrischmann/go-metrics-influxdb"
	"time"
)
// 查询DB时使用下面sql
// select * from "quux.meter"
// 因为带. 所以measurement上必须带双引号
func main(){
	m := metrics.NewMeter()
	metrics.Register("quux", m)
	m.Mark(1)

	go influxdb.InfluxDB(metrics.DefaultRegistry,
		time.Second * 5,
		"http://127.0.0.1:8086",
		"mydb",
		"",
		"")

	var j int64
	j = 1
	for true {
		time.Sleep(time.Second*1)
		m.Mark(j)
		j += 1
	}
}
