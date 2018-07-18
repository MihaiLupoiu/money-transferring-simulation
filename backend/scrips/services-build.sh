#!/bin/bash

# TODO: Automate by path

cd ..
backendDir=$PWD
alpineVersion=1.10-alpine 

function createSeviceImage () {
    local _serviceName=$1
    
    echo -e "Compiling ${_serviceName}\n"
    docker run --rm -v $GOPATH:/go myhay/builder:$alpineVersion sh -c "cd /go/src/github.com/MihaiLupoiu/money-transferring-simulation/backend/services/${_serviceName} && CGO_ENABLED=0 GOOS=linux go build -installsuffix nocgo -o ${_serviceName}.alpine"
    ret=$?
    if [ $ret -ne 0 ]; then
        exit -1
    fi

    repository_name=mihailupoiu/${_serviceName}

    echo -e "Building ${_serviceName} docker image: ${repository_name}\n"
    echo -e "\n-------------------------------------------------------------------------"
    docker build -t $repository_name ${backendDir}/services/${_serviceName}
    ret=$?
    if [ $ret -ne 0 ]; then
        exit -1
    fi
    echo -e "-------------------------------------------------------------------------\n"
    
    echo -e "Removing ${_serviceName} binary"     
    rm -f $backendDir/services/${_serviceName}/${_serviceName}.alpine

}

##############################################################################
#                                   MAIN
##############################################################################

cd $backendDir/services/users
echo $PWD2
createSeviceImage "users"


if [ $(docker images -f "dangling=true" -q | wc -l) -ne 0 ]
then
    echo "Cleaning docker images not used."
    docker rmi $(docker images -f "dangling=true" -q) -f
fi