using Manabie.Togo.Core.Commands;
using Manabie.Togo.Core.Events;
using System.Threading.Tasks;

namespace Manabie.Togo.Core.Bus
{
    public interface IMediatorHandler
    {
        Task<TResult> SendCommand<TRequest, TResult>(TRequest command) where TRequest : ICommand<TResult>;
        Task RaiseEvent<T>(T @event) where T : Event;
    }
}
