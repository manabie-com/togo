using Manabie.BasicIdentityServer.Application;
using Manabie.BasicIdentityServer.Infrastructure;
using Manabie.BasicIdentityServer.Infrastructure.Persistence;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.

// Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
builder.Services.AddControllers(c => c.EnableEndpointRouting = false);
builder.Services.AddEndpointsApiExplorer();
builder.Services.AddSwaggerGen();

builder.Services.AddApplication();
builder.Services.AddInfrastructure();
builder.Services.AddMvcCore();
var app = builder.Build();



using (var scope = app.Services.CreateScope())
{
    var initialiser = scope.ServiceProvider.GetRequiredService<DataSeed>();
    await initialiser.InitialiseAsync();
    await initialiser.SeedAsync();
}

app.UseHttpsRedirection();

app.UseCors("MyPolicy");

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.UseSwagger();
    app.UseSwaggerUI();
}

app.UseRouting();

app.UseAuthentication();
app.UseAuthorization();

app.UseEndpoints(options =>
{
    options.MapControllers();
    options.MapDefaultControllerRoute();
});

app.Run();

// Make the implicit Program class public so test projects can access it
public partial class Program { }
