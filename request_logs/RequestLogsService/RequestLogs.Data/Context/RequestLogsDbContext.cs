using System;
using System.Collections.Generic;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using RequestLogs.Data.Entities;

namespace RequestLogs.Data.Context;

public partial class RequestLogsDbContext : DbContext
{
    public RequestLogsDbContext()
    {
    }

    public RequestLogsDbContext(DbContextOptions<RequestLogsDbContext> options)
        : base(options)
    {
    }

    public virtual DbSet<TblRequestLog> TblRequestLogs { get; set; }

    private string GetConnectionString()
        {
            IConfiguration config = new ConfigurationBuilder()
                .SetBasePath(Directory.GetCurrentDirectory())
                .AddJsonFile("appsettings.json", true, true)
                .Build();
            var strConn = config.GetConnectionString("DefaultConnection");
    
            return strConn;
        }
    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
#warning To protect potentially sensitive information in your connection string, you should move it out of source code. You can avoid scaffolding the connection string by using the Name= syntax to read it from configuration - see https://go.microsoft.com/fwlink/?linkid=2131148. For more guidance on storing connection strings, see https://go.microsoft.com/fwlink/?LinkId=723263.
        => optionsBuilder.UseNpgsql(GetConnectionString());

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.HasPostgresExtension("uuid-ossp");

        modelBuilder.Entity<TblRequestLog>(entity =>
        {
            entity.HasKey(e => e.Id).HasName("tbl_request_logs_pkey");

            entity.ToTable("tbl_request_logs");

            entity.Property(e => e.Id)
                .HasDefaultValueSql("uuid_generate_v4()")
                .HasColumnName("id");
            entity.Property(e => e.ApiKeyId).HasColumnName("api_key_id");
            entity.Property(e => e.Cost).HasColumnName("cost");
            entity.Property(e => e.Endpoint).HasColumnName("endpoint");
            entity.Property(e => e.RequestTime)
                .HasDefaultValueSql("now()")
                .HasColumnType("timestamp without time zone")
                .HasColumnName("request_time");
            entity.Property(e => e.StatusCode).HasColumnName("status_code");
        });

        OnModelCreatingPartial(modelBuilder);
    }

    partial void OnModelCreatingPartial(ModelBuilder modelBuilder);
}
