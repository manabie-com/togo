FROM mcr.microsoft.com/dotnet/core/aspnet:2.1-stretch-slim AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/core/sdk:2.1-stretch AS build
WORKDIR /src
COPY ["TogoManabie/TogoManabie.csproj", "TogoManabie/"]
RUN dotnet restore "TogoManabie/TogoManabie.csproj"
COPY . .
WORKDIR "/src/TogoManabie"
RUN dotnet build "TogoManabie.csproj" -c Release -o /app

FROM build AS publish
RUN dotnet publish "TogoManabie.csproj" -c Release -o /app

FROM base AS final
WORKDIR /app
COPY --from=publish /app .
ENTRYPOINT ["dotnet", "TogoManabie.dll"]