using Todo.Domain;
using Todo.Domain.Interfaces;
using Todo.Storage.Contract.Interfaces;
using Todo.Storage.Sql;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
var connectionString = builder.Configuration.GetConnectionString("Database");

// Dependency Injection
builder.Services.AddTransient<ITodoManagementService, TodoManagementService>();
builder.Services.AddTransient<IUserManagementService, UserManagementService>();
builder.Services.AddTransient<ISerializer, JsonSerializer>();
builder.Services.AddTransient<ITodoRepository>(r => 
    new TodoSqlRepository(
        connectionString, 
        r.GetRequiredService<ISerializer>()));

builder.Services.AddTransient<IUserRepository>(r =>
    new UserSqlRepository(
        connectionString,
        r.GetRequiredService<ISerializer>()));

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseHttpsRedirection();

app.UseAuthorization();

app.MapControllers();

app.Run();

// For integration tests
public partial class Program { }