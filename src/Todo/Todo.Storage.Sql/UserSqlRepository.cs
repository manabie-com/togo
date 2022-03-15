using System.Data;
using System.Data.SqlClient;
using Dapper;
using Todo.Storage.Contract;
using Todo.Storage.Contract.Interfaces;

namespace Todo.Storage.Sql;

public class UserSqlRepository : IUserRepository
{
    private readonly string _connectionString;

    private readonly ISerializer _serializer;

    public UserSqlRepository(string connectionString, ISerializer serializer)
    {
        if (string.IsNullOrEmpty(connectionString))
        {
            throw new ArgumentException("Connection string must not be null or empty.", nameof(connectionString));
        }
        
        _connectionString = connectionString;
        _serializer = serializer ?? throw new ArgumentNullException(nameof(serializer));
    }

    public async Task<UserResponse> AddAsync(UserRequest request)
    {
        var sql = "dbo.AddUser";

        try
        {
            // Open an SQL connection
            using (var connection = new SqlConnection(_connectionString))
            {
                connection.Open();

                // Run the stored procedure
                var result = await connection.QueryFirstOrDefaultAsync(
                    sql,
                    new
                    {
                        request.FirstName,
                        request.LastName,
                        request.DailyTaskLimit
                    },
                    commandType: CommandType.StoredProcedure);

                if (result == null)
                {
                    throw new Exception($"Unable to read result from command: {sql}");
                }

                // Create and return the user response based on query results
                return new UserResponse
                {
                    Id = result.ID,
                    FirstName = result.FirstName,
                    LastName = result.LastName,
                    DailyTaskLimit = result.DailyTaskLimit
                };
            }
        }
        catch (SqlException)
        {
            throw;
        }
        catch (Exception)
        {
            throw;
        }
    }

    public async Task<UserResponse> GetAsync(long userID)
    {
        var getUserSql = "dbo.GetUser";

        try
        {
            // Open an SQL connection
            using (var connection = new SqlConnection(_connectionString))
            {
                connection.Open();

                // Run the stored procedure to get the User
                var result = await connection.QueryFirstOrDefaultAsync(
                    getUserSql,
                    new { userID },
                    commandType: CommandType.StoredProcedure);

                if (result == null)
                {
                    throw new Exception($"Unable to read result from command: {getUserSql}");
                }

                var response = new UserResponse
                {
                    Id = result.ID,
                    FirstName = result.FirstName,
                    LastName = result.LastName,
                    Todos = "",
                    DailyTaskLimit = result.DailyTaskLimit
                };

                var getTodosSql = "dbo.GetUserTodoItems";
                
                // Run the stored procedure to get the user's Todos
                var results = await connection.QueryAsync(
                    getTodosSql,
                    new { userID },
                    commandType: CommandType.StoredProcedure);

                if (results == null)
                {
                    throw new Exception($"Unable to read result from command: {getTodosSql}");
                }

                var json = _serializer.Serialize(results);

                if (json == "[]")
                {
                    return response;
                }

                response.Todos = json;

                return response;
            }
        }
        catch (SqlException)
        {
            throw;
        }
        catch (Exception)
        {
            throw;
        }
    }

    public async Task<long> DeleteAsync(long userID)
    {
        var sql = "dbo.DeleteUser";

        try
        {
            // Open an SQL connection
            using (var connection = new SqlConnection(_connectionString))
            {
                connection.Open();

                // Run the stored procedure
                var result = await connection.ExecuteAsync(
                    sql,
                    new { userID },
                    commandType: CommandType.StoredProcedure);

                return userID;
            }
        }
        catch (SqlException)
        {
            throw;
        }
        catch (Exception)
        {
            throw;
        }
    }
}