requireAll = require 'require-all'

module.exports = requireAll(
  dirname: __dirname
  filter: /(.+Model)\.coffee$/
  resolve: (Model) ->
    return new Model()
)
