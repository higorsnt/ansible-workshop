#!/bin/bash

while ! nc -z kafka 9092; do
  sleep 1
done

java -jar app.jar