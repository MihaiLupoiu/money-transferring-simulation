#!/bin/bash

source env/bin/activate
buildbot restart ./master
buildbot-worker restart ./worker
