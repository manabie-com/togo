express = require 'express'
routes  = require './routes'
app     = express()

app.use '/', routes['user']

module.exports = app
