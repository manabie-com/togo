USE TodoItemDb;
GO
CREATE PROCEDURE [dbo].[AddTodo]
(
    @name NVARCHAR(MAX),
    @description NVARCHAR(MAX),
    @dateCreatedUTC DATETIME,
    @dateModifiedUTC DATETIME,
	@userId BIGINT
)
AS
BEGIN
    SET NOCOUNT ON;
	DECLARE @transactionDate DATETIME = GETUTCDATE();
	INSERT INTO [dbo].[TodoItems]
    (
        [Name],
        [Description],
        [DateCreatedUTC],
        [DateModifiedUTC],
		[UserID]
    )
    OUTPUT
        @transactionDate,
        INSERTED.[ID],
        INSERTED.[Name],
        INSERTED.[Description],
		INSERTED.[DateCreatedUTC],
        INSERTED.[DateModifiedUTC],
		INSERTED.[UserID]
    VALUES
    (@name, @description, @transactionDate, @transactionDate, @userId);
END;
GO