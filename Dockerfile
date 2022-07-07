# Builds
FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build
WORKDIR /build
COPY . .
RUN dotnet restore
WORKDIR /build/TogoService.API
RUN dotnet build "TogoService.API.csproj" -c Release -o /app
RUN dotnet publish -c release -r debian-x64 -o /app

# Runs
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1 AS final
WORKDIR /app
ENV ASPNETCORE_URLS http://+:80
EXPOSE 80

ARG NewRelic="./newrelic"
COPY $NewRelic ./newrelic

RUN dpkg -i ./newrelic/newrelic-netcore20-agent*.deb
COPY --from=build /app .
ENTRYPOINT ["dotnet", "TogoService.API.dll"]