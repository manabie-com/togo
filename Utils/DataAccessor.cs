namespace ManabieTodo.Utils
{
    public static class DataAccessor
    {
        public static string GetString(this IDictionary<string, object> dic, string key)
        {
            if (dic.ContainsKey(key))
            {
                return dic[key]?.ToString() ?? "";
            }

            return "";
        }

        public static int GetInt(this IDictionary<string, object> dic, string key, int defaultValue = 0)
        {
            if (dic.ContainsKey(key))
            {
                return int.Parse(dic[key]?.ToString() ?? defaultValue.ToString());
            }

            return defaultValue;
        }

        public static int? GetNullableInt(this IDictionary<string, object> dic, string key)
        {
            int value = 0;

            if (dic.ContainsKey(key))
            {
                if (int.TryParse(dic[key]?.ToString(), out value))
                {
                    return value;
                }
            }

            return null;
        }

        public static DateTime GetDateTime(this IDictionary<string, object> dic, string key, DateTime defaultValue = new DateTime())
        {
            if (dic.ContainsKey(key))
            {
                return DateTime.Parse(dic[key]?.ToString() ?? defaultValue.ToString());
            }

            return defaultValue;
        }

        public static DateTime? GetNullableDateTime(this IDictionary<string, object> dic, string key)
        {
            DateTime value = new DateTime();

            if (dic.ContainsKey(key))
            {
                if (DateTime.TryParse(dic[key]?.ToString(), out value))
                {
                    return value;
                }
            }

            return null;
        }

        public static bool GetBoolean(this IDictionary<string, object> dic, string key, bool defaultValue = false)
        {
            if (dic.ContainsKey(key))
            {
                return bool.Parse(dic[key]?.ToString() ?? defaultValue.ToString());
            }

            return defaultValue;
        }

        public static bool? GetNullableBoolean(this IDictionary<string, object> dic, string key)
        {
            bool value = false;

            if (dic.ContainsKey(key))
            {
                if (bool.TryParse(dic[key]?.ToString(), out value))
                {
                    return value;
                }
            }

            return null;
        }
    }
}