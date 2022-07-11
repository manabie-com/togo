using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Utilities.Exceptions
{
    public class MyTodoException : Exception
    {
        public MyTodoException()
        {
        }

        public MyTodoException(string message)
            : base(message)
        {
        }

        public MyTodoException(string message, Exception inner)
            : base(message, inner)
        {
        }
    }
}
