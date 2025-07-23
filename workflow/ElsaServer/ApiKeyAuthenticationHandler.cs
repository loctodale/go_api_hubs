using Microsoft.AspNetCore.Authentication;
using Microsoft.Extensions.Options;
using System.Security.Claims;
using System.Text.Encodings.Web;

namespace ElsaServer;

public class ApiKeyAuthenticationSchemeOptions : AuthenticationSchemeOptions
{
    public const string DefaultScheme = "CustomApiKey";
    public string ApiKeyHeaderName { get; set; } = "X-API-Key";
}

public class ApiKeyAuthenticationHandler : AuthenticationHandler<ApiKeyAuthenticationSchemeOptions>
{
    private readonly ITenantLookupService _tenantLookupService;

    public ApiKeyAuthenticationHandler(
        IOptionsMonitor<ApiKeyAuthenticationSchemeOptions> options,
        ILoggerFactory logger,
        UrlEncoder encoder,
        ITenantLookupService tenantLookupService)
        : base(options, logger, encoder)
    {
        _tenantLookupService = tenantLookupService;
    }

    protected override async Task<AuthenticateResult> HandleAuthenticateAsync()
    {
        // Check if API key is present
        var apiKey = GetApiKey();
        if (string.IsNullOrEmpty(apiKey))
        {
            return AuthenticateResult.NoResult();
        }

        // Validate API key and get tenant
        var tenant = await _tenantLookupService.GetTenantByApiKeyAsync(apiKey);
        if (tenant == null)
        {
            return AuthenticateResult.Fail("Invalid API key");
        }

        // Create claims for the authenticated API key
        var claims = new[]
        {
            new Claim(ClaimTypes.Name, $"ApiKey-{tenant.Id}"),
            new Claim(ClaimTypes.NameIdentifier, tenant.Id),
            new Claim("tenant_id", tenant.Id),
            new Claim("auth_method", "apikey")
        };

        var identity = new ClaimsIdentity(claims, Scheme.Name);
        var principal = new ClaimsPrincipal(identity);
        var ticket = new AuthenticationTicket(principal, Scheme.Name);

        return AuthenticateResult.Success(ticket);
    }

    private string? GetApiKey()
    {
        // Try different header patterns
        return Context.Request.Headers["X-API-Key"].FirstOrDefault()
               ?? Context.Request.Headers["Authorization"].FirstOrDefault()?.Replace("Bearer ", "")
               ?? Context.Request.Headers["X-Tenant-Key"].FirstOrDefault();
    }
} 