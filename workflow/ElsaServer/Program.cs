using Elsa;
using Elsa.Common.Multitenancy;
using Elsa.EntityFrameworkCore;
using Elsa.EntityFrameworkCore.Extensions;
using Elsa.EntityFrameworkCore.Modules.Management;
using Elsa.EntityFrameworkCore.Modules.Runtime;
using Elsa.EntityFrameworkCore.Modules.Tenants;
using Elsa.Extensions;
using Elsa.Identity.Multitenancy;
using Elsa.Identity.Options;
using Elsa.Tenants;
using Elsa.Tenants.AspNetCore;
using Elsa.Tenants.Extensions;
using ElsaServer;
using ElsaServer.Models;
using FastEndpoints.Swagger;
using Microsoft.AspNetCore.Mvc;
using ITenantAccessor = Elsa.Services.ITenantAccessor;

WebApplicationBuilder builder = WebApplication.CreateBuilder(args);
ConfigurationManager configuration = builder.Configuration;

var identitySection = configuration.GetSection("Identity");
var identityTokenSection = identitySection.GetSection("Tokens");

builder.Services.AddControllers();
builder.Services.AddHttpContextAccessor();
builder.Services.AddSingleton<IHttpContextAccessor, HttpContextAccessor>();
builder.Services.AddDbContext<TenantServiceContext>();
builder.Services.Configure<IdentityTokenOptions>(options =>
{
    options.TenantIdClaimsType = "http://schemas.microsoft.com/identity/claims/tenantid";
});

builder.Services.AddScoped<ITenantStore,MemoryTenantStore>();
builder.Services.AddElsa(elsa =>
{
    var dbContextOptions = new ElsaDbContextOptions();
    string postgresConnectionString = configuration.GetConnectionString("Postgresql")!;
    string schema = configuration.GetConnectionString("Schema")!;

    if (!string.IsNullOrEmpty(schema))
    {
        dbContextOptions.SchemaName = schema;
        dbContextOptions.MigrationsAssemblyName = typeof(Program).Assembly.GetName().Name;
    }
    elsa.UseWorkflowManagement(management => 
        management
            .UseEntityFrameworkCore(ef => ef.UsePostgreSql(postgresConnectionString, dbContextOptions)));
    elsa.UseWorkflowRuntime(runtime => runtime.UseEntityFrameworkCore(ef => ef.UsePostgreSql(postgresConnectionString, dbContextOptions)));

    elsa.UseSasTokens()
         .UseIdentity(identity =>
         {
             identity.TokenOptions = options => identityTokenSection.Bind(options);
             identity.UseConfigurationBasedApplicationProvider(options => identitySection.Bind(options));
             identity.UseConfigurationBasedRoleProvider(options => identitySection.Bind(options));
         })
         .UseDefaultAuthentication();
    // elsa.UseTenantHttpRouting(http => http.WithTenantHeader("X-Tenant-Id"));
    elsa.UseTenants(tenantsFeature =>
    {
      
        tenantsFeature.UseTenantManagement(manage =>
        {
            manage.Services.AddScoped<ITenantsProvider, DatabaseTenantsProvider>();
            manage.UseEntityFrameworkCore(ef => ef.UsePostgreSql(configuration.GetConnectionString("Tenant")));
        });
        tenantsFeature.ConfigureMultitenancy(options =>
        {
            options.TenantResolverPipelineBuilder
                .Clear()
                .Append<ClaimsTenantResolver>()
                .Append<CurrentUserTenantResolver>()
                .Append<HeaderTenantResolver>();
        });

       
    });

    elsa
        .UseHttp()
        .AddFastEndpointsAssembly<Program>()
        .UseWorkflowsApi();

    elsa.AddWorkflowsFrom<Program>();
});
builder.Services.AddScoped<ITenantAccessor, MyTenantAccessor>();
builder.Services.AddScoped<ITenantsProvider, DatabaseTenantsProvider>();

builder.Services.SwaggerDocument(options =>
{
    options.DocumentSettings = documentSetting =>
    {
        documentSetting.Title = "Elsa API";
        documentSetting.Version = "v1";
    };
});
builder.Services.AddCors(cors => cors.AddDefaultPolicy(policy => policy.AllowAnyHeader().AllowAnyMethod().AllowAnyOrigin().WithExposedHeaders("*")));
builder.Services.AddRazorPages(options => options.Conventions.ConfigureFilter(new IgnoreAntiforgeryTokenAttribute()));
var app = builder.Build();

// Configure the HTTP request pipeline.
app.UseHttpsRedirection();
app.MapControllers();
app.UseWorkflows();
app.UseWorkflowsApi();
app.UseBlazorFrameworkFiles();
app.UseRouting();
app.UseCors();
app.UseStaticFiles();
app.UseAuthentication();
app.UseAuthorization();
app.UseTenants();
app.MapFallbackToPage("/_Host");
if (!app.Environment.IsProduction())
{
    EndpointSecurityOptions.SecurityIsEnabled = false;
    app.UseDeveloperExceptionPage();
    app.UseOpenApi();
    app.UseSwaggerUi();
    app.UseReDoc();
}

app.Run();