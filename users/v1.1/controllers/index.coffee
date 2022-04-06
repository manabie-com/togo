requireAll = require 'require-all'

module.exports = requireAll(
  dirname: __dirname
  filter: /(.+Controller)\.coffee$/
  resolve: (Controller) ->
    return new Controller()
)
