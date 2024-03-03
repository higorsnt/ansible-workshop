#!/bin/bash

while ! nc -z kafka 9092; do
  sleep 1
done

poetry run start