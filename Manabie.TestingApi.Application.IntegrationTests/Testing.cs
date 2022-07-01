using Manabie.BasicIdentityServer.Infrastructure.Identity;
using Manabie.Testing.Infrastructure.Persistence;
using MediatR;
using Microsoft.AspNetCore.Identity;
using Microsoft.AspNetCore.Mvc.Testing;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using NUnit.Framework;
using Respawn;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Manabie.TestingApi.Application.IntegrationTests;

[SetUpFixture]
public partial class Testing
{
    private static WebApplicationFactory<Program> _factory = null!;
    private static IConfiguration _configuration = null!;
    private static IServiceScopeFactory _scopeFactory = null!;
    private static Checkpoint _checkpoint = null!;
    private static string? _currentUserId;
    private static string? _currentRole;

    [OneTimeSetUp]
    public void RunBeforeAnyTests()
    {
        _factory = new CustomWebApplicationFactory();
        _scopeFactory = _factory.Services.GetRequiredService<IServiceScopeFactory>();
        _configuration = _factory.Services.GetRequiredService<IConfiguration>();

        // No need check point to reset state of db, because using InMemory DB
        //_checkpoint = new Checkpoint
        //{
        //    TablesToIgnore = new[] { new Respawn.Graph.Table("__EFMigrationsHistory") }
        //};
    }

    public static async Task<TResponse> SendAsync<TResponse>(IRequest<TResponse> request)
    {
        using var scope = _scopeFactory.CreateScope();

        var mediator = scope.ServiceProvider.GetRequiredService<ISender>();

        return await mediator.Send(request);
    }


    public static async Task ResetState()
    {
        // No need check point to reset state of db, because using InMemory DB
        //await _checkpoint.Reset("MyTestDB");
        _currentRole = null;
        _currentUserId = null;
    }

    public static async Task<TEntity?> FindAsync<TEntity>(params object[] keyValues)
        where TEntity : class
    {
        using var scope = _scopeFactory.CreateScope();

        var context = scope.ServiceProvider.GetRequiredService<ManabieDbContext>();

        return await context.FindAsync<TEntity>(keyValues);
    }

    public static async Task AddAsync<TEntity>(TEntity entity)
        where TEntity : class
    {
        using var scope = _scopeFactory.CreateScope();

        var context = scope.ServiceProvider.GetRequiredService<ManabieDbContext>();

        context.Add(entity);

        await context.SaveChangesAsync();
    }

    public static async Task AddRangeAsync<TEntity>(IList<TEntity> entities)
        where TEntity : class
    {
        using var scope = _scopeFactory.CreateScope();

        var context = scope.ServiceProvider.GetRequiredService<ManabieDbContext>();

        context.AddRange(entities);

        await context.SaveChangesAsync();
    }

    public static async Task<int> CountAsync<TEntity>() where TEntity : class
    {
        using var scope = _scopeFactory.CreateScope();

        var context = scope.ServiceProvider.GetRequiredService<ManabieDbContext>();

        return await context.Set<TEntity>().CountAsync();
    }

    public static async Task<(string, string)> RunAsDefaultUserAsync()
    {
        return await RunAsUserAsync("test@local", "Testing1234!", new[] { "User" });
    }

    public static async Task<(string,string)> RunAsAdministratorAsync()
    {
        return await RunAsUserAsync("administrator@local", "Administrator1234!", new[] { "Administrator" });
    }

    public static async Task<(string,string)> RunAsUserAsync(string userName, string password, string[] roles)
    {
        using var scope = _scopeFactory.CreateScope();

        _currentUserId = Guid.NewGuid().ToString();
        _currentRole = roles.First();
        return (_currentUserId, _currentRole);
    }

    [OneTimeTearDown]
    public void RunAfterAnyTests()
    {
    }
}
