using Microsoft.Extensions.DependencyInjection;
using Payment.Data.Context;
using RequestLogs.Data.Context;
using RequestLogs.Domain.Contracts.UnitOfWork;

namespace RequestLogs.Repository.UnitOfWork;

public class UnitOfWork : IUnitOfWork
{
    protected RequestLogsDbContext _context;
    private readonly IServiceProvider _serviceProvider;

    public UnitOfWork(RequestLogsDbContext context, IServiceProvider serviceProvider)
    {
        _context = context;
        _serviceProvider = serviceProvider;
    }

    public async Task<bool> SaveChangesAsync(CancellationToken cancellationToken = default)
    {
        var result = await _context.SaveChangesAsync(cancellationToken);
        return result > 0;
    }

    public IRequestLogsRepository RequestLogs => _serviceProvider.GetService<IRequestLogsRepository>();
}