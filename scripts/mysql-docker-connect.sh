#!/usr/bin/env bash

# Use this to connect to the server and run commands directly
# you can also just use your favorite admin

host=127.0.0.1
pass=YOUR_PASSWORD_HERE

mysql --host=$host --port=3306 \
  --user=root --password=$pass \
  --database=litt


