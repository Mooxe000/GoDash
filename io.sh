#!/usr/bin/env bash

set -e

curl -o ./io.xml -L \
  "https://raw.githubusercontent.com/apache/activemq/d6682e5476cd8cbefca04227ffa26a5d508d2494/assembly/src/release/conf/jetty.xml"
