USE TodoItemDb;
GO
CREATE PROCEDURE [dbo].[GetUser]
(
    @userID BIGINT
)
AS
BEGIN
    SET NOCOUNT ON;
	SELECT ID, FirstName, LastName, DailyTaskLimit
	FROM [dbo].[Users]
	WHERE ID = @userID
END;
GO