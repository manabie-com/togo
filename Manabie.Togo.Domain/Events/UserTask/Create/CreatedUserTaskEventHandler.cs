using Manabie.Togo.RedisRepository.Interface;
using MediatR;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace Manabie.Togo.Domain.Events.UserTask.Create
{
    public class CreatedUserTaskEventHandler : INotificationHandler<CreatedUserTaskEvent>
    {
        private IUserTaskRepositoryRedis _userTaskRepositoryRedis;
        public CreatedUserTaskEventHandler(IUserTaskRepositoryRedis userTaskRepositoryRedis)
        {
            _userTaskRepositoryRedis = userTaskRepositoryRedis;
        }

        public async Task Handle(CreatedUserTaskEvent notification, CancellationToken cancellationToken)
        {
            // Save to redis
            await _userTaskRepositoryRedis.SaveAsync(notification.UserTaskEntity);
        }
    }
}
