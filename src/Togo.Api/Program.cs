using Microsoft.EntityFrameworkCore;
using Togo.Core;
using Togo.Infrastructure;
using Togo.Infrastructure.Identities;
using Togo.Infrastructure.Persistence;

var builder = WebApplication.CreateBuilder(args);

// Dependency Injection
var togoAppSettings = builder.Configuration.Get<TogoAppSettings>();
builder.Services.AddSingleton(togoAppSettings);
builder.Services.AddCoreServices();
builder.Services.AddInfrastructureServices(togoAppSettings);

var app = builder.Build();

// Database migration and data seeding
using var scope = app.Services.CreateScope();
var dbContext = scope.ServiceProvider.GetRequiredService<AppDbContext>();
await dbContext.Database.MigrateAsync();

var userService = scope.ServiceProvider.GetRequiredService<IUserService>();
await userService.SeedAdminUserAsync();

// Middlewares
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
