namespace ManabieTodo.Services
{
    public interface IDatabaseService
    {
        IEnumerable<T> ExecuteReader<T>(string cmd, IDictionary<string, object> args);
        IAsyncEnumerable<T> ExecuteReaderAsync<T>(string cmd, IDictionary<string, object> args);
        T ExecuteObjectReader<T>(string cmd, IDictionary<string, object> args);
        Task<T> ExecuteObjectReaderAsync<T>(string cmd, IDictionary<string, object> args);
        int ExecuteNonQuery(string cmd, IDictionary<string, object> args);
        Task<int> ExecuteNonQueryAsync(string cmd, IDictionary<string, object> args);
    }
}