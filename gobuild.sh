#!/bin/bash

echo "Clearing dist folder..."
rm -rf ./dist
mkdir ./dist
echo "Compiling code..."
go build -o dist/workday-api
echo "Copying files..."
cp -r static/. dist/static/
cp -r templates/. dist/templates/
cp -r data/. dist/data/
echo "Done"