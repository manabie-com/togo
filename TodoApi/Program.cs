using System.Reflection;
using System.Text.Json.Serialization;
using FluentValidation;
using MediatR;
using Microsoft.EntityFrameworkCore;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

builder.Services.AddControllers();
// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

// config db
var server = builder.Configuration["DbServer"] ?? "localhost";
var port = builder.Configuration["DbPort"] ?? "1433"; // Default SQL Server port
var user = builder.Configuration["DbUser"] ?? "SA"; // Warning do not use the SA account
var password = builder.Configuration["Password"] ?? "Str0ngPa$$w0rd";
var database = builder.Configuration["Database"] ?? "ToDoDb";
var connectionString = $"Server={server}, {port};Initial Catalog={database};User ID={user};Password={password}";
builder.Services.AddDbContext<ToDoDbContext>(options => options.UseSqlServer(connectionString));
//builder.Services.AddDbContext<ToDoDbContext>(options => options.UseSqlServer(builder.Configuration.GetConnectionString("DefaultConnection")));

builder.Services.AddScoped<IUserRepository, UserRepository>();
builder.Services.AddMediatR(typeof(Program));
builder.Services.AddTransient(typeof(IPipelineBehavior<,>), typeof(ValidationBehavior<,>));
builder.Services.AddValidatorsFromAssembly(Assembly.GetExecutingAssembly());

var app = builder.Build();

using (var scope = app.Services.CreateScope())
{
    var dataContext = scope.ServiceProvider.GetRequiredService<ToDoDbContext>();
    dataContext.Database.Migrate();
}

// Configure the HTTP request pipeline.
app.UseSwagger();
app.UseSwaggerUI();

app.MapFallback(() => Results.Redirect("/swagger"));

// app.UseAuthorization();

app.MapControllers();

app.MapGet("/api/user/{id}",
    async (int id, ISender sender) => await sender.Send(new UserQuery { Id = id }));

app.MapPost("/api/user",
    async (CreateUserCommand user, ISender sender) => await sender.Send(user));

app.MapPost("/api/todo",
    async (CreateToDoCommand toDo, ISender sender) => await sender.Send(toDo));    

app.Run();
