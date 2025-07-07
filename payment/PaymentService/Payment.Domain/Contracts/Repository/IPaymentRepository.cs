using Payment.Data.Entities;

namespace Payment.Domain.Contracts.Repository;

public interface IPaymentRepository : IGenericRepository<TblPayment>
{
}