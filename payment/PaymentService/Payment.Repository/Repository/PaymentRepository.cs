using Payment.Data.Context;
using Payment.Data.Entities;
using Payment.Domain.Contracts.Repository;

namespace Payment.Repository.Repository;

public class PaymentRepository : GenericRepository<TblPayment>, IPaymentRepository
{
    public PaymentRepository(PaymentDbContext context) : base(context)
    {
        
    }
}