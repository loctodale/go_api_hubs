FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS base
USER $APP_UID
WORKDIR /app
EXPOSE 8080
ENV ASPNETCORE_URLS=http://+:8080
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
ARG BUILD_CONFIGURATION=Release
WORKDIR /src

COPY ["workflow/Tenant.Grpc/Tenant.Grpc.csproj", "Tenant.Grpc/"]
RUN dotnet restore "Tenant.Grpc/Tenant.Grpc.csproj"

COPY ./workflow/Tenant.Grpc/ ./Tenant.Grpc/

WORKDIR "/src/Tenant.Grpc"
RUN dotnet build "Tenant.Grpc.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "Tenant.Grpc.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "Tenant.Grpc.dll"]
