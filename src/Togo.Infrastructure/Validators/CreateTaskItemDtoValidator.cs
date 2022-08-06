using FluentValidation;
using Togo.Core.AppServices.TaskItems.Dtos;

namespace Togo.Infrastructure.Validators;

public class CreateTaskItemDtoValidator : AbstractValidator<CreateTaskItemDto>
{
    public CreateTaskItemDtoValidator()
    {
        RuleFor(task => task.Title)
            .NotNull()
            .NotEmpty();
    }
}
