#!/bin/bash

cd ..
backendDir=$PWD

serviceName=$1
tag=$2

serviceRunning=$(helm list -a -q ${serviceName} | wc -l )
    if [ $serviceRunning -eq 0 ]; then
        echo -e "\nHelm install service ${serviceName}"

        helm install --name=${serviceName} --set image.tag=$tag -- ../charts/${serviceName}
        ret=$?
        if [ $ret -ne 0 ]; then
            exit -1
        fi

    else
        echo -e "\nHelm upgrading service ${_serviceName}"
        helm upgrade ${serviceName} --set image.tag=$tag -- ../charts/${serviceName}
        ret=$?
        if [ $ret -ne 0 ]; then
            exit -1
        fi
    fi
