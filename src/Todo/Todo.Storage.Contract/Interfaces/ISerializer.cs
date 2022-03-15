namespace Todo.Storage.Contract.Interfaces;

/* 
 * This interface should be used when serializing data.
*/
public interface ISerializer
{
    /*
     * This method Serializes an object value to a string.
     */
    string Serialize(object value);

    /*
     * This method deserialize a string value to the specified type T.
     */
    T Deserialize<T>(string value) where T : class;
}