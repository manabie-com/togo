#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 5000

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["src/Services/Manabie.BasicIdentityServer.API/Manabie.BasicIdentityServer.API.csproj", "src/Services/Manabie.BasicIdentityServer.API/"]
COPY ["src/Services/Manabie.BasicIdentityServer.Domain/Manabie.BasicIdentityServer.Domain.csproj", "src/Services/Manabie.BasicIdentityServer.Domain/"]
COPY ["src/Services/Manabie.BasicIdentityServer.Infrastructure/Manabie.BasicIdentityServer.Infrastructure.csproj", "src/Services/Manabie.BasicIdentityServer.Infrastructure/"]
COPY ["src/Services/Manabie.BasicIdentityServer.Application/Manabie.BasicIdentityServer.Application.csproj", "src/Services/Manabie.BasicIdentityServer.Application/"]
RUN dotnet restore "src/Services/Manabie.BasicIdentityServer.API/Manabie.BasicIdentityServer.API.csproj"
COPY . .
WORKDIR "/src/src/Services/Manabie.BasicIdentityServer.API"
RUN dotnet build "Manabie.BasicIdentityServer.API.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Manabie.BasicIdentityServer.API.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Manabie.BasicIdentityServer.API.dll"]