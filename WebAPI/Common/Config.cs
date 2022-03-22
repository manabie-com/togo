using Microsoft.Extensions.Configuration;
using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Common
{
    public class Config
    {
        private static IConfiguration Configuration { get; set; }

        public static void ConfigStartup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        /// <summary>
        /// Gets the application setting.
        /// </summary>
        /// <param name="key">The key.</param>
        /// <param name="throwException">if set to <c>true</c> [throw exception].</param>
        /// <returns>System.String.</returns>
        public static string GetAppSetting(string key, bool throwException = true) =>
            GetAppSetting<string>(key, throwException, null);
        /// <summary>
        /// Gets the multi value application setting.
        /// </summary>
        /// <param name="key">The key.</param>
        /// <param name="throwException">if set to <c>true</c> [throw exception].</param>
        /// <param name="delimiter">The delimiter.</param>
        /// <returns>System.String[].</returns>
        public static string[] GetMultiValueAppSetting(string key, bool throwException = true, string delimiter = ";")
        {
            string[] strArray;
            string appSetting = GetAppSetting(key, throwException);
            if (appSetting == null)
            {
                strArray = null;
            }
            else
            {
                string[] separator = new string[] { delimiter };
                strArray = appSetting.Split(separator, StringSplitOptions.RemoveEmptyEntries);
            }
            return strArray;
        }
        /// <summary>
        /// Gets the application setting.
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="key">The key.</param>
        /// <param name="throwException">if set to <c>true</c> [throw exception].</param>
        /// <param name="defaultValue">The default value.</param>
        /// <returns>T.</returns>
        /// <exception cref="Exception">The app setting {key}</exception>
        public static T GetAppSetting<T>(string key, bool throwException, T defaultValue = default(T))
        {
            T local;
            if (Configuration == null)
            {
                throw new Exception($"Please add Config.ConfigStartUp(Configuration); in Startup method in Startup.cs");
            }
            string str = Configuration.GetSection($"AppSettings:{key}").Value;
            if (str == null)
            {
                if (throwException)
                {
                    throw new Exception($"The app setting {key} was not found, this is required. Please app {key} to AppSettings block, Ex: \"AppSettings\":{{\"{key}\":\"value\"}}'");
                }
                local = defaultValue;
            }
            else
            {
                string str2 = str;
                Type underlyingType = Nullable.GetUnderlyingType(typeof(T));
                if (underlyingType != null)
                {
                    local = (T)Convert.ChangeType(str2, underlyingType);
                }
                else
                {
                    local = (T)TypeDescriptor.GetConverter(typeof(T)).ConvertFromInvariantString(str2);
                }
            }
            return local;
        }
        /// <summary>
        /// Gets the root application setting.
        /// </summary>
        /// <typeparam name="T"></typeparam>
        /// <param name="key">The key to get param Ex:  "Mail":{"Listmail":{"MailDevCspt":"dev@cspt.com.vn"}} ->  "Mail:Listmail:MailDevCspt" return dev@cspt.com.vn.</param>
        /// <param name="throwException">if set to <c>true</c> [throw exception].</param>
        /// <param name="defaultValue">The default value.</param>
        /// <returns></returns>
        /// <exception cref="Exception">
        /// </exception>
        public static T GetRootAppSetting<T>(string key, bool throwException, T defaultValue = default(T))
        {
            T local;
            if (Configuration == null)
            {
                throw new Exception($"Please add Config.ConfigStartup(Configuration); to Startup method in Startup.cs");
            }
            string str = Configuration.GetSection(key).Value;
            if (str == null)
            {
                if (throwException)
                {
                    throw new Exception($"The {key} was not found, this is required.");
                }
                local = defaultValue;
            }
            else
            {
                string str2 = str;
                Type underlyingType = Nullable.GetUnderlyingType(typeof(T));
                if (underlyingType != null)
                {
                    local = (T)Convert.ChangeType(str2, underlyingType);
                }
                else
                {
                    local = (T)TypeDescriptor.GetConverter(typeof(T)).ConvertFromInvariantString(str2);
                }
            }
            return local;
        }
    }
}
