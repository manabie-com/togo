using Microsoft.OpenApi.Models;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Conventions;
using MongoDB.Driver;
using ToDo.Api.Repositories;
using ToDo.Api.Validators;

namespace ToDo.Api
{
    public class Startup
    {
        private readonly IConfiguration _configuration;

        public Startup(IConfiguration configuration)
        {
            _configuration = configuration;
        }

        // TODO Add MongoDb settings in appsettings
        public void ConfigureServices(IServiceCollection services)
        {


            services.AddControllers();


            #region Swagger Configurations

            var assemblyPath = typeof(Startup).Assembly.Location;

            services.AddSwaggerGen(c =>
            {
                c.SwaggerDoc("v1",
                    new OpenApiInfo
                    {
                        Title = "Todo Api MongoDB v1",
                        Version = "v1",
                        Description = "Todo Api MongoDB v1"
                    });
            });

            #endregion


            #region Mongo Configurations

            var mongoClientSettings = MongoClientSettings.FromConnectionString(_configuration.GetConnectionString("Mongo"));

            services.AddSingleton(mongoClientSettings);

            services.AddSingleton<IConventionPack>(new ConventionPack
            {
                new EnumRepresentationConvention(BsonType.String),
                new IgnoreExtraElementsConvention(true),
                new IgnoreIfNullConvention(true)
            });

            services.AddSingleton<IMongoClient>(new MongoClient(mongoClientSettings));
            services.AddScoped(typeof(IMongoBaseRepository<>), typeof(MongoBaseRepository<>));
            services.AddScoped<CreateUserValidator>();
            services.AddScoped<CreateToDoValidator>();

            #endregion
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
                app.UseSwagger();
                app.UseSwaggerUI(c => c.SwaggerEndpoint("/swagger/v1/swagger.json", "TODO.Api v1"));
            }

            app.UseHttpsRedirection();

            app.UseRouting();

            app.UseAuthorization();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
            });
        }
    }
}
