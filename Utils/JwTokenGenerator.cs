using System.Text;
using System.Security.Claims;
using ManabieTodo.Models;
using System.IdentityModel.Tokens.Jwt;
using Microsoft.IdentityModel.Tokens;
using ManabieTodo.Constants;

namespace ManabieTodo.Utils
{
    public class JwTokenGenerator
    {
        private SymmetricSecurityKey _key { get; }

        public JwTokenGenerator(string secretKey)
        {
            byte[] bytes = Encoding.ASCII.GetBytes(secretKey);
            _key = new SymmetricSecurityKey(bytes);
        }

        public string Create(UserModel model)
        {
            JwtSecurityTokenHandler tokenHandler = new JwtSecurityTokenHandler();
            SecurityTokenDescriptor tokenDescriptor = new SecurityTokenDescriptor
            {
                Subject = new ClaimsIdentity(
                    new[] {
                    new Claim(ClaimTag.Id, model.Id.ToString()),
                    new Claim(ClaimTag.Name, model.Name ?? "Unknown"),
                    new Claim(ClaimTag.AllowedTaskDay, model.AllowedTaskDay.ToString()),
                    }
                ),
                Expires = DateTime.UtcNow.AddHours(8),
                SigningCredentials = new SigningCredentials(
                    _key,
                    SecurityAlgorithms.HmacSha256Signature
                )
            };

            SecurityToken token = tokenHandler.CreateToken(tokenDescriptor);

            return tokenHandler.WriteToken(token);
        }
    }
}