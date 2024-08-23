#!/bin/bash

export DATABASE_NAME="orchard.db"

# This script will generate the database inside the db folder
cd ../data
go run ../cmd/orchard/main.go