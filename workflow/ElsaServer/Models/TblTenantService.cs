using System;
using System.Collections.Generic;

namespace ElsaServer.Models;

public partial class TblTenantService
{
    public Guid Id { get; set; }

    public Guid TenantId { get; set; }

    public Guid UserId { get; set; }

    public string Name { get; set; } = null!;

    public DateTime? CreatedAt { get; set; }
}
