using Manabie.Api.Controllers;
using Manabie.Api.Insfrastructures;
using Manabie.Api.Middleware;
using Manabie.Api.Services;
using Manabie.Api.Utilities;

using Microsoft.AspNetCore.Builder;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Caching.Memory;
using Microsoft.Extensions.DependencyInjection;

var builder = WebApplication.CreateBuilder(args);

// Add configurations
ConfigurationManager configuration = builder.Configuration;

// Add services to the container.

builder.Services.AddControllers();
builder.Services.AddAuthentication();
builder.Services.AddAuthorization();

// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddDbContext<ManaDbContext>(options =>
    options.UseSqlite("Data Source=sqlite.data")
);
builder.Services.Configure<AppSettings>(configuration.GetSection("AppSettings"));

// configure DI for application services
builder.Services.AddScoped<IUserService, UserService>();
builder.Services.AddScoped<ITaskService, TaskService>();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseDeveloperExceptionPage();
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseRouting();

app.UseHttpsRedirection();

app.UseAuthentication();
app.UseAuthorization();

app.UseMiddleware<JWTMiddleware>();

app.UseEndpoints(endpoints =>
{
    endpoints.MapControllers();
});

app.MapGet("/", () => "hello world!");

app.Run();

public partial class Program
{
}