#!/bin/bash

cd /tmp

if [ $(uname -v | grep Debian | wc -l) -ne 0 ]; then
    sudo apt-get install -y jq \
            build-essential \
            python-pip \
            apt-transport-https \
            ca-certificates \
            curl 
    pip install virtualenv

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

    echo -e "\n--------------------- Installing last stable version of kubernetes ------------------------"
    curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
    chmod +x ./kubectl
    sudo mv ./kubectl /usr/local/bin/kubectl


    echo -e "\n--------------------- Installing last stable version of minikube known ------------------------"
    content=$(curl https://api.github.com/repos/kubernetes/minikube/tags)
    version=$(echo $content | jq -r '.[0] | .name') 
    echo "Last version available ${version}"
    version=v0.28.2
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

elif [ $(uname -v | grep Darwin | wc -l) -ne 0 ]; then
    # https://gist.github.com/kevin-smets/b91a34cea662d0c523968472a81788f7

    which -s brew
    if [[ $? != 0 ]] ; then
        # Install Homebrew
        /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
    else
        brew update
    fi

    brew install kubernetes-cli
    brew cask install docker minikube virtualbox
    brew install kubernetes-helm

else
    echo "Not a Debian machine. Manual installation required. View script required steps."
fi







