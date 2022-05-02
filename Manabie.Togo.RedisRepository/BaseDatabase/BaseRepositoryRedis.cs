using Manabie.Togo.Data.Base;
using Newtonsoft.Json;
using StackExchange.Redis;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Togo.RedisRepository.BaseDatabase
{
    public class BaseRepositoryRedis<T> : IBaseRepositoryRedis<T> where T : BaseEntity
    {
        protected readonly IConnectionMultiplexer _redisConnection;
        protected readonly IDatabase _database;

        /// <summary>
        /// The Namespace is the first part of any key created by this Repository, e.g. "tenant" or "company"
        /// </summary>
        protected readonly string _namespace;

        public BaseRepositoryRedis(IConnectionMultiplexer redisConnection, IDatabase database, string nameSpace)
        {
            _redisConnection = redisConnection;
            _database = database;
            _namespace = nameSpace;
        }

        public bool Exists(Guid id)
        {
            return Get(MakeKey(id)) != null;
        }

        public bool Exists(string key)
        {
            return _database.KeyExists(key);
        }

        public IEnumerable<T> GetAll()
        {
            var key = MakeKey("all");
            var serializedObject = _database.StringGet(key);
            if (serializedObject.IsNullOrEmpty)
            {
                return new List<T>();
            }
            return JsonConvert.DeserializeObject<List<T>>(serializedObject.ToString()) ?? new List<T>();
        }

        public async Task<IEnumerable<T>> GetAllAsync()
        {
            var key = MakeKey("all");
            var serializedObject = await _database.StringGetAsync(key);
            if (serializedObject.IsNullOrEmpty)
            {
                return new List<T>();
            }
            return JsonConvert.DeserializeObject<List<T>>(serializedObject.ToString()) ?? new List<T>();
        }

        public T GetByID(Guid id)
        {
            return Get(MakeKey(id));
        }

        public async Task<IEnumerable<T>> GetAllBykeyAsync(string key)
        {
            try
            {
				var dbExc = _database.Execute($"KEYS", new string[1] { $"{key}:*" });
				var keyR = (RedisKey[])dbExc;
				var value = await _database.StringGetAsync(keyR);
                if (value.Count() > 0)
                {                 
                    var result = value.Select(serializedObject=> JsonConvert.DeserializeObject<IEnumerable<T>>(serializedObject.ToString()));
                    return result.SelectMany(x => x);
                }
                return new List<T>();
            }
            catch (Exception ex)
			{
                return new List<T>();
			}            
        }

        public async Task<T> GetByIDAsync(Guid id)
        {
            return await GetAsync(MakeKey(id));
        }

        public IEnumerable<T> GetMultiple(IEnumerable<Guid> ids)
        {
            if (ids == null || !ids.Any())
            {
                return new List<T>();
            }
            return Get(ids.Select(x => MakeKey(x)).ToList());
        }

        public async Task<IEnumerable<T>> GetMultipleAsync(IEnumerable<Guid> ids)
        {
            if (ids == null || !ids.Any())
            {
                return new List<T>();
            }
            return await GetAsync(ids.Select(x => MakeKey(x)).ToList());
        }

        public virtual void Save(T item)
        {
            _database.StringSet(MakeKey(item.ID), JsonConvert.SerializeObject(item));

            MergeIntoAllCollection(item);
        }

        public virtual async Task SaveByKeyAsync(string key,T item)
        {
            var setDataTask = new List<Task>();

            var taskItem = _database.StringSetAsync(key, JsonConvert.SerializeObject(item));
            setDataTask.Add(taskItem);

            var taskAll = MergeIntoAllCollectionAsync(key,item);
            setDataTask.Add(taskAll);

            await Task.WhenAll(setDataTask);
        }

        public virtual async Task SaveAsync(T item)
        {
            var setDataTask = new List<Task>();

            var taskItem = _database.StringSetAsync(MakeKey(item.ID), JsonConvert.SerializeObject(item));
            setDataTask.Add(taskItem);

            var taskAll = MergeIntoAllCollectionAsync(item);
            setDataTask.Add(taskAll);

            await Task.WhenAll(setDataTask);
        }

        public void Save(IEnumerable<T> items)
        {
            var dataToSet = items.ToDictionary(x => (RedisKey)MakeKey(x.ID), x => (RedisValue)JsonConvert.SerializeObject(x)).ToArray();
            _database.StringSet(dataToSet);

            MergeIntoAllCollection(items);
        }

        public async Task SaveAsync(IEnumerable<T> items)
        {
            var dataToSet = items.ToDictionary(x => (RedisKey)MakeKey(x.ID), x => (RedisValue)JsonConvert.SerializeObject(x)).ToArray();

            var setDataTask = new List<Task>();

            var taskItem = _database.StringSetAsync(dataToSet);
            setDataTask.Add(taskItem);

            var taskAll = MergeIntoAllCollectionAsync(items);
            setDataTask.Add(taskAll);

            await Task.WhenAll(setDataTask);
        }

        public void Remove(T item)
        {
            Remove(item.ID);
        }

        public async Task RemoveAsync(T item)
        {
            await RemoveAsync(item.ID);
        }

        public void Remove(IEnumerable<T> items)
        {
            Remove(items.Select(x => x.ID).ToList());
        }

        public async Task RemoveAsync(IEnumerable<T> items)
        {
            await RemoveAsync(items.Select(x => x.ID).ToList());
        }

        public void Remove(Guid id)
        {
            Remove(MakeKey(id));
            RemoveFromAllCollection(id);
        }

        public async Task RemoveAsync(Guid id)
        {
            var removeDataTask = new List<Task>();

            var taskItem = RemoveAsync(MakeKey(id));
            removeDataTask.Add(taskItem);

            var taskAll = RemoveFromAllCollectionAsync(id);
            removeDataTask.Add(taskAll);

            await Task.WhenAll(removeDataTask);
        }

        public void Remove(IEnumerable<Guid> ids)
        {
            Remove(ids.Select(x => MakeKey(x)).ToList());
            RemoveFromAllCollection(ids);
        }

        public async Task RemoveAsync(IEnumerable<Guid> ids)
        {
            var removeDataTask = new List<Task>();

            var taskItem = RemoveAsync(ids.Select(x => MakeKey(x)).ToList());
            removeDataTask.Add(taskItem);

            var taskAll = RemoveFromAllCollectionAsync(ids);
            removeDataTask.Add(taskAll);

            await Task.WhenAll(removeDataTask);
        }

        public void Remove(string key)
        {
            _database.KeyDelete(key);
        }

        public async Task RemoveAsync(string key)
        {
            await _database.KeyDeleteAsync(key);
        }

        public void Remove(IEnumerable<string> keys)
        {
            _database.KeyDelete(keys.Select(x => (RedisKey)x).ToArray());
        }

        public async Task RemoveAsync(IEnumerable<string> keys)
        {
            await _database.KeyDeleteAsync(keys.Select(x => (RedisKey)x).ToArray());
        }

        public void ClearAll()
        {
            var result = _database.Execute($"KEYS", new string[1] { $"{_namespace}:*" });
            var keys = (RedisKey[])result;

            _database.KeyDelete(keys);
        }

        public async Task ClearAllAsync()
        {
            var result = _database.Execute($"KEYS", new string[1] { $"{_namespace}:*" });
            var keys = (RedisKey[])result;

            await _database.KeyDeleteAsync(keys);
        }

        #region Private methods
        protected string MakeKey(Guid id)
        {
            return AddPrefix(id.ToString());
        }

        private string MakeKey(string keySuffix)
        {
            if (keySuffix.StartsWith($"{_namespace}:"))
            {
                return keySuffix;
            }
            return AddPrefix(keySuffix);
        }

        private string AddPrefix(string keySuffix)
        {
            return $"{_namespace}:{keySuffix}";
        }

        private T Get(string key)
        {
            var serializedObject = _database.StringGet(key);
            if (serializedObject.IsNullOrEmpty)
            {
                return null;
            }
            return JsonConvert.DeserializeObject<T>(serializedObject.ToString());
        }

        private async Task<T> GetAsync(string key)
        {
            var serializedObject = await _database.StringGetAsync(key);
            if (serializedObject.IsNullOrEmpty)
            {
                return null;
            }
            return JsonConvert.DeserializeObject<T>(serializedObject.ToString());
        }

        private IEnumerable<T> Get(IEnumerable<string> keys)
        {
            var redisKeys = keys.Select(x => (RedisKey)x).ToList();
            return Get(redisKeys);
        }

        private async Task<IEnumerable<T>> GetAsync(IEnumerable<string> keys)
        {
            var redisKeys = keys.Select(x => (RedisKey)x).ToList();
            return await GetAsync(redisKeys);
        }

        private IEnumerable<T> Get(IEnumerable<RedisKey> keys)
        {
            var redisKeys = keys.ToArray();
            return Get(redisKeys);
        }

        private async Task<IEnumerable<T>> GetAsync(IEnumerable<RedisKey> keys)
        {
            var redisKeys = keys.ToArray();
            return await GetAsync(redisKeys);
        }

        private IEnumerable<T> Get(RedisKey[] keys)
        {
            var serializedObjects = _database.StringGet(keys);

            var result = new List<T>();

            if (serializedObjects != null)
            {
                foreach (var serializedObject in serializedObjects)
                {
                    result.Add(JsonConvert.DeserializeObject<T>(serializedObject.ToString()));
                }
            }

            return result;
        }

        private async Task<IEnumerable<T>> GetAsync(RedisKey[] keys)
        {
            var serializedObjects = await _database.StringGetAsync(keys);

            var result = new List<T>();

            if (serializedObjects != null)
            {
                foreach (var serializedObject in serializedObjects)
                {
                    result.Add(JsonConvert.DeserializeObject<T>(serializedObject.ToString()));
                }
            }

            return result;
        }

        protected void MergeIntoAllCollection(T item)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Except(items.Where(x => x.ID == item.ID), new BaseDtoCompare<T>()).ToList();

            //Add the modified district to the ALL collection
            items.Add(item);

            _database.StringSet(MakeKey("all"), JsonConvert.SerializeObject(items));
        }

        protected void MergeIntoAllCollection(IEnumerable<T> items)
        {
            List<T> itemsOld = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            itemsOld = itemsOld.Except(itemsOld.Where(x => items.Any(y => y.ID == x.ID)), new BaseDtoCompare<T>()).ToList();

            //Add the modified district to the ALL collection
            itemsOld.AddRange(items);

            _database.StringSet(MakeKey("all"), JsonConvert.SerializeObject(itemsOld));
        }

        private void RemoveFromAllCollection(T item)
        {
            RemoveFromAllCollection(item.ID);
        }

        private void RemoveFromAllCollection(Guid id)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Except(items.Where(x => x.ID == id), new BaseDtoCompare<T>()).ToList();

            _database.StringSet(MakeKey("all"), JsonConvert.SerializeObject(items));
        }

        private async Task MergeIntoAllCollectionAsync(string key,T item)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Except(items.Where(x => x.ID == item.ID), new BaseDtoCompare<T>()).ToList();

            //Add the modified district to the ALL collection
            items.Add(item);

            await _database.StringSetAsync(key, JsonConvert.SerializeObject(items));
        }

        private async Task MergeIntoAllCollectionAsync(T item)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Except(items.Where(x => x.ID == item.ID), new BaseDtoCompare<T>()).ToList();

            //Add the modified district to the ALL collection
            items.Add(item);

            await _database.StringSetAsync(MakeKey("all"), JsonConvert.SerializeObject(items));
        }

        private async Task RemoveFromAllCollectionAsync(T item)
        {
            await RemoveFromAllCollectionAsync(item.ID);
        }

        private async Task RemoveFromAllCollectionAsync(Guid id)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Except(items.Where(x => x.ID == id), new BaseDtoCompare<T>()).ToList();

            await _database.StringSetAsync(MakeKey("all"), JsonConvert.SerializeObject(items));
        }

        private void MergeIntoAllCollectionơ(IEnumerable<T> itemsToSave)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Except(itemsToSave, new BaseDtoCompare<T>()).ToList();

            //Add the modified district to the ALL collection
            items.AddRange(itemsToSave);

            _database.StringSet(MakeKey("all"), JsonConvert.SerializeObject(items));
        }

        private void RemoveFromAllCollection(IEnumerable<T> itemsToRemove)
        {
            RemoveFromAllCollection(itemsToRemove.Select(x => x.ID).ToList());
        }

        private void RemoveFromAllCollection(IEnumerable<Guid> idsToRemove)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Where(x => !idsToRemove.Contains(x.ID)).ToList();

            _database.StringSet(MakeKey("all"), JsonConvert.SerializeObject(items));
        }

        private async Task MergeIntoAllCollectionAsync(IEnumerable<T> itemsToSave)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Except(itemsToSave, new BaseDtoCompare<T>()).ToList();

            //Add the modified district to the ALL collection
            items.AddRange(itemsToSave);

            await _database.StringSetAsync(MakeKey("all"), JsonConvert.SerializeObject(items));
        }

        private async Task RemoveFromAllCollectionAsync(IEnumerable<T> itemsToRemove)
        {
            await RemoveFromAllCollectionAsync(itemsToRemove.Select(x => x.ID).ToList());
        }

        private async Task RemoveFromAllCollectionAsync(IEnumerable<Guid> idsToRemove)
        {
            List<T> items = GetAll().ToList();

            //If the district already exists in the ALL collection, remove that entry
            items = items.Where(x => !idsToRemove.Contains(x.ID)).ToList();

            await _database.StringSetAsync(MakeKey("all"), JsonConvert.SerializeObject(items));
        }

        #endregion
    }
}
