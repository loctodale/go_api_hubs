using System;
using System.Collections.Generic;
using Microsoft.EntityFrameworkCore;
using Tenant.Grpc.Entities;

namespace Tenant.Grpc.Context;

public partial class TenantContext : DbContext
{
    public TenantContext()
    {
    }

    public TenantContext(DbContextOptions<TenantContext> options)
        : base(options)
    {
    }

    public virtual DbSet<Entities.Tenant> Tenants { get; set; }
    private string GetConnectionString()
    {
        IConfiguration config = new ConfigurationBuilder()
            .SetBasePath(Directory.GetCurrentDirectory())
            .AddJsonFile("appsettings.json", true, true)
            .Build();
        var strConn = config.GetConnectionString("Postgresql");

        return strConn;
    }
    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        => optionsBuilder.UseNpgsql(GetConnectionString());

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.HasPostgresExtension("uuid-ossp");

        modelBuilder.Entity<Entities.Tenant>(entity =>
        {
            entity.ToTable("Tenants", "Elsa");

            entity.HasIndex(e => e.Name, "IX_Tenant_Name");

            entity.HasIndex(e => e.TenantId, "IX_Tenant_TenantId");
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
}
