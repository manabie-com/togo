using System.ComponentModel.DataAnnotations;

namespace WebApi.Requests
{
    public class CreateTaskRequest
    {
        [Required(ErrorMessage = "The Content field is required")]
        public string Content { get; set; }
    }
}
