using Microsoft.Data.Sqlite;
using Newtonsoft.Json;
using ManabieTodo.Utils;

namespace ManabieTodo.Services
{
    public class DatabaseService : IDatabaseService
    {
        private SqliteConnection sqlite { get; }

        public DatabaseService(string connStr)
        {
            sqlite = new SqliteConnection(connStr);
            sqlite.Open();
        }

        public IEnumerable<T> ExecuteReader<T>(string cmd, IDictionary<string, object> args)
        {
            SqliteCommand command = sqlite.CreateCommand();
            command.CommandText = cmd;
            if (args != null)
            {
                foreach (var key in args.Keys)
                {
                    command.Parameters.AddWithValue($"${key}", args.GetString(key));
                }
            }

            using (SqliteDataReader reader = command.ExecuteReader())
            {
                while (reader.Read())
                {
                    if (GetBoolByColumn(reader, "IS_ACTIVE"))
                    {
                        yield return ConvertSqlite<T>(reader);
                    }
                }
            }
        }

        public async IAsyncEnumerable<T> ExecuteReaderAsync<T>(string cmd, IDictionary<string, object> args)
        {
            SqliteCommand command = sqlite.CreateCommand();
            command.CommandText = cmd;
            if (args != null)
            {
                foreach (var key in args.Keys)
                {
                    command.Parameters.AddWithValue($"${key}", args.GetString(key));
                }
            }

            using (SqliteDataReader reader = await command.ExecuteReaderAsync())
            {
                while (reader.Read())
                {
                    if (GetBoolByColumn(reader, "IS_ACTIVE"))
                    {
                        yield return ConvertSqlite<T>(reader);
                    }
                }
            }
        }

        public T? ExecuteObjectReader<T>(string cmd, IDictionary<string, object> args)
        {
            T? result = default(T);

            SqliteCommand command = sqlite.CreateCommand();
            command.CommandText = cmd;
            if (args != null)
            {
                foreach (var key in args.Keys)
                {
                    command.Parameters.AddWithValue($"${key}", args.GetString(key));
                }
            }

            using (SqliteDataReader reader = command.ExecuteReader())
            {
                if (reader.Read())
                {
                    if (GetBoolByColumn(reader, "IS_ACTIVE"))
                    {
                        result = ConvertSqlite<T>(reader);
                    }
                }
            }

            return result;
        }

        public async Task<T?> ExecuteObjectReaderAsync<T>(string cmd, IDictionary<string, object> args)
        {
            T? result = default(T);

            SqliteCommand command = sqlite.CreateCommand();
            command.CommandText = cmd;
            if (args != null)
            {
                foreach (var key in args.Keys)
                {
                    command.Parameters.AddWithValue($"${key}", args.GetString(key));
                }
            }

            using (SqliteDataReader reader = await command.ExecuteReaderAsync())
            {
                if (reader.Read())
                {
                    if (GetBoolByColumn(reader, "IS_ACTIVE"))
                    {
                        result = ConvertSqlite<T>(reader);
                    }
                }
            }

            return result;
        }

        public int ExecuteNonQuery(string cmd, IDictionary<string, object> args)
        {
            SqliteCommand command = sqlite.CreateCommand();
            command.CommandText = cmd;
            if (args != null)
            {
                foreach (var key in args.Keys)
                {
                    command.Parameters.AddWithValue($"${key}", args.GetString(key));
                }
            }

            return command.ExecuteNonQuery();
        }

        public async Task<int> ExecuteNonQueryAsync(string cmd, IDictionary<string, object> args)
        {
            SqliteCommand command = sqlite.CreateCommand();
            command.CommandText = cmd;
            if (args != null)
            {
                foreach (var key in args.Keys)
                {
                    command.Parameters.AddWithValue($"${key}", args.GetString(key));
                }
            }

            return await command.ExecuteNonQueryAsync();
        }

        private bool GetBoolByColumn(SqliteDataReader reader, string columnName)
        {
            IDictionary<string, object> data = Enumerable
                            .Range(0, reader.FieldCount)
                            .ToDictionary(reader.GetName, reader.GetValue);

            return data.GetBoolean(columnName, true);
        }

        private T? ConvertSqlite<T>(SqliteDataReader reader)
        {
            IDictionary<string, object> data = Enumerable
                            .Range(0, reader.FieldCount)
                            .ToDictionary(reader.GetName, reader.GetValue);

            string json = JsonConvert.SerializeObject(data);

            return JsonConvert.DeserializeObject<T>(json);
        }
    }
}