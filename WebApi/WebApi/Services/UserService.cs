﻿using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.IdentityModel.Tokens;
using System;
using System.IdentityModel.Tokens.Jwt;
using System.Security.Claims;
using System.Text;
using System.Threading.Tasks;
using WebApi.Models;
using WebApi.ViewModels;

namespace WebApi.Services
{
    public class UserService : IUserService
    {
        private readonly DemoDbContext _demoDbContext;
        private readonly IConfiguration _configuration;

        public UserService(DemoDbContext demoDbContext, IConfiguration configuration)
        {
            _demoDbContext = demoDbContext;
            _configuration = configuration;
        }

        /// <summary>
        /// Login service
        /// </summary>
        /// <param name="userId"></param>
        /// <param name="password"></param>
        /// <returns></returns>
        public async Task<LoginViewModel> Login(Guid userId, string password)
        {
            var user = await _demoDbContext.Users.SingleOrDefaultAsync(x => x.Id == userId && password == password.Trim());
            if (user == null) return null;

            var token = GenerateJwtToken(user);

            return new LoginViewModel
            {
                Id = user.Id,
                Token = token
            };
        }

        #region Private methods
        /// <summary>
        /// Generate JWT token by user info
        /// </summary>
        /// <param name="user"></param>
        /// <returns></returns>
        private string GenerateJwtToken(User user)
        {
            // generate token that is valid for 7 days
            var tokenHandler = new JwtSecurityTokenHandler();
            var key = Encoding.ASCII.GetBytes(_configuration.GetValue<string>("AppSettings:Secret"));
            var tokenDescriptor = new SecurityTokenDescriptor
            {
                Subject = new ClaimsIdentity(new[] { new Claim("id", user.Id.ToString()) }),
                Expires = DateTime.UtcNow.AddDays(7),
                SigningCredentials = new SigningCredentials(new SymmetricSecurityKey(key), SecurityAlgorithms.HmacSha256Signature)
            };
            var token = tokenHandler.CreateToken(tokenDescriptor);
            return tokenHandler.WriteToken(token);
        }
        #endregion
    }
}
