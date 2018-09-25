#!/bin/bash
echo "GET http://192.168.99.100:30080/api/v1/users" | vegeta attack -duration=10s -rate=100 | tee results-GET.bin | vegeta report