using System;
using Bogus;
using TogoService.API.Dto;
using TogoService.API.Model;

namespace TogoService.UnitTest.Helper
{
    public class FakeData
    {
        private static int GenerateRandomArrLength()
        {
            Random random = new Random();
            return random.Next(1, 11);
        }

        public static TaskRequest[] GenerateTaskRequests(int numberOfItems = 0)
        {
            if (numberOfItems == 0)
            {
                numberOfItems = GenerateRandomArrLength();
            }

            var fakerTaskRequest = new Faker<TaskRequest>()
                .RuleFor(s => s.Name, f => f.Name.JobTitle())
                .RuleFor(s => s.Description, f => f.Name.JobDescriptor());

            return fakerTaskRequest.Generate(numberOfItems).ToArray();
        }

        public static NewTaskRequest GenerateNewTaskRequest(int numberOfTasks = 0)
        {
            var fakerNewTaskRequest = new Faker<NewTaskRequest>()
                .RuleFor(s => s.Date, f => DateTime.UtcNow)
                .RuleFor(s => s.Tasks, _ => GenerateTaskRequests(numberOfTasks));

            return fakerNewTaskRequest.Generate();
        }

        public static User GenerateUser(uint maxDailyTask)
        {
            var fakerUser = new Faker<User>()
                .RuleFor(s => s.Id, _ => Guid.NewGuid())
                .RuleFor(s => s.MaxDailyTasks, _ => maxDailyTask);

            return fakerUser.Generate();
        }

        public static TodoTask[] GenerateTodoTasks(int numberOfItems = 0)
        {
            if (numberOfItems == 0)
            {
                numberOfItems = GenerateRandomArrLength();
            }

            var fakerTodoTask = new Faker<TodoTask>()
                .RuleFor(s => s.Name, f => f.Name.JobTitle())
                .RuleFor(s => s.Description, f => f.Name.JobDescriptor());

            return fakerTodoTask.Generate(numberOfItems).ToArray();
        }
    }
}