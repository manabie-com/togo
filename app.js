/*
Implement one single API which accepts a todo task and records it
There is a maximum limit of N tasks per user that can be added per day.
Different users can have different maximum daily limit.
Write integration (functional) tests
Write unit tests
Choose a suitable architecture to make your code simple, organizable, and maintainable
*/
const express = require('express')
const app = express()
const taskRoutes = require('./routes/task')

app.use(taskRoutes)

app.listen(3000)

module.exports = app
