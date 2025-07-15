using Grpc.Core;
using Grpc.Net.Client;
using Payment.Data.Context;
using Pb;
using RequestLogs.Data.Entities;
using RequestLogs.Domain.Contracts.UnitOfWork;

namespace RequestLogsService.Services;

public class RequestLogsService : global::RequestLogsService.RequestLogsService.RequestLogsServiceBase
{
    private readonly IRequestLogsRepository _requestLogsRepository;
    private readonly IUnitOfWork _unitOfWork;
    public RequestLogsService(IUnitOfWork unitOfWork)
    {
        _requestLogsRepository = unitOfWork.RequestLogs;
        _unitOfWork = unitOfWork;
    }

    public override async Task<CreateRequestLogsResponse> CreateRequestLogs(CreateRequestLogsRequest request, ServerCallContext context)
    {
        var apisChannel = GrpcChannel.ForAddress("http://apis-service:8080");
        var apisKeyClient = new ApisKeyService.ApisKeyServiceClient(apisChannel);
        var apisKeyIsExisted = apisKeyClient.CheckApisKeyIsExisted(new CheckApisKeyIsExistedRequest
        {
            Id = request.ApiKeyId
        });
        if (apisKeyIsExisted.BaseResponse.Code != 200 || !apisKeyIsExisted.IsExisted)
        {
            return new CreateRequestLogsResponse
            {
                BaseResponse = new BaseResponse
                {
                    Code = apisKeyIsExisted.BaseResponse.Code,
                    Message = apisKeyIsExisted.BaseResponse.Message
                }
            };
        }
        await _requestLogsRepository.AddAsync(new TblRequestLog
        {
            Cost = request.Cost,
            Endpoint = request.Endpoint,
            StatusCode = request.StatusCode,
            ApiKeyId = Guid.Parse(request.ApiKeyId),
            RequestTime = DateTime.UtcNow,
        });

        var result = await _unitOfWork.SaveChangesAsync();
        if (!result)
        {
            return new CreateRequestLogsResponse
            {
                BaseResponse = new BaseResponse
                {
                    Code = 500,
                    Message = "Save failed",
                }
            };
        }

        return new CreateRequestLogsResponse
        {
            BaseResponse = new BaseResponse
            {
                Code = 200,
                Message = "Success",
            }
        };
    }
}