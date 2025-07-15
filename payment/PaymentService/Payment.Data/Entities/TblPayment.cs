using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using Payment.Data.Enums;

namespace Payment.Data.Entities;

public partial class TblPayment
{
    public Guid Id { get; set; }
    
    public Guid UserId { get; set; }
    public int Amount { get; set; }

    public string Method { get; set; } = null!;

    public string? ReferenceId { get; set; }

    public PaymentType PaymentType { get; set; }

    public DateTime? CreatedAt { get; set; }
}
