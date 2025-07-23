
using Elsa.Services;

namespace ElsaServer;

public class MyTenantAccessor : ITenantAccessor
{
    private readonly IHttpContextAccessor _httpContextAccessor;

    public MyTenantAccessor(IHttpContextAccessor httpContextAccessor)
    {
        _httpContextAccessor = httpContextAccessor;
    }

    public Task<string> GetTenantIdAsync(CancellationToken cancellationToken = new CancellationToken())
    {
        var context = _httpContextAccessor.HttpContext;
        
        if (context == null)
            return Task.FromResult("default");

        // Try to get tenant ID from claims first
        var tenantClaim = context.User.Claims.FirstOrDefault(c => c.Type == "http://schemas.microsoft.com/identity/claims/tenantid");
        if (tenantClaim != null && !string.IsNullOrEmpty(tenantClaim.Value))
        {
            return Task.FromResult(tenantClaim.Value);
        }

        // Fallback to tenant ID from HttpContext items (set by tenant resolvers)
        if (context.Items.TryGetValue("TenantId", out var tenantIdObj) && tenantIdObj is string tenantId)
        {
            return Task.FromResult(tenantId);
        }

        return Task.FromResult("default");
    }
}