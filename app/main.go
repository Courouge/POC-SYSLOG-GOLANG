package main

import (
    "fmt"
    "github.com/segmentio/kafka-go"
    "context"
    "net/http"

    "github.com/logrusorgru/grokky"

    log "github.com/Sirupsen/logrus"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

//Define the metrics we wish to expose
var    currentOffsetApp = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "current_offset_app", Help: "current offset app"})

func createHost() grokky.Host {
	h := grokky.New()
	// add patterns to the Host
	h.Must("YEAR", `(?:(?:19|20)[0-9]{2})`)
	h.Must("MONTHNUM", `(?:0?[1-9]|1[0-2])`)
	h.Must("MONTHDAY", `(?:(?:0[1-9])|(?:[12][0-9])|(?:3[01])|[1-9])`)
	h.Must("HOUR", `(?:2[0123]|[01]?[0-9])`)
	h.Must("MINUTE", `(?:[0-5][0-9])`)
	h.Must("SECOND", `(?:(?:[0-5]?[0-9]|60)(?:[:.,][0-9]+)?)`)
    h.Must("DATA", `\[.+.]`)
    h.Must("GREEDYDATA", `.*`)
    h.Must("IPV4", `(?:(?:[0-1]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])[.](?:[0-1]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])[.](?:[0-1]?[0-9]{1,2}|2[0-4][0-9]|25[0-5])[.](?:[0-1]?[0-9]{1,2}|2[0-4][0-9]|25[0-5]))`)
    h.Must("ISO8601_TIMEZONE", `(?:Z|[+-]%{HOUR}(?::?%{MINUTE}))`)
    h.Must("LOGLEVEL", `([Ii]nfo|INFO|[Tt]race|TRACE|[Dd]ebug|DEBUG|WARN?(?:ING)?)`)
	h.Must("TIMESTAMP_ISO8601", `%{YEAR}-%{MONTHNUM}-%{MONTHDAY}[T ]%{HOUR}:?%{MINUTE}(?::?%{SECOND})?%{ISO8601_TIMEZONE}`)


	return h
}

// 10.252.21.10 2016-07-11T23:56:42.000+00:00 INFO [MySecretApp.com.Transaction.Manager]:Starting transaction for session -464410bf-37bf-475a-afc0-498e0199f008


func main() {
    //go prometheus_exporter()
    consumer()
}

func consumer() {
// make a new reader that consumes from topic-A, partition 0, at offset 42
r := kafka.NewReader(kafka.ReaderConfig{
    Brokers:   []string{"127.0.0.1:9092"},
    Topic:     "events",
    GroupID:   "myappoffset",
    MinBytes:  10e3, // 10KB
    MaxBytes:  10e6, // 10MB

})


h := createHost()
// compile the pattern for RFC3339 time
p, err := h.Compile("%{IPV4:ipv4} %{TIMESTAMP_ISO8601:time} %{LOGLEVEL:loglevel} %{DATA:class}:%{GREEDYDATA:msg}")
if err != nil {
    log.Fatal(err)
}

for {
    m, err := r.ReadMessage(context.Background())
    if err != nil {
        fmt.Printf("Kafka server is not connected")
        break
    }

    currentOffsetApp.Set(float64(m.Offset))

    if len(string(m.Value)) > 0 {
        for k, v := range p.Parse(string(m.Value)) {
            if k == "ipv4" {
            producer(v)
            fmt.Printf("%s: %v\n", k, v)
            }
            if k == "msg" {
            producer(v)
            fmt.Printf("%s: %v\n", k, v)
            }
        }
        // producer(string(m.Value))
    }
}

r.Close()

}

func producer(msg string) {
topic := "error"
partition := 0
conn, _ := kafka.DialLeader(context.Background(), "tcp", "127.0.0.1:9092", topic, partition)
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

	//Register metrics with prometheus
	prometheus.MustRegister(currentOffsetApp)

}