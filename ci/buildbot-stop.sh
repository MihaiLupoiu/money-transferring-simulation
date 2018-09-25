#!/bin/bash

source env/bin/activate
buildbot stop ./master
buildbot-worker stop ./worker
