using Payment.Domain.Contracts.Repository;

namespace Payment.Domain.Contracts.UnitOfWork;

public interface IUnitOfWork
{
    Task<bool> SaveChangesAsync(CancellationToken cancellationToken = default);
    IPaymentRepository Payment { get; }
}