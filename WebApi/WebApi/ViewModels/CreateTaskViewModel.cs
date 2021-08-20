using System;

namespace WebApi.ViewModels
{
    public class CreateTaskViewModel
    {
        public Guid Id { get; set; }
        public string Content { get; set; }
    }
}
