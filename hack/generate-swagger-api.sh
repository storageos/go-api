#!/bin/sh
set -eu

swagger generate model -f swagger.yaml \
    -m types --skip-validator -C swagger-gen.yaml \
    -n ErrorResponse \
    -n Deployment \
    -n Volume \
    -n CreateVolumeOptions
