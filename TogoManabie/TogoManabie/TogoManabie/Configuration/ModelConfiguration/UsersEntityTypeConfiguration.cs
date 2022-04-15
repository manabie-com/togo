using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Models;

namespace TogoManabie.Configuration.ModelConfiguration
{
    public class UsersEntityTypeConfiguration : IEntityTypeConfiguration<User>
    {
        public void Configure(EntityTypeBuilder<User> builder)
        {
            builder
            .ToTable("Users");

            builder
            .HasKey(b => b.id);

            builder
            .Property(b => b.passWord)
            .IsRequired();

            builder
            .Property(b => b.maxTodo)
            .IsRequired();
        }
    }
}
