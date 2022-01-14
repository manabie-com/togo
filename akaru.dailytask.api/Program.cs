using akaru.dailytask.api.Database;

var builder = WebApplication.CreateBuilder(args);

// Inject Database in the Services for dependency injection
builder.Services.AddDbContext<DatabaseContext>();

var context = builder.Services.BuildServiceProvider().GetService<DatabaseContext>();
context.Database.EnsureCreated();

// Add services to the container.
builder.Services.AddControllers();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (!app.Environment.IsDevelopment())
{
    app.UseExceptionHandler("/Home/Error");
}
app.UseStaticFiles();

app.UseRouting();

app.UseAuthorization();

app.MapControllerRoute(
    name: "default",
    pattern: "{controller=User}/{action=Index}");

app.Run();

