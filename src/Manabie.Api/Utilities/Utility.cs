using System;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using System.Security.Cryptography;
using Manabie.Api.Entities;

using Microsoft.IdentityModel.Tokens;
using Microsoft.AspNetCore.Cryptography.KeyDerivation;
using Microsoft.AspNetCore.Mvc.Diagnostics;
using Microsoft.EntityFrameworkCore.Storage.ValueConversion.Internal;

namespace Manabie.Api.Utilities;

public static class Utility
{
    public static string GenerateJwtToken(User user, string secret)
    {
        // generate token that is valid for 7 days
        var tokenHandler = new JwtSecurityTokenHandler();
        var key = Encoding.UTF8.GetBytes(secret);
        var tokenDescriptor = new SecurityTokenDescriptor
        {
            Subject = new ClaimsIdentity(new[] { new Claim("Id", user.Id.ToString()) }),
            Expires = DateTime.UtcNow.AddDays(7),
            SigningCredentials = new SigningCredentials(new SymmetricSecurityKey(key), SecurityAlgorithms.HmacSha256Signature)
        };
        var token = tokenHandler.CreateToken(tokenDescriptor);
        return tokenHandler.WriteToken(token);
    }

    public static string HashSh256(string input, string secret)
    {
        byte[] salt = Encoding.UTF8.GetBytes(secret);

        string hashed = Convert.ToBase64String(KeyDerivation.Pbkdf2(
            password: input,
            salt: salt,
            prf: KeyDerivationPrf.HMACSHA256,
            iterationCount: 100000,
            numBytesRequested: 256 / 8));
        return hashed;
    }
}

