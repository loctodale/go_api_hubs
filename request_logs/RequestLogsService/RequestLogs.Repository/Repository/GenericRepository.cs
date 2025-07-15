using Microsoft.EntityFrameworkCore;
using Payment.Data.Context;
using RequestLogs.Data.Context;

namespace RequestLogs.Repository.Repository;

public class GenericRepository<T> : IGenericRepository<T> where T : class
{
    protected readonly RequestLogsDbContext _context;
    public GenericRepository(RequestLogsDbContext context) => _context = context;


    public async Task<T?> GetByIdAsync(Guid id)
    {
        return await _context.Set<T>().FindAsync(id);
    }

    public async Task<IEnumerable<T>> GetAllAsync()
    {
        return await _context.Set<T>().ToListAsync();
    }

    public async Task AddAsync(T entity)
    {
        await _context.Set<T>().AddAsync(entity);
    }

    public async Task AddRangeAsync(IEnumerable<T> entities)
    {
        await _context.Set<T>().AddRangeAsync(entities);
    }

    public void Update(T entity)
    {
        _context.Set<T>().Update(entity);
    }

    public void DeleteAsync(T entity)
    {
        _context.Set<T>().Remove(entity);
    }
}
