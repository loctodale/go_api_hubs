using Grpc.Core;

namespace PaymentService.Services;

public class PaymentService : global::PaymentService.PaymentService.PaymentServiceBase
{
    public PaymentService()
    {
    }

    public override Task<CreateWalletResponse> CreateNewWallet(CreateWalletRequest request, ServerCallContext context)
    {
        return base.CreateNewWallet(request, context);
    }

    public override Task<AddToWalletResponse> AddToWallet(AddToWalletRequest request, ServerCallContext context)
    {
        return base.AddToWallet(request, context);
    }
}