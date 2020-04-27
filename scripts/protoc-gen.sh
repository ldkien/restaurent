#!/bin/bash

/Users/ledoankien/Downloads/protoc-3.11.4-osx-x86_64/bin/protoc --proto_path=/Users/ledoankien/Go-Projects/src/restaurant/backend-entity/proto --proto_path=/Users/ledoankien/Downloads/protoc-3.11.4-osx-x86_64/include --go_out=plugins=grpc:/Users/ledoankien/Go-Projects/src/restaurant/backend-entity /Users/ledoankien/Go-Projects/src/restaurant/backend-entity/proto/*.proto