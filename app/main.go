package main

import (
    "fmt"
    "github.com/segmentio/kafka-go"
    "context"
    "strings"
    "net/http"

    log "github.com/Sirupsen/logrus"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    //fmt.Printf("start main function\n")
    go prometheus_exporter()
    consumer()
}

func consumer() {
// make a new reader that consumes from topic-A, partition 0, at offset 42
r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"172.17.0.1:9092"},
    Topic:     "events",
    GroupID:   "test",
    Partition: 0,
    MinBytes:  10e3, // 10KB
    MaxBytes:  10e6, // 10MB

})

for {

    m, err := r.ReadMessage(context.Background())
    if err != nil {
        fmt.Printf("Kafka server is not connected")
        break
    }
    //fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
    stringSlice := strings.Split(string(m.Value), "|")
    fmt.Printf("%s\n",stringSlice)
    producer(string(m.Value))

}

r.Close()

}

func producer(msg string) {
topic := "error"
partition := 0
conn, _ := kafka.DialLeader(context.Background(), "tcp", "172.17.0.1:9092", topic, partition)
    conn.WriteMessages(
        kafka.Message{Value: []byte(msg)},
    )

conn.Close()
}

func prometheus_exporter() {
  //This section will start the HTTP server and expose
  //any metrics on the /metrics endpoint.
  http.Handle("/metrics", promhttp.Handler())
  log.Info("Beginning to serve on port :8080")
  log.Fatal(http.ListenAndServe(":8080", nil))

}

func init(){
//Define the metrics we wish to expose
var fooMetric = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "foo_metric", Help: "current offset app"})

var barMetric = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "bar_metric", Help: "Shows whether a bar has occurred in our cluster"})

	//Register metrics with prometheus
	prometheus.MustRegister(fooMetric)
	prometheus.MustRegister(barMetric)

	//Set fooMetric to 1
	fooMetric.Set(0)

	//Set barMetric to 0
	barMetric.Set(1)
}