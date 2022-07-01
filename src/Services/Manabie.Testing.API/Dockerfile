#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 500

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["src/Services/Manabie.Testing.API/Manabie.Testing.API.csproj", "src/Services/Manabie.Testing.API/"]
COPY ["src/Services/Manabie.Testing.Application/Manabie.Testing.Application.csproj", "src/Services/Manabie.Testing.Application/"]
COPY ["src/Services/Manabie.Testing.Domain/Manabie.Testing.Domain.csproj", "src/Services/Manabie.Testing.Domain/"]
COPY ["src/Services/Manabie.Testing.Infrastructure/Manabie.Testing.Infrastructure.csproj", "src/Services/Manabie.Testing.Infrastructure/"]
RUN dotnet restore "src/Services/Manabie.Testing.API/Manabie.Testing.API.csproj"
COPY . .
WORKDIR "/src/src/Services/Manabie.Testing.API"
RUN dotnet build "Manabie.Testing.API.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Manabie.Testing.API.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Manabie.Testing.API.dll"]