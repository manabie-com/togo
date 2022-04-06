express        = require 'express'
bodyParser     = require 'body-parser'
code           = require "#{process.cwd()}/config/code"
config         = require  "#{process.cwd()}/config/development"
mongoose       = require 'mongoose'
middlewares    = require './middlewares'
validateParams = require './middlewares/validateParams'
port           = config?.app?.port || 8004
app            = express()

mongoose.connect config.mongoConnectionString
db = mongoose.connection
db.on "error", console.error.bind(console, "connection error: ")
db.once "open", () -> console.log("Connected successfully")

app.use bodyParser.json({ limit: '50mb' })
app.use bodyParser.urlencoded({ limit: '50mb', extended: false })

app.use middlewares.validateParams

app.use '/users', require './users'
app.use '/operation', require './operation'

app.get '/', (req, res) -> res.json { msg: 'Hello world' }

# catch 404 and forward to error handler
app.use (req, res, next) ->
  res.json(code: code.CODE_UNSUPPORTED_API, message: "unsupported api")

# error handler
app.use (err, req, res, next) ->
  if err?.message?.indexOf('have_no_privilege') > -1
    return res.json {
      code: code.CODE_FORBIDDEN
      message: 'No Privilege...'
      data: {}
    }
  else
    return res.json { code: code.CODE_ERROR, message: err.stack }

app.listen port, () -> console.log "Express server listening on port #{port}"

module.exports = app
