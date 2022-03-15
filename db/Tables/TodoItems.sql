USE TodoItemDb;

CREATE TABLE [dbo].[TodoItems](
	[ID] [bigint] IDENTITY(1,1) NOT NULL,
	[Name] [NVARCHAR](MAX) NOT NULL,
	[Description] [NVARCHAR](MAX) NOT NULL,
	[DateCreatedUTC] [DATETIME] NOT NULL,
	[DateModifiedUTC] [DATETIME] NOT NULL,
	[UserID] [bigint] REFERENCES Users(ID)
);