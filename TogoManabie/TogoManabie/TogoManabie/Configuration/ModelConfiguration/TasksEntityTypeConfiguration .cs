using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using TogoManabie.Models;

namespace TogoManabie.Configuration.ModelConfiguration
{
    public class TasksEntityTypeConfiguration : IEntityTypeConfiguration<Models.Task>
    {
        public void Configure(EntityTypeBuilder<Models.Task> builder)
        {
            builder
            .ToTable("Tasks");

            builder
            .HasKey(b => b.id);

            builder
            .Property(b => b.content)
            .IsRequired();

            builder
            .Property(b => b.created_date)
            .HasColumnType("Date")
            .HasDefaultValueSql("GetDate()");

            builder
            .HasOne(b => b.user)
            .WithMany(u => u.tasks)
            .HasForeignKey(b => b.user_id);
        }
    }
}
