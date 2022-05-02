using System.Collections.Generic;

namespace Manabie.Togo.Core.Base
{
	public class ErrorCodeMessage
	{
		public static readonly KeyValuePair<int, string> Success = new KeyValuePair<int, string>(0, "The operation completed successfully.");
		public static readonly KeyValuePair<int, string> IncorrectFunction = new KeyValuePair<int, string>(1, "Incorrect function.");
		public static readonly KeyValuePair<int, string> UserExisted = new KeyValuePair<int, string>(2, "This email already exist.");
		public static readonly KeyValuePair<int, string> UserNotExisted = new KeyValuePair<int, string>(3, "User does not existed");
		public static readonly KeyValuePair<int, string> RoleNotExisted = new KeyValuePair<int, string>(4, "This role not existed.");
		public static readonly KeyValuePair<int, string> RoleEmpty = new KeyValuePair<int, string>(4, "RoleCategorys can not empty.");
	}
}
