#!/bin/bash
helm install --name=postgresql --set postgresPassword="abc123"../../charts/postgresql