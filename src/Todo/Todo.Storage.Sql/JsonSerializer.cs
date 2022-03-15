
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Todo.Storage.Contract.Interfaces;

namespace Todo.Domain;

public class JsonSerializer : ISerializer
{
    public string Serialize(object value)
    {
        if (value == null)
        {
            throw new ArgumentNullException(nameof(value));
        }

        return JsonConvert.SerializeObject(value, Formatting.Indented, OnGetJsonSettings());
    }

    public T Deserialize<T>(string value) where T : class
    {
        if (string.IsNullOrEmpty(value))
        {
            throw new ArgumentException(nameof(value));
        }

        return (JsonConvert.DeserializeObject<T>(value))!;
    }
    
    public virtual JsonSerializerSettings OnGetJsonSettings()
    {
        return new JsonSerializerSettings()
        {
            NullValueHandling = NullValueHandling.Ignore,
            Converters = new List<JsonConverter> { new IsoDateTimeConverter() }
        };
    }

}