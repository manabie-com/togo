using FluentValidation;
using ToDo.Api.Requests;

namespace ToDo.Api.Validators
{
    public class CreateToDoValidator : AbstractValidator<CreateTodoRequest>
    {
        public CreateToDoValidator()
        {
            RuleFor(request => request.UserId).NotNull().NotEmpty();
            RuleFor(request => request.TodoName).NotNull().NotEmpty()
                .MaximumLength(100);
            RuleFor(request => request.TodoDescription).NotNull().NotEmpty()
                .MaximumLength(1000);
        }
    }
}
