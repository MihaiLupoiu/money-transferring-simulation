#!/bin/bash

cd /tmp

echo -e "\n--------------------- Installing last stable version of kubernetes ------------------------"
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl


# TODO: Update to 0.28.0
echo -e "\n--------------------- Installing last stable version of minikube known ------------------------"
content=$(curl https://api.github.com/repos/kubernetes/minikube/tags)
version=$(echo $content | jq -r '.[0] | .name') 
echo "Last version available ${version}"
version=v0.25.2
echo "Installing version ${version}"
curl -Lo minikube https://storage.googleapis.com/minikube/releases/$version/minikube-linux-amd64 && chmod +x minikube 
sudo mv minikube /usr/local/bin/

echo -e "\n--------------------- Installing last stable version of helm ------------------------"
content=$(curl https://api.github.com/repos/kubernetes/helm/tags) 
version=$(echo $content | jq -r '.[] | .name' | grep -v rc | awk 'NR==1{print $1}') 
echo $version
curl -LO https://storage.googleapis.com/kubernetes-helm/helm-$version-linux-amd64.tar.gz && tar -zxvf helm-$version-linux-amd64.tar.gz 
sudo mv linux-amd64/helm /usr/local/bin/helm && rm -r linux-amd64 && rm helm*

 # Install Go dep
if [ $(dep version | wc -l) -eq 0 ]; then
    go get -u -v github.com/golang/dep/cmd/dep && go install github.com/golang/dep/cmd/dep
fi