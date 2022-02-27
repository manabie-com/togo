using System;
using Microsoft.EntityFrameworkCore.Migrations;

namespace TODO.Repositories.Migrations
{
    public partial class init : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "TodoStatus",
                columns: table => new
                {
                    TodoStatusId = table.Column<int>(type: "int", nullable: false),
                    StatusName = table.Column<string>(type: "nvarchar(max)", nullable: true),
                    StatusDescription = table.Column<string>(type: "nvarchar(max)", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_TodoStatus", x => x.TodoStatusId);
                });

            migrationBuilder.CreateTable(
                name: "User",
                columns: table => new
                {
                    UserId = table.Column<int>(type: "int", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    LastName = table.Column<string>(type: "nvarchar(max)", nullable: true),
                    FirstName = table.Column<string>(type: "nvarchar(max)", nullable: true),
                    MiddleName = table.Column<string>(type: "nvarchar(max)", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_User", x => x.UserId);
                });

            migrationBuilder.CreateTable(
                name: "Todo",
                columns: table => new
                {
                    TodoId = table.Column<int>(type: "int", nullable: false)
                        .Annotation("SqlServer:Identity", "1, 1"),
                    UserId = table.Column<int>(type: "int", nullable: false),
                    StatusId = table.Column<int>(type: "int", nullable: false),
                    TodoName = table.Column<string>(type: "nvarchar(max)", nullable: true),
                    TodoDescription = table.Column<string>(type: "nvarchar(max)", nullable: true),
                    Priority = table.Column<int>(type: "int", nullable: false),
                    DateCreated = table.Column<DateTime>(type: "datetime2", nullable: true),
                    DateModified = table.Column<DateTime>(type: "datetime2", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Todo", x => x.TodoId);
                    table.ForeignKey(
                        name: "FK_Todo_TodoStatus_StatusId",
                        column: x => x.StatusId,
                        principalTable: "TodoStatus",
                        principalColumn: "TodoStatusId",
                        onDelete: ReferentialAction.Cascade);
                    table.ForeignKey(
                        name: "FK_Todo_User_UserId",
                        column: x => x.UserId,
                        principalTable: "User",
                        principalColumn: "UserId",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "UserTodoConfig",
                columns: table => new
                {
                    UserId = table.Column<int>(type: "int", nullable: false),
                    DailyTaskLimit = table.Column<int>(type: "int", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_UserTodoConfig", x => x.UserId);
                    table.ForeignKey(
                        name: "FK_UserTodoConfig_User_UserId",
                        column: x => x.UserId,
                        principalTable: "User",
                        principalColumn: "UserId",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.InsertData(
                table: "TodoStatus",
                columns: new[] { "TodoStatusId", "StatusDescription", "StatusName" },
                values: new object[,]
                {
                    { 0, null, "TO DO" },
                    { 1, null, "DONE" },
                    { 2, null, "IN PROGRESS" }
                });

            migrationBuilder.InsertData(
                table: "User",
                columns: new[] { "UserId", "FirstName", "LastName", "MiddleName" },
                values: new object[,]
                {
                    { 1, "Michael", "Jordan", null },
                    { 2, "Isiah", "Thomas", null }
                });

            migrationBuilder.InsertData(
                table: "UserTodoConfig",
                columns: new[] { "UserId", "DailyTaskLimit" },
                values: new object[] { 1, 10 });

            migrationBuilder.InsertData(
                table: "UserTodoConfig",
                columns: new[] { "UserId", "DailyTaskLimit" },
                values: new object[] { 2, 5 });

            migrationBuilder.CreateIndex(
                name: "IX_Todo_StatusId",
                table: "Todo",
                column: "StatusId");

            migrationBuilder.CreateIndex(
                name: "IX_Todo_UserId",
                table: "Todo",
                column: "UserId");
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "Todo");

            migrationBuilder.DropTable(
                name: "UserTodoConfig");

            migrationBuilder.DropTable(
                name: "TodoStatus");

            migrationBuilder.DropTable(
                name: "User");
        }
    }
}
