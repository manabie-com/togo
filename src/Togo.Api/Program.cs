using System.Text;
using Microsoft.AspNetCore.Authentication.JwtBearer;
using Microsoft.AspNetCore.Identity;
using Microsoft.EntityFrameworkCore;
using Microsoft.IdentityModel.Tokens;
using Togo.Core.Interfaces;
using Togo.Infrastructure;
using Togo.Infrastructure.Identities;
using Togo.Infrastructure.Persistence;
using Togo.Infrastructure.Services;

var builder = WebApplication.CreateBuilder(args);

var togoAppSettings = builder.Configuration.Get<TogoAppSettings>();
builder.Services.AddSingleton(togoAppSettings);

builder.Services.AddControllers();
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddDbContext<AppDbContext>(options =>
    options.UseNpgsql(builder.Configuration.GetConnectionString("Default")));

builder.Services.AddScoped<IAppDbContext, AppDbContext>();

builder.Services.AddIdentityCore<AppUser>()
    .AddRoles<IdentityRole>()
    .AddRoleManager<RoleManager<IdentityRole>>()
    .AddEntityFrameworkStores<AppDbContext>();

builder.Services.AddScoped<IUserService, UserService>();

builder.Services
    .AddAuthentication(authenticationOptions =>
    {
        authenticationOptions.DefaultAuthenticateScheme = JwtBearerDefaults.AuthenticationScheme;
        authenticationOptions.DefaultChallengeScheme = JwtBearerDefaults.AuthenticationScheme;
    })
    .AddJwtBearer(jwtBearerOptions =>
{
    jwtBearerOptions.RequireHttpsMetadata = false;
    jwtBearerOptions.SaveToken = true;
    jwtBearerOptions.TokenValidationParameters = new TokenValidationParameters
    {
        ValidateIssuer = true,
        ValidIssuer = togoAppSettings.JwtBearer.Issuer,
        ValidateIssuerSigningKey = true,
        IssuerSigningKey = new SymmetricSecurityKey(
            Encoding.ASCII.GetBytes(
                togoAppSettings.JwtBearer.SecurityKey)),
        
        ValidateAudience = true,
        ValidAudience = togoAppSettings.JwtBearer.Audience
    };
});

builder.Services.AddHttpContextAccessor();
builder.Services.AddSingleton<ICurrentUserService, CurrentUserService>();

var app = builder.Build();

using var scope = app.Services.CreateScope();
var dbContext = scope.ServiceProvider.GetRequiredService<AppDbContext>();
await dbContext.Database.MigrateAsync();

var userService = scope.ServiceProvider.GetRequiredService<IUserService>();
await userService.SeedAdminUserAsync();

if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthentication();
app.UseAuthorization();

app.MapControllers();

app.Run();

public partial class Program { }
