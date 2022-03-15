using System.Data;
using System.Data.SqlClient;
using Dapper;
using Todo.Storage.Contract;
using Todo.Storage.Contract.Interfaces;

namespace Todo.Storage.Sql;

public class TodoSqlRepository : ITodoRepository
{
    private readonly string _connectionString;

    private readonly ISerializer _serializer;

    public TodoSqlRepository(string connectionString, ISerializer serializer)
    {
        if (string.IsNullOrEmpty(connectionString))
        {
            throw new ArgumentException("Connection string must not be null or empty.", nameof(connectionString));
        }
        
        _connectionString = connectionString;
        _serializer = serializer ?? throw new ArgumentNullException(nameof(serializer));
    }

    public async Task<TodoResponse> AddAsync(TodoRequest request)
    {
        var sql = "dbo.AddTodo";

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
                        request.Name,
                        request.Description,
                        request.DateCreatedUTC,
                        request.DateModifiedUTC,
                        request.UserId
                    },
                    commandType: CommandType.StoredProcedure);

                if (result == null)
                {
                    throw new Exception($"Unable to read result from command: {sql}");
                }

                return new TodoResponse
                {
                    Id = result.ID,
                    Name = result.Name,
                    Description = result.Description,
                    DateCreatedUTC = result.DateCreatedUTC,
                    DateModifiedUTC = result.DateModifiedUTC,
                    UserId = result.UserID
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
    
    public async Task<long> DeleteAsync(long todoID)
    {
        var sql = "dbo.DeleteTodo";

        try
        {
            // Open an SQL connection
            using (var connection = new SqlConnection(_connectionString))
            {
                connection.Open();

                // Run the stored procedure
                var result = await connection.ExecuteAsync(
                    sql,
                    new { todoID },
                    commandType: CommandType.StoredProcedure);

                return todoID;
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