
using Microsoft.EntityFrameworkCore;
using Microsoft.EntityFrameworkCore.Metadata.Builders;

public class ToDoConfiguration : IEntityTypeConfiguration<ToDo>
{
    public void Configure(EntityTypeBuilder<ToDo> toDoConfiguration)
    {
        toDoConfiguration.HasOne(x => x.User)
        .WithMany(y => y.ToDos)
        .HasForeignKey(y => y.UserId)
        .OnDelete(DeleteBehavior.Cascade);
    }
}
