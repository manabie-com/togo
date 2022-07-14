namespace Manabie.Testing.Domain.Common
{
    public class BaseEntity<T>
    {
        public T Id { get; set; }
    }

    public class BaseEntity : BaseEntity<int>
    {
    }
}
