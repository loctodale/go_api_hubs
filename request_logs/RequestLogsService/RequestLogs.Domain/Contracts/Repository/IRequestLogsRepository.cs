using RequestLogs.Data.Entities;

namespace Payment.Data.Context;

public interface IRequestLogsRepository : IGenericRepository<TblRequestLog>
{
}