module.exports =
  type: "object"
  properties:
    userName:
      type: "string"
      minLength: 6
      maxLength: 100
    email:
      type: "string"
      minLength: 6
      maxLength: 100
    password:
      type: "string"
      minLength: 6
      maxLength: 100
    dailyTaskLimit:
      type: "number"
  required: ["userName", "email", "password"]
