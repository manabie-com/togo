set srcIdentityServerPath=./src/services/Manabie.BasicIdentityServer.API
set srcTestingPath=./../Manabie.Testing.API

set identityServerOutputPath=./../../../publish/Manabie.BasicIdentityServer.API
set testingOutputPath=./../../../publish/Manabie.Testing.API

set configuration=Debug

cd %srcIdentityServerPath%
dotnet publish -c %configuration% -o %identityServerOutputPath%

cd %srcTestingPath%
dotnet publish -c %configuration% -o %testingOutputPath%

cd ../../../

call RunIdentity.bat

call RunTestingApi.bat

pause 