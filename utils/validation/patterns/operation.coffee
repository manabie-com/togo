module.exports =
  "v1.1":
    GET: []
    POST: [
      {
        route: "/operation/v1.1/tasks/assign"
        pattern: new RegExp("^\/operation\/v1.1\/tasks\/assign(?:\/(?=$))?$", "i")
        schema: "AssignTaskSchema"
      }
    ]
    PUT: []
    DELETE: []
