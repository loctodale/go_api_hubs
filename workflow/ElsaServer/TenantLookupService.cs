using Elsa.Common.Multitenancy;

namespace ElsaServer;

public interface ITenantLookupService
{
    Task<Elsa.Common.Multitenancy.Tenant> GetTenantByApiKeyAsync(string apiKey);
}

public class TenantLookupService : ITenantLookupService
{
    private readonly Dictionary<string, Elsa.Common.Multitenancy.Tenant> _apiKeyToTenantMap = new()
    {
        // Sample API keys for testing
        { "acme-api-key-xyz123", new Elsa.Common.Multitenancy.Tenant { Id = "tenant-1", Name = "Tenant 1" } },
        { "contoso-api-key-abc789", new Elsa.Common.Multitenancy.Tenant { Id = "tenant-2", Name = "Tenant 2" } },
        { "globex-api-key-def456", new Elsa.Common.Multitenancy.Tenant { Id = "tenant-1", Name = "Tenant 1" } },
        
        // Add a default API key that maps to tenant-1 for testing
        { "default-api-key", new Elsa.Common.Multitenancy.Tenant { Id = "tenant-1", Name = "Tenant 1" } },
        { "test-api-key", new Elsa.Common.Multitenancy.Tenant { Id = "tenant-1", Name = "Tenant 1" } }
    };
    
    public Task<Elsa.Common.Multitenancy.Tenant> GetTenantByApiKeyAsync(string apiKey)
    {
        _apiKeyToTenantMap.TryGetValue(apiKey, out var tenant);
        
        // If no specific mapping found, return a default tenant
        tenant ??= new Elsa.Common.Multitenancy.Tenant { Id = "tenant-1", Name = "Tenant 1" };
        
        return Task.FromResult(tenant);
    }
}