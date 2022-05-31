using FluentValidation;
using ToDo.Api.Requests;

namespace ToDo.Api.Validators
{
    public class CreateUserValidator : AbstractValidator<CreateUserRequest>
    {
        public CreateUserValidator()
        {
            RuleFor(request => request.FullName).NotNull().NotEmpty()
                .MaximumLength(32);
            RuleFor(request => request.DailyTaskLimit).NotNull().NotEmpty()
                .LessThan(100);
        }
    }
}
