# -*- coding: utf-8 -*-

from buildbot.plugins import *
from buildbot.process.properties import WithProperties

tasks = util.BuildFactory()

go_env={
        'GOPATH': util.Interpolate('%(prop:builddir)s/go'),
        'GOROOT': '/usr/local/go',
        'GOBIN': util.Interpolate('%(prop:builddir)s/go/bin'),
        'PATH': util.Interpolate('/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin:/root/go/bin:/%(prop:builddir)s/go/bin')
}

basic_workdir='go/src/github.com/MihaiLupoiu/money-transferring-simulation'

# ==============================================================================

gofmt = steps.ShellCommand(name="go fmt",
        command=["gofmt", "-l", "models", "queries", "services"],
        env=go_env,
        workdir=basic_workdir+'/backend',
        haltOnFailure=False,
        flunkOnFailure=False)

govet = steps.ShellCommand(name="go vet",
        command=["go", "tool", "vet", "-shadow=true", "models/"],
        env=go_env,
        workdir=basic_workdir+'/backend',
        haltOnFailure=False,
        flunkOnFailure=False)

# ==============================================================================

# check out the source
tasks.addStep(steps.Git(
    repourl=util.Interpolate("%(src:tasks:repository)s"),
	mode='incremental',
	method='clean',
	branch=util.Interpolate('%(src:tasks:branch)s'),      
	codebase='tasks',
	workdir=basic_workdir,
        getDescription={
        "always":True,
        "tags": True,
        "long": True,
        "abbrev": 8}))

# TODO: Check if already installed
tasks.addStep(steps.ShellCommand(name="Go get dep",
        command=["go", "get", "-u", "github.com/golang/dep/cmd/dep"],
        env=go_env,
        workdir=basic_workdir,
        haltOnFailure=False,
        flunkOnFailure=False))

# TODO: Check if already installed
tasks.addStep(steps.ShellCommand(name="Go install dep",
        command=["go", "install", "github.com/golang/dep/cmd/dep"],
        env=go_env,
        workdir=basic_workdir,
        haltOnFailure=False,
        flunkOnFailure=False))

tasks.addStep(steps.ShellCommand(name="Ensure backend dependencies",
        command=["dep", "ensure"],
        env=go_env,
        workdir=basic_workdir+'/backend',
        haltOnFailure=True))

tasks.addSteps([gofmt, govet])

tasks.addStep(steps.ShellCommand(name="Build Service",
        command=["./services-build.sh", "tasks"],
        env=go_env,
        workdir=basic_workdir+'/backend/scripts',
        haltOnFailure=True))

tasks.addStep(steps.ShellCommand(name="Tag new docker image",
        command=util.renderer(lambda props: ["docker", "tag", "mihailupoiu/tasks:latest", "myhay/tasks:{}".format(props.getProperty('commit-description')['tasks'])]), 
        env=go_env,
        workdir=basic_workdir+'/backend/scripts',
        haltOnFailure=True))

tasks.addStep(steps.ShellCommand(name="Upload new docker image to dockerhub",
        command=util.renderer(lambda props: ["docker", "push", "myhay/tasks:{}".format(props.getProperty('commit-description')['tasks'])]), 
        env=go_env,
        workdir=basic_workdir+'/backend/scripts',
        haltOnFailure=True))

tasks.addStep(steps.ShellCommand(name="Deploy new docker image in Kubernetes",
        command=util.renderer(lambda props: ["./services-deploy.sh", "tasks", "{}".format(props.getProperty('commit-description')['tasks'])]), 
        env=go_env,
        workdir=basic_workdir+'/backend/scripts',
        haltOnFailure=True))
