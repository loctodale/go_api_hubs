FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS base
USER $APP_UID
WORKDIR /app
EXPOSE 8080
EXPOSE 8081

FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src
COPY ["ApiHubs.WebApp/ApiHubs.WebApp.csproj", "ApiHubs.WebApp/"]
RUN dotnet restore "ApiHubs.WebApp/ApiHubs.WebApp.csproj"
COPY . .
WORKDIR "/src/ApiHubs.WebApp"
RUN dotnet build "ApiHubs.WebApp.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "ApiHubs.WebApp.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "ApiHubs.WebApp.dll"]
