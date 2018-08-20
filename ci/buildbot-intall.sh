#!/bin/bash

pip install --upgrade pip
virtualenv env
source env/bin/activate
pip install 'buildbot[bundle]'

buildbot create-master ./master

mkdir worker
buildbot-worker create-worker ./worker localhost worker pass
