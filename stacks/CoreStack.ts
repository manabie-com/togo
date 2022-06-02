import { StackContext, Api, Table } from "@serverless-stack/resources";

export function CoreStack({ stack }: StackContext) {
  const isProd = stack.stage === "prod";
  const logLevel = process.env.LOG_LEVEL || isProd ? "INFO" : "DEBUG";

  const usersTable = new Table(stack, "UsersTable", {
    fields: { id: "string" },
    primaryIndex: { partitionKey: "id" },
    cdk: { table: { pointInTimeRecovery: isProd } },
  });

  const tasksTable = new Table(stack, "TasksTable", {
    fields: { user_id: "string", id: "string", due_date: "number" },
    primaryIndex: { partitionKey: "user_id", sortKey: "id" },
    globalIndexes: {
      userIdDueDateIndex: { partitionKey: "user_id", sortKey: "due_date" },
    },
    cdk: { table: { pointInTimeRecovery: isProd } },
  });

  new Api(stack, "Api", {
    defaults: {
      function: {
        environment: {
          LOG_LEVEL: logLevel,
          USERS_TABLE_NAME: usersTable.tableName,
          TASKS_TABLE_NAME: tasksTable.tableName,
        },
      },
    },
    accessLog: {
      retention: isProd ? "three_months" : "one_day",
    },
    routes: {
      "POST /users": "functions/users/create/lambda.go",
      "POST /users/{userId}/tasks": "functions/tasks/create/lambda.go",
    },
  });
}
