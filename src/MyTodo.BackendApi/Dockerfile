#See https://aka.ms/containerfastmode to understand how Visual Studio uses this Dockerfile to build your images for faster debugging.

FROM mcr.microsoft.com/dotnet/aspnet:3.1 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:3.1 AS build
WORKDIR /src
COPY ["src/MyTodo.BackendApi/MyTodo.BackendApi.csproj", "src/MyTodo.BackendApi/"]
COPY ["src/MyTodo.Utilities/MyTodo.Utilities.csproj", "src/MyTodo.Utilities/"]
COPY ["src/MyTodo.Services/MyTodo.Services.csproj", "src/MyTodo.Services/"]
COPY ["src/MyTodo.Data.EntityFramework/MyTodo.Data.EntityFramework.csproj", "src/MyTodo.Data.EntityFramework/"]
COPY ["src/MyTodo.Infrastructure/MyTodo.Infrastructure.csproj", "src/MyTodo.Infrastructure/"]
COPY ["src/MyTodo.Data/MyTodo.Data.csproj", "src/MyTodo.Data/"]
RUN dotnet restore "src/MyTodo.BackendApi/MyTodo.BackendApi.csproj"
COPY . .
WORKDIR "/src/src/MyTodo.BackendApi"
RUN dotnet build "MyTodo.BackendApi.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "MyTodo.BackendApi.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "MyTodo.BackendApi.dll"]
