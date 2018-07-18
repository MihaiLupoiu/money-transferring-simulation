#!/bin/bash

pip install --upgrade pip
virtualenv env
source env/bin/activate
pip install 'buildbot[bundle]'

buildbot create-master ./master


buildbot-worker create-worker . localhost worker pass
