#!/bin/sh

./main http &
./main grpc &
wait
