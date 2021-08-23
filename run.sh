#!/bin/bash

# Clean up previous builds
rm -rf ./publish

mkdir ./publish
dotnet restore ./src/TodoNET/Api/TodoNet.Api/TodoNet.Api.csproj
dotnet build --configuration release --no-restore ./src/TodoNET/Api/TodoNet.Api/TodoNet.Api.csproj
dotnet publish --configuration release --no-restore --output ./publish ./src/TodoNET/Api/TodoNet.Api/TodoNet.Api.csproj
cd publish
dotnet TodoNet.Api.dll