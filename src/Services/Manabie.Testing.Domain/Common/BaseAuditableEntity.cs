namespace Manabie.Testing.Domain.Common
{
    public class BaseAuditableEntity<TKey> : BaseEntity<TKey>
    {
        public DateTime Created { get; set; }

        public string? CreatedBy { get; set; }

        public DateTime? LastModified { get; set; }

        public string? LastModifiedBy { get; set; }
    }
}
