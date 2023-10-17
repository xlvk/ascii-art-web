#!/bin/bash

docker build -t ascii-web .
docker run -p 2003:2003 ascii-web