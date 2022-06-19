#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:6.0 AS base
WORKDIR /app

FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /src
COPY ["Todo.Api/Todo.Api.csproj", "Todo.Api/"]
COPY ["Todo.Application/Todo.Application.csproj", "Todo.Application/"]
COPY ["Todo.Domain/Todo.Domain.csproj", "Todo.Domain/"]
COPY ["Todo.Infrastructure/Todo.Infrastructure.csproj", "Todo.Infrastructure/"]
COPY ["Todo.Api.Test/Todo.Api.Test.csproj", "Todo.Api.Test/"]
COPY ["Todo.Application.Test/Todo.Application.Test.csproj", "Todo.Application.Test/"]
RUN dotnet restore "Todo.Api/Todo.Api.csproj"
RUN dotnet restore "Todo.Api.Test/Todo.Api.Test.csproj"
RUN dotnet restore "Todo.Application.Test/Todo.Application.Test.csproj"
COPY . .

RUN dotnet test Todo.Api.Test
RUN dotnet test Todo.Application.Test

WORKDIR "/src/Todo.Api"
RUN dotnet build "Todo.Api.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "Todo.Api.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Todo.Api.dll"]