#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/runtime:6.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["tests/Manabie.TestingApi.Application.IntegrationTests/Manabie.TestingApi.Application.IntegrationTests.csproj", "tests/Manabie.TestingApi.Application.IntegrationTests/"]
COPY ["src/Services/Manabie.Testing.Infrastructure/Manabie.Testing.Infrastructure.csproj", "src/Services/Manabie.Testing.Infrastructure/"]
COPY ["src/Services/Manabie.Testing.Application/Manabie.Testing.Application.csproj", "src/Services/Manabie.Testing.Application/"]
COPY ["src/Services/Manabie.Testing.Domain/Manabie.Testing.Domain.csproj", "src/Services/Manabie.Testing.Domain/"]
COPY ["src/Services/Manabie.Testing.API/Manabie.Testing.API.csproj", "src/Services/Manabie.Testing.API/"]
RUN dotnet restore "tests/Manabie.TestingApi.Application.IntegrationTests/Manabie.TestingApi.Application.IntegrationTests.csproj"
COPY . .
WORKDIR "/src/tests/Manabie.TestingApi.Application.IntegrationTests"
RUN dotnet build "Manabie.TestingApi.Application.IntegrationTests.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Manabie.TestingApi.Application.IntegrationTests.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Manabie.TestingApi.Application.IntegrationTests.dll"]