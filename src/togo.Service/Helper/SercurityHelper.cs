using System;
using System.Linq;
using System.Security.Cryptography;
using System.Text;

namespace togo.Service.Interface
{
    public class SercurityHelper
    {
        public static string GenerateSalt()
        {
            var bytes = new byte[128 / 8];
            var rng = new RNGCryptoServiceProvider();
            rng.GetBytes(bytes);
            return Convert.ToBase64String(bytes);
        }

        public static string ComputeHash(string password, string salt)
        {
            byte[] unhashedBytes = Encoding.Unicode.GetBytes(String.Concat(salt, password));

            SHA256Managed sha256 = new SHA256Managed();
            byte[] hashedBytes = sha256.ComputeHash(unhashedBytes);

            return Convert.ToBase64String(hashedBytes);
        }

        public static bool ComparePassword(string password, string salt, string hash)
        {
            var attemptedHash = ComputeHash(password, salt);
            return attemptedHash == hash;
        }
    }
}
