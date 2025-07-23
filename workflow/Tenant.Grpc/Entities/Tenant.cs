using System;
using System.Collections.Generic;

namespace Tenant.Grpc.Entities;

public partial class Tenant
{
    public string Id { get; set; } = null!;

    public string Name { get; set; } = null!;

    public string Configuration { get; set; } = null!;

    public string? TenantId { get; set; }
}
