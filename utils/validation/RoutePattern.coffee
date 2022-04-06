patterns        = require './patterns'

class RoutePattern
  constructor: () ->
    @routePattern = patterns

  getRoutePatterns: (module, version, method) ->
    # console.log "module", module
    # console.log "version", version
    # console.log "method", method

    if not module? or not version? or not method? then return []

    return @routePattern[module]?[version]?[method] || []

module.exports = RoutePattern
