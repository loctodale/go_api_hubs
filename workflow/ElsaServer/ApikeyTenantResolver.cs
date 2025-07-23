using Elsa.Common.Multitenancy;

namespace ElsaServer;

public class ApiKeyTenantResolver : ITenantResolver
{
    private readonly IHttpContextAccessor _httpContextAccessor;
    private readonly ITenantLookupService _tenantLookupService;
    
    public ApiKeyTenantResolver(
        IHttpContextAccessor httpContextAccessor,
        ITenantLookupService tenantLookupService)
    {
        _httpContextAccessor = httpContextAccessor;
        _tenantLookupService = tenantLookupService;
    }
    
    public async Task<TenantResolverResult> ResolveAsync(TenantResolverContext context)
    {
        var httpContext = _httpContextAccessor.HttpContext;
        if (httpContext == null)
            return TenantResolverResult.Unresolved();
            
        // Try multiple header patterns
        var apiKey = GetApiKey(httpContext);
        if (string.IsNullOrEmpty(apiKey))
            return TenantResolverResult.Unresolved();
            
        // Look up tenant by API key
        var tenant = await _tenantLookupService.GetTenantByApiKeyAsync(apiKey);
        if (tenant == null)
            return TenantResolverResult.Unresolved();
            
        // Store tenant ID in HttpContext for MyTenantAccessor to access
        httpContext.Items["TenantId"] = tenant.Id;
            
        return TenantResolverResult.Resolved(tenant.Id);
    }
    
    private string GetApiKey(HttpContext httpContext)
    {
        // Try different header names
        return httpContext.Request.Headers["X-API-Key"].FirstOrDefault()
               ?? httpContext.Request.Headers["Authorization"].FirstOrDefault()?.Replace("Bearer ", "")
               ?? httpContext.Request.Headers["X-Tenant-Key"].FirstOrDefault()
               ?? httpContext.Request.Headers["X-Api-Key"].FirstOrDefault(); // Added this variant
    }
}