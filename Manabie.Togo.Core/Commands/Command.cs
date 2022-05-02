using FluentValidation.Results;
using Manabie.Togo.Core.Events;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Togo.Core.Commands
{
    public abstract class Command : Message
    {
        public int Id { get; set; }

        public string Name { get; set; }

        public ValidationResult ValidationResult { get; set; }
    }
}
