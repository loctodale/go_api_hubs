using Payment.Data.Context;
using RequestLogs.Data.Context;
using RequestLogs.Data.Entities;

namespace RequestLogs.Repository.Repository;

public class RequestLogsRepository : GenericRepository<TblRequestLog>, IRequestLogsRepository
{
    public RequestLogsRepository(RequestLogsDbContext context) : base(context)
    {
        
    } 
}