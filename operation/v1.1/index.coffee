express = require 'express'
routes  = require './routes'
app     = express()

app.use '/tasks', routes['task']

module.exports = app
