# Stage 1: Build
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src

# Copy project file (from docker-compose context: . = root, so workflow/ElsaStudio/ElsaStudio.csproj)
COPY workflow/ElsaStudio/ElsaStudio.csproj ElsaStudio/
RUN dotnet restore ElsaStudio/ElsaStudio.csproj

# Copy all ElsaStudio files
COPY workflow/ElsaStudio/ ElsaStudio/
WORKDIR /src/ElsaStudio
RUN dotnet publish -c Release -o /out

# Stage 2: Runtime - Use nginx to serve static files
FROM nginx:alpine AS runtime
WORKDIR /usr/share/nginx/html

# Copy the published Blazor WASM files
COPY --from=build /out/wwwroot .

# Copy nginx configuration for Blazor WASM
COPY --from=build /src/ElsaStudio/nginx.conf /etc/nginx/nginx.conf

# Expose port 8080
EXPOSE 8080

# Start nginx
CMD ["nginx", "-g", "daemon off;"]
