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

    public virtual DbSet<TblTenantService> TblTenantServices { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
#warning To protect potentially sensitive information in your connection string, you should move it out of source code. You can avoid scaffolding the connection string by using the Name= syntax to read it from configuration - see https://go.microsoft.com/fwlink/?linkid=2131148. For more guidance on storing connection strings, see https://go.microsoft.com/fwlink/?LinkId=723263.
        => optionsBuilder.UseNpgsql("Host=ep-winter-unit-a12g657d-pooler.ap-southeast-1.aws.neon.tech; Database=tenant_service; Username=account_service_owner; Password=npg_OV47RfoAulnp; SSL Mode=VerifyFull; Channel Binding=Require;");

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.HasPostgresExtension("uuid-ossp");

        modelBuilder.Entity<TblTenantService>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("tbl_tenant_service_pkey");

            entity.ToTable("tbl_tenant_service");

            entity.Property(e => e.Id)
                .HasDefaultValueSql("uuid_generate_v4()")
                .HasColumnName("id");
            entity.Property(e => e.CreatedAt)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("created_at");
            entity.Property(e => e.Name).HasColumnName("name");
            entity.Property(e => e.TenantId).HasColumnName("tenantId");
            entity.Property(e => e.UserId).HasColumnName("user_id");
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
}
