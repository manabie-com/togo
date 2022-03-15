USE TodoItemDb;
GO
CREATE PROCEDURE [dbo].[DeleteUser]
(
    @userID BIGINT
)
AS
BEGIN
    SET NOCOUNT ON;
	DELETE
	FROM [dbo].[Users]
	WHERE ID = @userID
END;
GO