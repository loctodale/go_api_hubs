using Microsoft.Extensions.DependencyInjection;
using Payment.Data.Context;
using Payment.Domain.Contracts.Repository;
using Payment.Domain.Contracts.UnitOfWork;
using Payment.Repository.Repository;

namespace Payment.Repository.UnitOfWork;

public class UnitOfWork : IUnitOfWork
{
    protected PaymentDbContext _context;
    private readonly IServiceProvider _serviceProvider;

    public UnitOfWork(PaymentDbContext context, IServiceProvider serviceProvider)
    {
        _context = context;
        _serviceProvider = serviceProvider;
    }
    
    public async Task<bool> SaveChangesAsync(CancellationToken cancellationToken = default)
    {
        var result = await _context.SaveChangesAsync(cancellationToken) > 0;
        
        return result;
    }
    public IPaymentRepository Payment => _serviceProvider.GetService<IPaymentRepository>();

}