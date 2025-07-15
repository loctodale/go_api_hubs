using Payment.Data.Context;

namespace RequestLogs.Domain.Contracts.UnitOfWork;

public interface IUnitOfWork
{
    Task<bool> SaveChangesAsync(CancellationToken cancellationToken = default);
    IRequestLogsRepository RequestLogs { get; }
}