using Payment.Data.Context;
using RequestLogs.Domain.Contracts.UnitOfWork;
using RequestLogs.Repository.Repository;
using RequestLogs.Repository.UnitOfWork;
using RequestLogsService.Services;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddGrpc();
builder.Services.AddScoped<IRequestLogsRepository, RequestLogsRepository>();
builder.Services.AddScoped<IUnitOfWork, UnitOfWork>();
var app = builder.Build();

// Configure the HTTP request pipeline.
app.MapGrpcService<RequestLogsService.Services.RequestLogsService>();
app.MapGet("/",
    () =>
        "Communication with gRPC endpoints must be made through a gRPC client. To learn how to create a client, visit: https://go.microsoft.com/fwlink/?linkid=2086909");

app.Run();