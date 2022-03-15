USE TodoItemDb;
GO

CREATE PROCEDURE [dbo].[AddUser]
(
    @firstName NVARCHAR(MAX),
    @lastName NVARCHAR(MAX),
    @dailyTaskLimit INT
)
AS
BEGIN
    SET NOCOUNT ON;
	INSERT INTO [dbo].[Users]
    (
        [FirstName],
        [LastName],
        [DailyTaskLimit]
    )
    OUTPUT
        INSERTED.[ID],
        INSERTED.[FirstName],
        INSERTED.[LastName],
        INSERTED.[DailyTaskLimit]
    VALUES
    (@firstName, @lastName, @dailyTaskLimit);
END;
GO