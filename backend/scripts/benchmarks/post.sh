#!/bin/bash
jq -ncM '{method: "POST", url: "http://192.168.99.100:30080/api/v1/users", body: { "firstname": "Jhon", "lastname": "Donals", "mail": "jd@fake.com"} | @base64, header: {"Content-Type": ["application/json"]}}' | vegeta attack -body addUser.json -format=json -duration=10s -rate=100  | tee results-POST.bin | vegeta report

#http://192.168.99.100:30080/api/v1/users
