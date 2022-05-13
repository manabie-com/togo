using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using ToDoApp.DTO.Entity;

namespace ToDoApp.DTO
{
    public class DatabaseContext : DbContext
    {
        public DatabaseContext(DbContextOptions<DatabaseContext> options) : base(options)
        {
           
        }

        public DbSet<ToDo> Todos { get; set; }
        public DbSet<User> Users { get; set; }
    }
}
