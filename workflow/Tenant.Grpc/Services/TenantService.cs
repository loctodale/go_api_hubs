using Grpc.Core;
using Tenant.Grpc.Context;
using Tenant.Grpc.Entities;
using TenantService;

namespace Tenant.Grpc.Services;

public class TenantService : TenantServiceGrpc.TenantServiceGrpcBase
{
    private readonly TenantContext _context;

    public TenantService(TenantContext context) {
        _context = context;
    }

    public override async Task<ResponseCreateTenant> CreateNewTenantId(RequestCreateTenant request, ServerCallContext context)
    {
        await _context.Tenants.AddAsync(new Entities.Tenant
        {
            Id = Guid.NewGuid().ToString(),
            Name = request.Name,
            TenantId = Guid.NewGuid().ToString(),
            Configuration = request.UserId
        });
        
        await _context.SaveChangesAsync();

        return new ResponseCreateTenant
        {
            BaseResponse = new BaseResponse
            {
                Message = "Success",
                Code = 200
            }
        };
    }
}