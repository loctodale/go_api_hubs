# Use .NET SDK to build the solution
FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /app

# Copy solution and restore dependencies
#COPY workflow/ElsaWorkflowService/ElsaWorkflowService.sln ./
COPY workflow/ElsaServer/ ./
RUN dotnet restore 

# Copy everything else and build the projects
#COPY . .
RUN dotnet publish ElsaServer.csproj -c Release -o /app/out

# Use ASP.NET runtime for the final image
FROM mcr.microsoft.com/dotnet/aspnet:8.0 AS runtime
WORKDIR /app
    
# Copy the built output
COPY --from=build /app/out .

## Expose ports
EXPOSE 8080

# Start Elsa Server
ENTRYPOINT ["dotnet", "ElsaServer.dll"]