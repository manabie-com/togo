USE TodoItemDb;
GO
CREATE PROCEDURE [dbo].[DeleteTodo]
(
    @todoID BIGINT
)
AS
BEGIN
    SET NOCOUNT ON;
	DELETE
	FROM [dbo].[TodoItems]
	WHERE ID = @todoID
END;
GO