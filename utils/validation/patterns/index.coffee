requireAll          = require 'require-all'
config              = require 'config'

module.exports = requireAll(
  dirname : __dirname
  filter  : /(.+)\.coffee$/
)
