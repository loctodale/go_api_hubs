using Grpc.Core;
using Payment.Data.Entities;
using Payment.Data.Enums;
using Payment.Domain.Contracts.Repository;
using Payment.Domain.Contracts.UnitOfWork;

namespace PaymentService.Services;

public class PaymentService : global::PaymentService.PaymentService.PaymentServiceBase
{
    private readonly IUnitOfWork _unitOfWork;
    private readonly IPaymentRepository _paymentRepository;
    public PaymentService(IUnitOfWork unitOfWork)
    {
        _unitOfWork = unitOfWork;
        _paymentRepository = unitOfWork.Payment;
    }

    public override async Task<CreateWalletResponse> CreateNewWallet(CreateWalletRequest request, ServerCallContext context)
    {
        // log request
        // Xử lí tạo tài khoản
        await _paymentRepository.AddAsync(new TblPayment
        {
            Amount = request.Amount,
            PaymentType = (PaymentType)request.Type,
            Method = request.Method,
            ReferenceId = request.ReferenceId,
            UserId = Guid.Parse(request.Userid),
        });
        var result = await _unitOfWork.SaveChangesAsync();

        if (!result)
        {
            return new CreateWalletResponse
            {
                BaseResponse = new BaseResponse
                {
                    Message = "Failed to create new wallet",
                    Code = 500
                }
            };
        }
        return new CreateWalletResponse
        {
            BaseResponse = new BaseResponse
            {
                Message = "Success",
                Code = 200
            },
        };
    }

    public override async Task<AddToWalletResponse> AddToWallet(AddToWalletRequest request, ServerCallContext context)
    {
        // Xử lí thêm tiền vào tài khoản
        var wallet = await _paymentRepository.GetByIdAsync(Guid.Parse(request.Id));
        if (wallet == null)
        {
            return new AddToWalletResponse
            {
                BaseResponse = new BaseResponse
                {
                    Message = "Wallet not found",
                    Code = 500
                },
            };
        }

        if (wallet.UserId != Guid.Parse(request.UserId))
        {
            return new AddToWalletResponse
            {
                BaseResponse = new BaseResponse
                {
                    Message = "Wallet invalid",
                    Code = 500
                },
            };
        }
        wallet.Amount += request.Amount;
        
        // logs lại tiền đã + | - vào tài khoản
         _paymentRepository.Update(wallet);
         var result = await _unitOfWork.SaveChangesAsync();
         if (!result)
         {
             return new AddToWalletResponse
             {
                 BaseResponse = new BaseResponse
                 {
                     Message = "Failed",
                     Code = 500
                 }
             };
         }

         return new AddToWalletResponse
         {
             BaseResponse = new BaseResponse
             {
                 Message = "Success",
                 Code = 200
             }
         };
    }
}