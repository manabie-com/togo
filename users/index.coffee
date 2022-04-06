express           = require 'express'

app = express()

app.use '/v1.1', require './v1.1'

module.exports = app
