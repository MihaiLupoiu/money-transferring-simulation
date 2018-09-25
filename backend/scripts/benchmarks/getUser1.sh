#!/bin/bash

while :
do
    sleep 1
    curl -i -X GET http://192.168.99.100:30080/api/v1/users/1
done