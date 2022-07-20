using System;
using System.Collections.Generic;
using System.Text;

namespace MyTodo.Data.Enums
{
    public enum TodoItemStatus
    {
        /// <summary>
        /// New
        /// </summary>
        New,
        /// <summary>
        /// Pending
        /// </summary>
        Pending,
        /// <summary>
        /// Assigned
        /// </summary>
        Assigned,
        /// <summary>
        /// InProgress
        /// </summary>
        InProgress,
        /// <summary>
        /// Done
        /// </summary>
        Done,
        /// <summary>
        /// Closed
        /// </summary>
        Closed,
        /// <summary>
        /// Canceled
        /// </summary>
        Canceled,
        /// <summary>
        /// Removed
        /// </summary>
        Removed
    }
}
