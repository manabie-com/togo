using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace togo.Service.Interface
{
    public interface ICurrentHttpContext
    {
        public string GetCurrentUserId();
    }
}
