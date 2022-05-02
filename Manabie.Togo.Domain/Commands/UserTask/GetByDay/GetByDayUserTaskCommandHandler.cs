using AutoMapper;
using Manabie.Togo.Core.Base;
using Manabie.Togo.Core.Bus;
using Manabie.Togo.Core.Handler;
using Manabie.Togo.JsonRepository.UserTask;
using Manabie.Togo.RedisRepository.Interface;
using MediatR;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace Manabie.Togo.Domain.Commands.UserTask.GetByDay
{
	public class GetByDayUserTaskCommandHandler : CommandHandler, IRequestHandler<GetByDayUserTaskCommand, GetByDayUserTaskResponse>
	{
		private IMediatorHandler _bus;
		private IUserTaskRepositoryRedis _userTaskRepositoryRedis;
		private readonly IUserTaskJsonRepository _userTaskJsonRepository;
		private readonly IMapper _mapper;
		public GetByDayUserTaskCommandHandler(IMediatorHandler bus, IUserTaskJsonRepository userTaskJsonRepository, IMapper mapper, IUserTaskRepositoryRedis userTaskRepositoryRedis) : base(bus)
		{
			_bus = bus;
			_userTaskJsonRepository = userTaskJsonRepository;
			_mapper = mapper;
			_userTaskRepositoryRedis = userTaskRepositoryRedis;
		}

		/// <summary>
		/// Handle command
		/// </summary>
		/// <param name="request"></param>
		/// <param name="cancellationToken"></param>
		/// <returns></returns>
		public async Task<GetByDayUserTaskResponse> Handle(GetByDayUserTaskCommand request, CancellationToken cancellationToken)
		{
			var response = new GetByDayUserTaskResponse();
			try
			{
				response.Data = await _userTaskRepositoryRedis.GetAllByDay(request.GetUserTaskDto.UserId, request.GetUserTaskDto.TaskDate);
				return response;
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
