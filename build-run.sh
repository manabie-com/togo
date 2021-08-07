#!/bin/bash

# Clean up previous builds
rm -rf ./publish

mkdir ./publish
dotnet restore ./src/togo.Api/togo.Api.csproj
dotnet build --configuration Release --no-restore ./src/togo.Api/togo.Api.csproj
dotnet publish --configuration Release --no-restore --output ./publish ./src/togo.Api/togo.Api.csproj
dotnet ./publish/togo.Api.dll