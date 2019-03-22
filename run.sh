#!/bin/bash

# Stop all container
# docker rm $(docker ps -a -q) -f

# Start platform
docker-compose up

# Injection de messages dans kafka
docker exec -it $(docker-compose ps -q kafka) kafka-console-producer.sh --broker-list localhost:9092 --topic events
May 11 10:40:48 scrooge disk-health-nurse[26783]: [ID 702911 user.error] m:SY-mon-full-500 c:H : partition health measures for /var did not suffice - still using 96% of partition space
May 11 10:40:48 scrooge disk-health-nurse[26783]: [ID 702911 user.error] m:SY-mon-full-500 c:H : partition health measures for /var did not suffice - still using 96% of partition space
May 11 10:40:48 scrooge disk-health-nurse[26783]: [ID 702911 user.error] m:SY-mon-full-500 c:H : partition health measures for /var did not suffice - still using 96% of partition space

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

