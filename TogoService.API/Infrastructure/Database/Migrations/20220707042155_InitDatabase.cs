using System;
using Microsoft.EntityFrameworkCore.Migrations;

namespace TogoService.API.Infrastructure.Database.Migrations
{
    public partial class InitDatabase : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "User",
                columns: table => new
                {
                    id = table.Column<Guid>(type: "TEXT", nullable: false),
                    name = table.Column<string>(type: "TEXT", maxLength: 500, nullable: false),
                    maxDailyTasks = table.Column<uint>(type: "INTEGER", nullable: false),
                    isDeleted = table.Column<bool>(type: "INTEGER", nullable: false),
                    updatedAt = table.Column<DateTime>(type: "TEXT", nullable: false),
                    createdAt = table.Column<DateTime>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_User", x => x.id);
                });

            migrationBuilder.CreateTable(
                name: "TodoTask",
                columns: table => new
                {
                    id = table.Column<Guid>(type: "TEXT", nullable: false),
                    name = table.Column<string>(type: "TEXT", maxLength: 500, nullable: false),
                    description = table.Column<string>(type: "TEXT", maxLength: 500, nullable: true),
                    todoDay = table.Column<DateTime>(type: "TEXT", nullable: false),
                    userId = table.Column<Guid>(type: "TEXT", nullable: false),
                    isDeleted = table.Column<bool>(type: "INTEGER", nullable: false),
                    updatedAt = table.Column<DateTime>(type: "TEXT", nullable: false),
                    createdAt = table.Column<DateTime>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_TodoTask", x => x.id);
                    table.ForeignKey(
                        name: "FK_TodoTask_User_userId",
                        column: x => x.userId,
                        principalTable: "User",
                        principalColumn: "id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateIndex(
                name: "IX_TodoTask_isDeleted",
                table: "TodoTask",
                column: "isDeleted");

            migrationBuilder.CreateIndex(
                name: "IX_TodoTask_userId",
                table: "TodoTask",
                column: "userId");

            migrationBuilder.CreateIndex(
                name: "IX_User_isDeleted",
                table: "User",
                column: "isDeleted");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "TodoTask");

            migrationBuilder.DropTable(
                name: "User");
        }
    }
}
