using Microsoft.EntityFrameworkCore.Migrations;

namespace ToDoBackend.Migrations
{
    public partial class UserAndSettingsDataSeed : Migration
    {
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.InsertData(
                table: "Users",
                columns: new[] { "Id", "FirstName", "LastName" },
                values: new object[] { "firstuser", "Ha", "Nguyen" });

            migrationBuilder.InsertData(
                table: "Users",
                columns: new[] { "Id", "FirstName", "LastName" },
                values: new object[] { "seconduser", "Ha", "Thanh Nguyen" });

            migrationBuilder.InsertData(
                table: "UserSettings",
                columns: new[] { "UserId", "MaxTasksPerDay" },
                values: new object[] { "firstuser", 5 });

            migrationBuilder.InsertData(
                table: "UserSettings",
                columns: new[] { "UserId", "MaxTasksPerDay" },
                values: new object[] { "seconduser", 15 });
        }

        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DeleteData(
                table: "UserSettings",
                keyColumn: "UserId",
                keyValue: "firstuser");

            migrationBuilder.DeleteData(
                table: "UserSettings",
                keyColumn: "UserId",
                keyValue: "seconduser");

            migrationBuilder.DeleteData(
                table: "Users",
                keyColumn: "Id",
                keyValue: "firstuser");

            migrationBuilder.DeleteData(
                table: "Users",
                keyColumn: "Id",
                keyValue: "seconduser");
        }
    }
}
