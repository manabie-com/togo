using System;
using System.Configuration;

using Manabie.Api.Entities;

using Microsoft.CodeAnalysis.CSharp.Syntax;
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.ChangeTracking;
using Microsoft.EntityFrameworkCore.Infrastructure.Internal;
using Microsoft.Net.Http.Headers;

namespace Manabie.Api.Insfrastructures;

public class ManaDbContext : DbContext
{
    public ManaDbContext(DbContextOptions options) : base(options)
    {

    }

    public DbSet<User> Users { get; set; }
    public DbSet<Entities.Task> Tasks { get; set; }

    protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
    {

    }

    protected override void OnModelCreating(ModelBuilder modelBuilder)
    {
        modelBuilder.Entity<User>()
            .HasIndex(n => n.Username)
            .IsUnique();

        modelBuilder.Entity<User>().HasData(new User(1,
                                      "Cuong",
                                      "Nguyen Sy Manh",
                                      "cuongnsm",
                                      5,
                                      "pBNZVHvtfjY3V13BsvZZvl2EdkJPq5xCgJomYUtQdvM="));
    }

    public override int SaveChanges()
    {
        foreach (var item in ChangeTracker.Entries<EntityBase>())
        {
            switch (item.State)
            {
                case EntityState.Added:
                    item.Entity.CreatedAt = DateTime.UtcNow;
                    break;
                default:
                    break;
            }
        }
        return base.SaveChanges();
    }
}
