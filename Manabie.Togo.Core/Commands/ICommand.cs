using MediatR;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Togo.Core.Commands
{
    public interface ICommand<TResult> : IRequest<TResult>
    {

    }
}
