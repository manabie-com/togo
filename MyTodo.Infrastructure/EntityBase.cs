using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Infrastructure
{
    public abstract class EntityBase<T>
    {
        public T Id { get; set; }
        public bool IsTransient()
        {
            return Id.Equals(default(T));
        }
    }
}
