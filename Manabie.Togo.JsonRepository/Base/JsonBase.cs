using Manabie.Togo.Data.Base;
using Newtonsoft.Json;
using System;
using System.Collections.Concurrent;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Manabie.Togo.JsonRepository.Base
{
	public class JsonBase<T> : IJsonBase<T> where T : BaseEntity
	{
		private ConcurrentDictionary<Guid, T> _dicData;
		private string _fullpath;

		public JsonBase()
		{
			var fileName = typeof(T).Name;
			_dicData = new ConcurrentDictionary<Guid, T>();
			_fullpath = Directory.GetCurrentDirectory() + @"\JsonFiles\" + fileName + ".json";
			_dicData = Init();
		}

		/// <summary>
		/// Init data from json file
		/// </summary>
		/// <returns></returns>
		private ConcurrentDictionary<Guid, T> Init()
		{
			try
			{
				lock (_dicData)
				{
					// Data empty
					if (!File.Exists(_fullpath))
					{
						if (!Directory.Exists(Path.GetDirectoryName(_fullpath)))
							Directory.CreateDirectory(Path.GetDirectoryName(_fullpath));
						System.IO.FileStream f = System.IO.File.Create(_fullpath);
						f.Close();
						return new ConcurrentDictionary<Guid, T>();
					}
					string jsonText;

					using (StreamReader sr = File.OpenText(_fullpath))
					{
						jsonText = sr.ReadToEnd();
					}

					var datas = JsonConvert.DeserializeObject<ConcurrentDictionary<Guid, T>>(jsonText);
					if (datas == null)
						return new ConcurrentDictionary<Guid, T>();
					return datas;
				}
			}
			catch (Exception ex)
			{
				return new ConcurrentDictionary<Guid, T>();
			}
		}

		/// <summary>
		/// Get object by id
		/// </summary>
		/// <param name="id"></param>
		/// <returns></returns>
		public T GetbyId(Guid id)
		{
			lock (_dicData)
			{
				if (_dicData.ContainsKey(id))
					return _dicData[id];
				else
					return default;
			}
		}

		/// <summary>
		/// Add a object
		/// </summary>
		/// <param name="data"></param>
		/// <returns></returns>
		public bool Add(T data)
		{
			lock (_dicData)
			{
				if (!_dicData.ContainsKey(data.ID))
					return _dicData.TryAdd(data.ID, data);
				else return false;
			}
		}

		/// <summary>
		/// Add many object
		/// </summary>
		/// <param name="datas"></param>
		/// <returns></returns>
		public bool Add(IEnumerable<T> datas)
		{
			lock (_dicData)
			{
				foreach (var data in datas)
				{
					if (!_dicData.ContainsKey(data.ID))
						return _dicData.TryAdd(data.ID, data);
				}
			}

			return true;
		}

		/// <summary>
		/// Update a obkect that exist
		/// </summary>
		/// <param name="data"></param>
		/// <returns></returns>
		public bool Update(T data)
		{
			try
			{
				lock (_dicData)
				{
					if (_dicData.ContainsKey(data.ID))
					{
						_dicData[data.ID] = data;
					}
					else
					{
						return false;
					}
				}
				return true;
			}
			catch (Exception ex)
			{
				return false;
			}
		}

		/// <summary>
		/// Save all change
		/// </summary>
		/// <returns></returns>
		public async Task<bool> SaveChange()
		{
			try
			{
				var dataAsync = Task.Run(delegate ()
				{
					//open file stream
					using (StreamWriter sw = new StreamWriter(_fullpath, true))
					{
						JsonSerializer serializer = new JsonSerializer();
						//serialize object directly into file stream
						serializer.Serialize(sw, _dicData);
					}
					return true;
				});
				return await dataAsync;
			}
			catch (Exception ex)
			{
				return false;
			}
		}

		/// <summary>
		/// Get all datas
		/// </summary>
		/// <returns></returns>
		public IEnumerable<T> GetAll()
		{
			return ConvertDicToIEnumer(_dicData);
		}

		private IEnumerable<T> ConvertDicToIEnumer(ConcurrentDictionary<Guid, T> datas)
		{
			if (datas == null)
				return default;
			return datas.Select(x => x.Value);
		}

		/// <summary>
		/// Delete a obkect that exist
		/// </summary>
		/// <param name="data"></param>
		/// <returns></returns>
		public bool Delete(Guid id)
		{
			lock (_dicData)
			{
				if (_dicData.ContainsKey(id))
				{
					_dicData.TryRemove(id, out _);
					return true;
				}
				else
				{
					return false;
				}
			}
		}
	}
}
