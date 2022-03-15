USE TodoItemDb;
GO
CREATE PROCEDURE [dbo].[GetUserTodoItems]
(
    @userID BIGINT
)
AS
BEGIN
    SET NOCOUNT ON;
	SELECT TodoItems.ID, TodoItems.Name, TodoItems.Description, TodoItems.DateCreatedUTC, TodoItems.DateModifiedUTC, TodoItems.UserID, Users.FirstName, Users.LastName, Users.DailyTaskLimit
	FROM [dbo].[TodoItems]
	INNER JOIN Users ON Users.ID = TodoItems.UserID
	WHERE TodoItems.UserID = @userID;
END;
GO