#!/bin/sh

grpcui --plaintext --import-path /work/hybrid/envoy/googleapis/ --import-path /work/hybrid/envoy/data-plane-api/ --import-path /work/hybrid/envoy/protoc-gen-validate  --proto /work/hybrid/envoy/protoc-gen-validate/validate/validate.proto --proto /work/hybrid/envoy/data-plane-api/envoy/api/v2/rds.proto --proto /work/hybrid/envoy/data-plane-api/envoy/api/v2/lds.proto --proto /work/hybrid/envoy/data-plane-api/envoy/api/v2/cds.proto --bind 0.0.0.0 --port 38080 localhost:12000
