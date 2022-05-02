using AutoMapper;
using Manabie.Togo.Core.Base;
using Manabie.Togo.Core.Bus;
using Manabie.Togo.Core.Handler;
using Manabie.Togo.Domain.Events.UserTask.Create;
using Manabie.Togo.JsonRepository.UserTask;
using MediatR;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace Manabie.Togo.Domain.Commands.UserTask.Create
{
	public class CreateUserTaskCommandHandler : CommandHandler, IRequestHandler<CreatedUserTaskCommand, CreateUserTaskResponse>
	{
		private IMediatorHandler _bus;
		private readonly IUserTaskJsonRepository _userTaskJsonRepository;
		private readonly IMapper _mapper;
        public CreateUserTaskCommandHandler(IMediatorHandler bus, IUserTaskJsonRepository userTaskJsonRepository, IMapper mapper) : base(bus)
        {
            _bus = bus;
            _userTaskJsonRepository = userTaskJsonRepository;
            _mapper = mapper;
        }

        /// <summary>
        /// Handle command
        /// </summary>
        /// <param name="request"></param>
        /// <param name="cancellationToken"></param>
        /// <returns></returns>
        public async Task<CreateUserTaskResponse> Handle(CreatedUserTaskCommand request, CancellationToken cancellationToken)
		{
			var response = new CreateUserTaskResponse();
			try
			{
				_userTaskJsonRepository.Add(request.UserTaskEntity);
				var result = await _userTaskJsonRepository.SaveChange();
			}
			catch (Exception ex)
			{
				response.Code = ErrorCodeMessage.IncorrectFunction.Key;
				response.Message = ErrorCodeMessage.IncorrectFunction.Value;
				response.ErrorDetail = ex.ToString();
			}
			return response;
		}
	}
}
