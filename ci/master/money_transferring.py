# -*- coding: utf-8 -*-

from buildbot.plugins import *
from buildbot.process.properties import WithProperties

money_transferring = util.BuildFactory()

go_env={
        'GOPATH': util.Interpolate('%(prop:builddir)s/go'),
        'GOROOT': '/usr/local/go',
        'GOBIN': util.Interpolate('%(prop:builddir)s/go/bin'),
        'PATH': util.Interpolate('/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin:/root/go/bin:/%(prop:builddir)s/go/bin')
}

basic_workdir='go/src/MihaiLupoiu/money-transferring-simulation'

# check out the source
money_transferring.addStep(steps.Git(
    repourl=util.Interpolate("%(src:money-transferring-simulation:repository)s"),
	mode='incremental',
	method='clean',
	branch=util.Interpolate('%(src:money-transferring-simulation:branch)s'),      
	codebase='money-transferring-simulation',
	workdir=basic_workdir))

# TODO: Chexk if already installed

money_transferring.addStep(steps.ShellCommand(name="Go get dep",
        command=["go", "get", "-u", "github.com/golang/dep/cmd/dep"],
        env=go_env,
        workdir=basic_workdir,
        haltOnFailure=False,
        flunkOnFailure=False))

money_transferring.addStep(steps.ShellCommand(name="Go install dep",
        command=["go", "install", "github.com/golang/dep/cmd/dep"],
        env=go_env,
        workdir=basic_workdir,
        haltOnFailure=False,
        flunkOnFailure=False))

money_transferring.addStep(steps.ShellCommand(name="Ensure backend dependencies",
        command=["dep", "ensure"],
        env=go_env,
        workdir=basic_workdir+'/backend',
        haltOnFailure=True))

money_transferring.addStep(steps.ShellCommand(name="Build Service",
        command=["services-build.sh"],
        env=go_env,
        workdir=basic_workdir+'/backend/scripts',
        haltOnFailure=True))

