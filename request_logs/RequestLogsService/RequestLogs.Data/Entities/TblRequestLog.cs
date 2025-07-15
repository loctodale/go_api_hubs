using System;
using System.Collections.Generic;

namespace RequestLogs.Data.Entities;

public partial class TblRequestLog
{
    public Guid Id { get; set; }

    public Guid ApiKeyId { get; set; }

    public DateTime? RequestTime { get; set; }

    public int StatusCode { get; set; }

    public string Endpoint { get; set; } = null!;

    public int? Cost { get; set; }
}
