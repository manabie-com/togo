module.exports =
  "v1.1":
    GET: []
    POST: [
      {
        route: "/users/v1.1/sign-up"
        pattern: new RegExp("^\/users\/v1.1\/sign-up(?:\/(?=$))?$", "i")
        schema: "SignUpSchema"
      }
    ]
    PUT: []
    DELETE: []
