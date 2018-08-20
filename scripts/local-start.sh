#!/bin/bash

minikube start \
    --cpus=2 \
    --memory=4096 \
    --v=5 --logtostderr 
    
    
    # -vm-driver=none \

sleep 5

# Create roles:
# Fix Dashboar
kubectl create clusterrolebinding add-on-cluster-admin --clusterrole=cluster-admin --serviceaccount=kube-system:default

helm init --history-max 4

maxRetrys=15
count=0
until [ $(helm list 2>&1 | wc -c ) -lt 2 ]
do
    sleep 5
    if [  $(helm list 2>&1 | wc -c ) -gt 150 ];then
        exit 0        
    fi

    if [ $count == $maxRetrys ];then
        exit 1
    fi
    let count=$count+1
    echo $count
done

# Enable Addons
# minikube addons enable efk 
minikube addons enable heapster
# minikube addons enable ingress
# minikube addons enable registry
