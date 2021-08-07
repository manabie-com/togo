using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace togo.Api
{
    public class ApiResponse<T> where T : class
    {
        public T Data { get; set; }

        public static implicit operator ApiResponse<T>(T data)
        {
            return new ApiResponse<T> { Data = data };
        }
    }
}
