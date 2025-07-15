FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
USER $APP_UID
WORKDIR /app
EXPOSE 8080
ENV ASPNETCORE_URLS=http://+:8080
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["/request_logs/RequestLogsService/RequestLogsService/RequestLogsService.csproj", "RequestLogsService/"]
RUN dotnet restore "RequestLogsService/RequestLogsService.csproj"
COPY ./request_logs/RequestLogsService .
WORKDIR "/src/RequestLogsService"
RUN dotnet build "RequestLogsService.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "RequestLogsService.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "RequestLogsService.dll"]
