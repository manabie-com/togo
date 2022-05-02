using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Togo.Core.Base
{
    public class ResponseBase<T>
    {
        public ResponseBase()
        {
            Message = ErrorCodeMessage.Success.Value;
            Code = ErrorCodeMessage.Success.Key;
        }

        public int Code { get; set; }
        public string Message { get; set; }

        public string ErrorDetail { get; set; }

        public bool IsSuccessful => Code == ErrorCodeMessage.Success.Key;

        public T Data { get; set; }
    }

    public class ResponseBase : ResponseBase<object>
    {
    }
}
