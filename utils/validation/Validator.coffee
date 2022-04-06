_                             = require 'lodash'
RoutePattern                  = require './RoutePattern.coffee'
schemas                       = require './schemas'
code                          = require "#{process.cwd()}/config/code.coffee"
routePattern                  = new RoutePattern()
ajvs                          = require './ajvs.coffee'


class Validator
  constructor: (req) ->
    if not req? then throw new Error("req is required")

    @req = _.cloneDeep req
    @method = @req.method
    @path = @req.path
    @schema = {}
    @params = {}

    @module = @getModule()
    @version = @getVersion()

    @ajv = ajvs
    @routePattern = @getRoutePattern()

    if @routePattern?
      @params = @buildParams(@req, @routePattern)
      @schema = @getSchema()

    @validate = @ajv.compile(@schema)

  getVersion: () =>
    if @path.split('/').length > 2
      return @path.split('/')[2]

    return null

  getModule: () =>
    if @path.split('/').length > 1
      return @path.split('/')[1]

    return null

  getSchema: () =>
    return schemas[@module]?[@version]?[@routePattern.schema]

  buildParams: (req, routePattern) ->
    params = {}
    # get req.params
    keys = []
    values = routePattern.pattern.exec(req.path)

    keys.map (key, i) ->
      params[key.name] = values[i + 1]
    switch req.method
      when "POST"
        params = _.extend req.body, req.query
      when "PUT"
        params = _.extend params, req.body
      when "GET"
        params = _.extend params, req.query
      when "DELETE"
        params = _.extend params, req.body, req.query
    return params

  getRoutePattern: () =>
    @routePatterns = routePattern.getRoutePatterns(@module, @version, @method)
    selectedRoutePattern = null

    for rp in @routePatterns
      if rp.pattern.test(@path)
        selectedRoutePattern = rp
        break
    return selectedRoutePattern

  validateParams: () =>
    valid = @validate(@params)

    codeErrors = _.map @validate.errors, (error) ->
      if error.keyword is 'enum'
        allowedValues = ": [ #{error?.params?.allowedValues || ""} ]"
        error.message = error.message + allowedValues

      return error.keyword

    codeResult = 0
    reason = @ajv.errorsText(@validate.errors)

    if codeErrors.indexOf("required") isnt -1
      codeResult = code.CODE_MISS_PARAMETER

    else if codeErrors.indexOf("type") isnt -1 or codeErrors.indexOf("maximum") isnt -1 or codeErrors.indexOf("minimum") isnt -1
      codeResult = code.CODE_ERROR

    else if codeErrors.indexOf("enum") isnt -1
      codeResult = code.CODE_ERROR

    else if codeErrors.indexOf("maxItems") isnt -1 or codeErrors.indexOf("maxLength") isnt -1 or codeErrors.indexOf("checkFieldSearch") isnt -1
      codeResult  = code.CODE_ERROR

    else if codeErrors.indexOf("minItems") isnt -1 or codeErrors.indexOf("minLength") isnt -1
      codeResult = code.CODE_ERROR

    else if codeErrors.indexOf("format") isnt -1 or codeErrors.indexOf("formatURL") isnt -1 or codeErrors.indexOf("formatDate") isnt -1 or codeErrors.indexOf("formatDateForChartTotal") isnt -1
      codeResult = code.CODE_ERROR

    else if codeErrors.indexOf("canNotStartDateAfterEndDate") isnt -1
      codeResult = code.CODE_ERROR

    else if codeErrors.indexOf("checkTimeForSegments") isnt -1
      codeResult = code.CODE_ERROR

    else if codeErrors.indexOf("checkRuleFilter") isnt -1
      codeResult = code.CODE_ERROR

    else if codeErrors.indexOf("checkName") isnt -1
      codeResult = code.CODE_ERROR

    return {
      valid: valid
      code: codeResult
      reason: reason
    }

module.exports = Validator
