using System;
using Microsoft.EntityFrameworkCore.Migrations;

namespace Infrastructure.Persistence.Migrations
{
    public partial class init : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "user",
                columns: table => new
                {
                    id = table.Column<string>(nullable: false),
                    created_by = table.Column<string>(nullable: false),
                    created_date = table.Column<DateTime>(nullable: false),
                    last_modified_by = table.Column<string>(nullable: true),
                    last_modified_date = table.Column<DateTime>(nullable: true),
                    password = table.Column<string>(nullable: false),
                    max_to_do = table.Column<int>(nullable: false, defaultValue: 5)
                        .Annotation("Sqlite:Autoincrement", true),
                    email = table.Column<string>(nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_user", x => x.id);
                });

            migrationBuilder.CreateTable(
                name: "task",
                columns: table => new
                {
                    id = table.Column<string>(nullable: false),
                    created_by = table.Column<string>(nullable: false),
                    create_date = table.Column<DateTime>(nullable: false),
                    last_modified_by = table.Column<string>(nullable: true),
                    last_modified_date = table.Column<DateTime>(nullable: true),
                    content = table.Column<string>(nullable: false),
                    user_id = table.Column<string>(nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_task", x => x.id);
                    table.ForeignKey(
                        name: "tasks_FK",
                        column: x => x.user_id,
                        principalTable: "user",
                        principalColumn: "id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_task_user_id",
                table: "task",
                column: "user_id");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "task");

            migrationBuilder.DropTable(
                name: "user");
        }
    }
}
