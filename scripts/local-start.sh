#!/bin/bash

# VirtualBox - Debian
if [ $(virtualbox -h | wc -l) -eq 0 ]; then
    echo -e "\n--------------------- Installing Virtualbox ------------------------"
    echo "deb https://download.virtualbox.org/virtualbox/debian stretch contrib" | sudo tee /etc/apt/sources.list.d/virtualbox.list
    wget -q https://www.virtualbox.org/download/oracle_vbox_2016.asc -O- | sudo apt-key add -
    wget -q https://www.virtualbox.org/download/oracle_vbox.asc -O- | sudo apt-key add -
    sudo apt-get update
    sudo apt-get install virtualbox-5.2 -y
fi

 # Install Docker
if [ $(docker --version | wc -l) -eq 0 ]; then
    echo -e "\n--------------------- Installing Docker ------------------------"
    curl -fsSL https://download.docker.com/linux/debian/gpg | sudo apt-key add -
    echo "deb [arch=amd64] https://download.docker.com/linux/debian   $(lsb_release -cs)    stable" | sudo tee /etc/apt/sources.list.d/docker.list
    sudo apt-get update
    sudo apt-get install docker-ce -y
    # sudo usermod -a -G docker $USER
fi

minikube start \
    --cpus=4 \
    --memory=4096 \
    --kubernetes-version v1.9.0 \
    --extra-config=kubelet.CAdvisorPort=4194 \
    --extra-config=apiserver.Authorization.Mode=RBAC \
    --extra-config=apiserver.ServiceNodePortRange=1-50000 \

# TODO: Change to 
    # --cpus=2 \
    # --memory=2048 \
    # --kubernetes-version v1.9.0 \
    
# -------------------------------------------------------------
    #--v=5 --logtostderr 
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
