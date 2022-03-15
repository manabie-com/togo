USE TodoItemDb;

CREATE TABLE [dbo].[Users](
	[ID] [bigint] IDENTITY(1,1) PRIMARY KEY NOT NULL,
	[FirstName] [NVARCHAR](MAX) NOT NULL,
	[LastName] [NVARCHAR](MAX) NOT NULL,
	[DailyTaskLimit] [int] NOT NULL
);

CREATE TABLE [dbo].[TodoItems](
	[ID] [bigint] IDENTITY(1,1) NOT NULL,
	[Name] [NVARCHAR](MAX) NOT NULL,
	[Description] [NVARCHAR](MAX) NOT NULL,
	[DateCreatedUTC] [DATETIME] NOT NULL,
	[DateModifiedUTC] [DATETIME] NOT NULL,
	[UserID] [bigint] REFERENCES Users(ID)
);

---

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

---

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

--

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

--

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

--

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

--

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