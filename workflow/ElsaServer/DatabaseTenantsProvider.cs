using Elsa.Common.Multitenancy;
using ElsaServer.Models;
using Microsoft.EntityFrameworkCore;

namespace ElsaServer;

public class DatabaseTenantsProvider : ITenantsProvider
{
    private readonly TenantServiceContext _context;
    public DatabaseTenantsProvider(TenantServiceContext context) => _context = context;

    private static Elsa.Common.Multitenancy.Tenant MapToTenant(TblTenantService entity)
    {
        return new Elsa.Common.Multitenancy.Tenant
        {
            Id = entity.TenantId.ToString(),
            Name = entity.Name,
            TenantId = entity.TenantId.ToString(),
        };
    }
    public async Task<IEnumerable<Tenant>> ListAsync(CancellationToken cancellationToken = default (CancellationToken))
    {
        try
        {
            var tenantEntities = await _context.TblTenantServices.ToListAsync();
            
            return tenantEntities.Select(MapToTenant);
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            return Enumerable.Empty<Elsa.Common.Multitenancy.Tenant>();
        }
    }

    public async Task<Elsa.Common.Multitenancy.Tenant?> FindAsync(TenantFilter filter, CancellationToken cancellationToken = new CancellationToken())
    {
        try
        {
            var tenantEntity = await _context.TblTenantServices.FirstOrDefaultAsync(x => x.TenantId.ToString() == filter.Id);
            return tenantEntity != null ? MapToTenant(tenantEntity) : null;
        }
        catch (Exception e)
        {
            Console.WriteLine(e);
            return null;
        }
    }
}