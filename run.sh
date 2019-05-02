#!/bin/bash

# Stop all container
# docker rm $(docker ps -a -q) -f

# Start platform
docker-compose up

# Injection de messages dans kafka
docker exec -it $(docker-compose ps -q kafka) kafka-console-producer.sh --broker-list localhost:9092 --topic events
2016-07-11T23:56:42.000+00:00 INFO [MySecretApp.com.Transaction.Manager]:Starting transaction for session -464410bf-37bf-475a-afc0-498e0199f008

Column 1 = "May 11 10:40:48" > Timestamp
Column 2 = "scrooge" > Loghost
Column 3 = "disk-health-nurse[26783]:" > Application/Process
Column 4 = "[ID 702911 user.error]" > Syslog facility.level
Column 5 = "m:SY-mon-full-500" > Message ID
Column 6 = "c:H : partition health..." > Message [possibly including rid, sid, ip]

# Lecture des messages dans kafka
docker exec -it $(docker-compose ps -q kafka) kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic error --from-beginning

# List topics in Kafka
docker exec -it $(docker-compose ps -q kafka) kafka-topics.sh --zookeeper zookeeper:2181 --list

# List topic configuration in Kafka
 docker exec -it $(docker-compose ps -q kafka) kafka-topics.sh --zookeeper zookeeper:2181  --describe

## new solution
docker-compose -f platform.yml up
go run app/main.go

## proxy conf docker
mkdir /etc/systemd/system/docker.service.d && vim /etc/systemd/system/docker.service.d/https-proxy.conf
Environment="HTTPS_PROXY=https://127.0.0.1:3128"

## proxy conf go get
https_proxy=127.0.0.1:3128 go get github.com/Sirupsen/logrus
https_proxy=127.0.0.1:3128 go get github.com/logrusorgru/grokky
https_proxy=127.0.0.1:3128 go get github.com/prometheus/client_golang/prometheus
https_proxy=127.0.0.1:3128 go get github.com/prometheus/client_golang/prometheus/promhttp
https_proxy=127.0.0.1:3128 go get github.com/segmentio/kafka-go

66.249.65.159 - [06/Nov/2014:19:10:38 +0600] GET /news/53f8d72920ba2744fe873ebc.html HTTP/1.1 404 177 - Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X)AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1;http://www.google.com/bot.html)



# Injection de messages dans kafka
docker exec -it $(docker-compose ps -q kafka) kafka-console-producer --broker-list localhost:9092 --topic events
2016-07-11T23:56:42.000+00:00 INFO [MySecretApp.com.Transaction.Manager]:Starting transaction for session -464410bf-37bf-475a-afc0-498e0199f008


# Create Ktable
docker exec -it $(docker-compose ps -q ksql-server)

CREATE TABLE mapping
  (ipsource VARCHAR,
   familly VARCHAR)
  WITH (KAFKA_TOPIC = 'mapping',
        VALUE_FORMAT='JSON',
        KEY = 'ipsource');