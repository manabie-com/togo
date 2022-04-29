using ManabieTodo.Models;
using ManabieTodo.Utils;
using Newtonsoft.Json;

namespace ManabieTodo.Services
{
    public class TodoService : ITodoService
    {
        private IDatabaseService _dbService { get; }
        private IFakeCache _fakeCache { get; }

        public TodoService(IDatabaseService dbService)
        {
            _dbService = dbService;
            _fakeCache = CacheFactory.FakeCache;
        }

        public bool ToggleComplete(int id)
        {
            string cmd = CompleteTaskCommand();

            bool isCompleteState = GetCompleteState(id);

            IDictionary<string, object> parameters = new Dictionary<string, object>
            {
                ["ID"] = id,
                ["IS_COMPLETE"] = !isCompleteState
            };

            int result = _dbService.ExecuteNonQuery(cmd, parameters);

            if (result > 0)
            {
                return !isCompleteState;
            }

            return isCompleteState;
        }


        public bool Delete(int id)
        {
            string cmd = UpdateUnactiveCommand();

            IDictionary<string, object> parameters = new Dictionary<string, object>
            {
                ["ID"] = id,
            };

            int result = _dbService.ExecuteNonQuery(cmd, parameters);

            return result > 0;
        }

        public bool DeleteAll()
        {
            string cmd = UpdateUnactiveCommand();

            IDictionary<string, object> parameters = new Dictionary<string, object>();

            int result = _dbService.ExecuteNonQuery(cmd, parameters);

            return result > 0;
        }

        public TodoModel Get(int id)
        {
            string cmd = GetAllCommand();

            cmd += "WHERE T.ID = $ID";

            IDictionary<string, object> parameters = new Dictionary<string, object>();

            return _dbService.ExecuteObjectReader<TodoModel>(cmd, parameters);
        }

        public Task<TodoModel> GetAsync(int id)
        {
            string cmd = GetAllCommand();

            cmd += "WHERE T.ID = $ID";

            IDictionary<string, object> parameters = new Dictionary<string, object>
            {
                ["ID"] = id
            };

            return _dbService.ExecuteObjectReaderAsync<TodoModel>(cmd, parameters);
        }


        public IEnumerable<TodoModel> GetAll()
        {
            string cmd = GetAllCommand();

            IDictionary<string, object> parameters = new Dictionary<string, object>();

            return _dbService.ExecuteReader<TodoModel>(cmd, parameters);
        }

        public IAsyncEnumerable<TodoModel> GetAllAsync()
        {
            string cmd = GetAllCommand();

            IDictionary<string, object> parameters = new Dictionary<string, object>();

            return _dbService.ExecuteReaderAsync<TodoModel>(cmd, parameters);
        }

        public int? Insert(TodoModel model)
        {
            string cmd = InsertCommand();

            string json = JsonConvert.SerializeObject(model);
            IDictionary<string, object> parameters = JsonConvert.DeserializeObject<IDictionary<string, object>>(json);

            InsertSupportModel result =
                _dbService.ExecuteObjectReader<InsertSupportModel>(cmd, parameters);

            if (result != null)
            {
                return result.seq;
            }

            return null;
        }

        public bool Update(TodoModel model)
        {
            string cmd = UpdateTaskCommand();

            IDictionary<string, object> parameters = new Dictionary<string, object>
            {
                ["ID"] = model.Id,
                ["DO"] = model.Do
            };

            int result = _dbService.ExecuteNonQuery(cmd, parameters);

            return result > 0;
        }

        private string GetAllCommand()
        {
            return @"
            SELECT
            T.ID,
            T.DO,
            T.IS_ACTIVE,
            T.IS_COMPLETE,
            T.CREATED_DATE,
            T.UPDATED_DATE,
            T.ASSIGNEE
            FROM TODO AS T
            ";
        }

        private string InsertCommand()
        {
            return @$"
            INSERT INTO TODO (DO, IS_ACTIVE, IS_COMPLETE, CREATED_DATE, ASSIGNEE)
            VALUES ($DO, $IS_ACTIVE, $IS_COMPLETE, $CREATED_DATE, $ASSIGNEE);
            SELECT SEQ FROM SQLITE_SEQUENCE WHERE NAME = 'TODO';
            ";
        }

        private string GetCompleteStateCommand()
        {
            return @"
            SELECT
            T.IS_COMPLETE
            FROM TODO AS T
            WHERE 
            T.ID = $ID
            ";
        }

        private string CompleteTaskCommand()
        {
            return @"
            UPDATE TODO SET IS_COMPLETE = $IS_COMPLETE
            WHERE ID = $ID
            ";
        }

        private string UpdateTaskCommand()
        {
            return @"
            UPDATE TODO SET DO = $DO
            WHERE ID = $ID
            ";
        }

        private string UpdateUnactiveCommand()
        {
            return @"
            UPDATE TODO SET IS_ACTIVE = 'False'
            ";
        }

        private bool GetCompleteState(int id)
        {
            string cmd = GetCompleteStateCommand();

            IDictionary<string, object> parameters = new Dictionary<string, object>
            {
                ["ID"] = id,
            };

            TodoModel result = _dbService.ExecuteObjectReader<TodoModel>(cmd, parameters);

            return result.IsComplete;
        }
    }
}