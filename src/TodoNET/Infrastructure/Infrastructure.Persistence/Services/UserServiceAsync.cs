using Application.DTOs.User;
using Application.Exceptions;
using Application.Interfaces;
using Application.Interfaces.Repositories;
using Application.Wrappers;
using AutoMapper;
using Domain.Entities;
using Domain.Settings;
using Infrastructure.Persistence.Contexts;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Options;
using Microsoft.IdentityModel.Tokens;
using System;
using System.Collections.Generic;
using System.IdentityModel.Tokens.Jwt;
using System.IO;
using System.Security.Claims;
using System.Security.Cryptography;
using System.Text;
using System.Threading.Tasks;

namespace Infrastructure.Persistence.Services
{
    public class UserServiceAsync : IUserServiceAsync
    {
        private readonly DbSet<User> _user;
        private readonly JWTSettings _jwtSettings;
        private readonly AppSettings _appSettings;
        private readonly IGenericRepositoryAsync<User> _genericRepositoryAsync;
        private readonly IMapper _mapper;

        public UserServiceAsync(ApplicationDbContext dbContext,
            IOptions<JWTSettings> jwtSettings,
            IOptions<AppSettings> appSettings,
            IGenericRepositoryAsync<User> genericRepositoryAsync
            , IMapper mapper)
        {
            _user = dbContext.Set<User>();
            _jwtSettings = jwtSettings.Value;
            _appSettings = appSettings.Value;
            _genericRepositoryAsync = genericRepositoryAsync;
            _mapper = mapper;
        }

        public async Task<Response<string>> CreateUserAsync(CreateUserRequest request)
        {
            if (request == null)
            {
                throw new ArgumentNullException(nameof(request));
            }
            var user = await _user.FirstOrDefaultAsync(_ => _.Email == request.Email);
            if (user != null)
            {
                throw new ApiException($"{request.Email} exists in system");
            }
            var hashPassword = Encrypt(request.Password);
            var newUser = new User
            {
                Id = Guid.NewGuid().ToString(),
                Email = request.Email,
                MaxTodo = request.MaxTodo,
                CreatedBy = string.Empty,
                Password = hashPassword,
            };
            await _genericRepositoryAsync.AddAsync(newUser);
            return new Response<string>(newUser.Id,string.Empty);
        }

        public async Task<Response<AuthenticationResponse>> AuthenticateAsync(AuthenticationRequest request)
        {
            if (request == null) throw new ArgumentNullException(nameof(request));

            var user = await _user.FirstOrDefaultAsync(_ => _.Email == request.Email);
            if (user == null)
            {
                throw new ApiException($"No Accounts Registered with {request.Email}.");
            }
            var result = Encrypt(request.Password) == user.Password;
            if (!result)
            {
                throw new ApiException($"Invalid Credentials for '{request.Email}'.");
            }

            JwtSecurityToken jwtSecurityToken = GenerateJWToken(user);
            AuthenticationResponse response = new AuthenticationResponse();
            response.Id = user.Id;
            response.JWToken = new JwtSecurityTokenHandler().WriteToken(jwtSecurityToken);
            response.Email = user.Email;
            return new Response<AuthenticationResponse>(response, $"Authenticated {user.Email}");
        }

        public async Task<Response<IReadOnlyList<AuthenticationResponse>>> GetUsers()
        {
            var users = await _genericRepositoryAsync.GetAllAsync();
            var result = _mapper.Map<IReadOnlyList<AuthenticationResponse>>(users);
            return new Response<IReadOnlyList<AuthenticationResponse>>(result);
        }
        public string Encrypt(string clearText)
        {
            string EncryptionKey = _appSettings.Secret;
            byte[] clearBytes = Encoding.Unicode.GetBytes(clearText);
            using (Aes encryptor = Aes.Create())
            {
                Rfc2898DeriveBytes pdb = new Rfc2898DeriveBytes(EncryptionKey, new byte[] { 0x49, 0x76, 0x61, 0x6e, 0x20, 0x4d, 0x65, 0x64, 0x76, 0x65, 0x64, 0x65, 0x76 });
                encryptor.Key = pdb.GetBytes(32);
                encryptor.IV = pdb.GetBytes(16);
                using (MemoryStream ms = new MemoryStream())
                {
                    using (CryptoStream cs = new CryptoStream(ms, encryptor.CreateEncryptor(), CryptoStreamMode.Write))
                    {
                        cs.Write(clearBytes, 0, clearBytes.Length);
                        cs.Close();
                    }
                    clearText = Convert.ToBase64String(ms.ToArray());
                }
            }
            return clearText;
        }
        #region Private
        private JwtSecurityToken GenerateJWToken(User user)
        {
            var claims = new[]
            {
                new Claim(JwtRegisteredClaimNames.Jti, Guid.NewGuid().ToString()),
                new Claim(JwtRegisteredClaimNames.Email, user.Email),
                new Claim("uid", user.Id),
            };

            var symmetricSecurityKey = new SymmetricSecurityKey(Encoding.UTF8.GetBytes(_jwtSettings.Key));
            var signingCredentials = new SigningCredentials(symmetricSecurityKey, SecurityAlgorithms.HmacSha256);

            var jwtSecurityToken = new JwtSecurityToken(
                issuer: _jwtSettings.Issuer,
                audience: _jwtSettings.Audience,
                claims: claims,
                expires: DateTime.UtcNow.AddMinutes(_jwtSettings.DurationInMinutes),
                signingCredentials: signingCredentials);
            return jwtSecurityToken;
        }
        
        #endregion
    }
}
