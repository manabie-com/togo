using ManabieTodo.Models;
using ManabieTodo.Services;
using Newtonsoft.Json;
using Xunit;

namespace ManabieTodo.Tests
{
    public class DatabaseServiceTest
    {
        private IDatabaseService _dbService { get; }

        private readonly TodoModel TODO_TEST_DATA_1 = new TodoModel
        {
            Id = 0,
            Do = "Do a job 1",
            IsActive = true,
            IsComplete = false,
            CreatedDate = new DateTime(2019, 1, 1),
        };

        public DatabaseServiceTest()
        {
            _dbService = new DatabaseService("Data Source=database.db;");
        }

        [Theory]
        [InlineData("Do a task 1", 999)]
        [InlineData("Do a task 2", 3)]
        [InlineData("Do a task 3", 999)]
        [InlineData("Do a task 4", 999)]
        [InlineData("Do a task 5", 999)]
        public void InsertTest(string todo, int assignee)
        {
            TodoModel model = new TodoModel
            {
                Do = todo,
                Assignee = assignee
            };

            string cmd = InsertCommand();

            string json = JsonConvert.SerializeObject(model);

            IDictionary<string, object> parameters = JsonConvert.DeserializeObject<IDictionary<string, object>>(json);

            InsertSupportModel result =
                _dbService.ExecuteObjectReader<InsertSupportModel>(cmd, parameters);


            Assert.NotEqual(0, result.seq);
        }

        [Theory]
        [InlineData(999)]
        public void DeleteAllTaskByUserTest(int assignee)
        {
            string cmd = DeleteLogicCommand();

            IDictionary<string, object> parameters = new Dictionary<string, object>
            {
                ["ID"] = assignee,
                ["IS_ACTIVE"] = false
            };

            int result = _dbService.ExecuteNonQuery(cmd, parameters);


            Assert.NotEqual(4, result);
        }


        private string InsertCommand()
        {
            return @$"
            INSERT INTO TODO (DO, IS_ACTIVE, IS_COMPLETE, CREATED_DATE)
            VALUES ($DO, $IS_ACTIVE, $IS_COMPLETE, $CREATED_DATE);

            SELECT SEQ FROM SQLITE_SEQUENCE WHERE NAME = 'TODO';
            ";
        }

        private string DeleteLogicCommand()
        {
            return @"
            UPDATE TODO SET IS_ACTIVE = $IS_ACTIVE
            WHERE ID = $ID
            ";
        }
    }
}