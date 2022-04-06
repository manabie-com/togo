module.exports =
  type: "object"
  properties:
    userId:
      type: "string"
    taskId:
      type: ["string", "null"]
    taskName:
      type: "string"
      minLength: 6
      maxLength: 100
    taskCode:
      type: "string"
      minLength: 6
      maxLength: 100
    taskDescription:
      type: "string"
    status:
      type: "string"
  required: ["userId", "taskId", "taskName", "taskDescription", "status", "taskCode"]
