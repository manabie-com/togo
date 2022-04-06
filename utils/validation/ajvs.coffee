_          = require 'lodash'
Ajv        = require 'ajv'

ajv = new Ajv.default { allErrors: true, allowUnionTypes: true }

module.exports = ajv
