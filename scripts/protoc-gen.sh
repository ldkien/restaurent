#!/bin/bash

../protoc/bin/protoc --proto_path=../backend-entity/proto --proto_path=../protoc/include --go_out=plugins=grpc:../backend-entity ../backend-entity/proto/*.proto