using Elsa.Common.Multitenancy;
using Elsa.Common.Services;
using Elsa.Tenants.Providers;
using ElsaServer.Models;
using JetBrains.Annotations;
using Microsoft.EntityFrameworkCore;

namespace Elsa.Tenants;

[UsedImplicitly]
public class MemoryTenantStore : ITenantStore
{
    private readonly MemoryStore<Tenant> store;

    public MemoryTenantStore(TenantServiceContext context)
    {
        store = new MemoryStore<Tenant>();
        var data = context.TblTenantServices.ToList().Select(x => new Tenant
        {
            Id = x.Id.ToString(),
            TenantId = x.TenantId.ToString(),
            Name = x.Name,
        });
        store.AddMany(data, x => x.Id);
        Console.WriteLine(store);
    }
    public Task<Tenant?> FindAsync(TenantFilter filter, CancellationToken cancellationToken = default)
    {
        var result = store.Query(query => Filter(query, filter)).FirstOrDefault();
        return Task.FromResult(result);
    }

    public Task<Tenant?> FindAsync(string id, CancellationToken cancellationToken = default)
    {
        Console.WriteLine($"FindAsync called with {id}");
        var filter = TenantFilter.ById(id);
        return FindAsync(filter, cancellationToken);
    }

    public Task<IEnumerable<Tenant>> FindManyAsync(TenantFilter filter, CancellationToken cancellationToken = default)
    {
        Console.WriteLine($"FindAsync called with {filter.Id}");
        var result = store.Query(query => Filter(query, filter)).ToList().AsEnumerable();
        return Task.FromResult(result);
    }

    public Task<IEnumerable<Tenant>> ListAsync(CancellationToken cancellationToken = default)
    {
        Console.WriteLine($"List async tenant");
        var result = store.List();
        return Task.FromResult(result);
    }

    public Task AddAsync(Tenant tenant, CancellationToken cancellationToken = default)
    {
        Console.WriteLine("AddAsync called");
        store.Add(tenant, GetId);
        return Task.CompletedTask;
    }

    public Task UpdateAsync(Tenant tenant, CancellationToken cancellationToken = default)
    {
        store.Update(tenant, GetId);
        return Task.CompletedTask;
    }

    public Task<bool> DeleteAsync(string id, CancellationToken cancellationToken = default)
    {
        var found = store.Delete(id);
        return Task.FromResult(found);
    }

    public Task<long> DeleteAsync(TenantFilter filter, CancellationToken cancellationToken = default)
    {
        var deletedCount = store.DeleteMany(filter.Apply(store.Queryable), GetId);
        return Task.FromResult(deletedCount);
    }
    
    private IQueryable<Tenant> Filter(IQueryable<Tenant> queryable, TenantFilter filter) => filter.Apply(queryable);

    private string GetId(Tenant tenant) => tenant.Id;
}