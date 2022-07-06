set testingOutputPath=./publish/Manabie.Testing.API
set testingUrl=http://localhost:7216
set configuration=Debug
set environment = Development


title API

cd %testingOutputPath%
dotnet Manabie.Testing.API.dll  --urls=%testingUrl% -c %configuration% -e %environment%