USE TodoItemDb;

CREATE TABLE [dbo].[Users](
	[ID] [bigint] IDENTITY(1,1) PRIMARY KEY NOT NULL,
	[FirstName] [NVARCHAR](MAX) NOT NULL,
	[LastName] [NVARCHAR](MAX) NOT NULL,
	[DailyTaskLimit] [int] NOT NULL
);