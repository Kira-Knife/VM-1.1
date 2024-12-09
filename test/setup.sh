#!/bin/bash

set -e

cd docker-openldap
sudo docker build -t my/openldap .
