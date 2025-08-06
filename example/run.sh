#!/bin/bash

PROMETHEUS_PORT=9090

teardown() {
  echo 'Stopping server...'
  kill -9 $SERVER_PID
  rm agent.jar
  exit 0
}

trap teardown SIGINT

cd "$(dirname "$0")"

echo "Installing latest binary"
go install github.com/picadoh/jload@latest

echo "Downloading Prometheus JMX exporter"
wget https://repo.maven.apache.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/1.0.1/jmx_prometheus_javaagent-1.0.1.jar -O agent.jar

echo "Starting Java HTTP server"
java -XX:+EnableDynamicAgentLoading HelloServer.java &

SERVER_PID=$!

while ! nc -z localhost 8080; do   
  sleep 1 
done

echo "Attaching agent to $SERVER_PID"
jload $SERVER_PID $PWD/agent.jar=$PROMETHEUS_PORT:$PWD/config.yml

echo "Visit http://localhost:$PROMETHEUS_PORT"

read -p "Press Ctrl+C to terminate"

wait $SERVER_PID 2>/dev/null

