using System;
using System.Collections.Generic;
using System.Linq;
using WebApi.Models;

namespace WebApi.Test
{
    /// <summary>
    /// This class for help data setup
    /// </summary>
    public class TestServiceUtilities
    {
        #region User model
        public static IQueryable<User> SetupDataForUserEntity(string userIds, string password)
        {
            var result = new List<User>();
            var userIdList = userIds.Split("|");
            foreach (var item in userIdList)
            {
                result.Add(new User
                {
                    Id = System.Guid.Parse(item),
                    MaxTodo = 5,
                    Password = password
                });
            }

            return result.AsQueryable();
        }
        #endregion

        #region Task model
        public static IQueryable<Task> SetupDataForTaskEntity(Guid userId, int totalRecords)
        {
            var result = new List<Task>();
            for (int i = 0; i < totalRecords; i++)
            {
                result.Add(new Task
                {
                    CreatedDate = DateTime.Now,
                    Content = $"Task {i + 1}",
                    Id = Guid.NewGuid(),
                    UserId = userId
                });
            }

            return result.AsQueryable();
        }
        #endregion
    }
}
