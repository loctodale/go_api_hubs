using Elsa.Common.Multitenancy;

namespace ElsaServer;

public class SubdomainTenantResolverAsync : ITenantResolver
{
    private readonly IHttpContextAccessor _httpContextAccessor;
    
    public SubdomainTenantResolverAsync(IHttpContextAccessor httpContextAccessor)
    {
        _httpContextAccessor = httpContextAccessor;
    }
    
    public async Task<TenantResolverResult> ResolveAsync(TenantResolverContext context)
    {
        var httpContext = _httpContextAccessor.HttpContext;
        
        if (httpContext == null)
            return TenantResolverResult.Unresolved();
            
        var host = httpContext.Request.Host.Value;
        var subdomain = ExtractSubdomain(host);
        
        if (string.IsNullOrEmpty(subdomain))
            return TenantResolverResult.Unresolved();
            
        // If you need to do async operations like database lookups:
        // var tenant = await _tenantService.GetTenantBySubdomainAsync(subdomain);
        
        var tenant = new Elsa.Common.Multitenancy.Tenant
        {
            Id = subdomain,
            Name = $"Tenant {subdomain}"
        };
        
        return TenantResolverResult.Resolved(tenant.Id);
    }
    
    private string ExtractSubdomain(string host)
    {
        var parts = host.Split('.');
        return parts.Length > 2 ? parts[0] : null;
    }
}