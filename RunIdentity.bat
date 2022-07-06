set identityServerOutputPath=./publish/Manabie.BasicIdentityServer.API
set identityServerUrl=http://localhost:5000
set configuration=Debug
set environment = Development
title Identity

cd %identityServerOutputPath%
dotnet Manabie.BasicIdentityServer.API.dll  --urls=%identityServerUrl% -c %configuration% -e %environment%
