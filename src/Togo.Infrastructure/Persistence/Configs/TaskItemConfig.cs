using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;
using Togo.Core.Entities;

namespace Togo.Infrastructure.Persistence.Configs;

public class TaskItemConfig : IEntityTypeConfiguration<TaskItem>
{
    public void Configure(EntityTypeBuilder<TaskItem> builder)
    {
        builder.HasKey(entity => entity.Id);

        builder
            .Property(entity => entity.Title)
            .IsRequired()
            .HasMaxLength(200);
        
        builder.HasIndex(entity => entity.Title);
    }
}
