const ValidatorJSClass = require('validatorjs');
const Message = require('../constants/Messages');

/**
 * An adapter of validatorjs library.
 */
class Validator {
  /**
   * Ensure the data matches all rules.
   *
   * @param {object} data
   * @param {object} rules
   * @param {object} messages
   * @return {void}
   *
   * @throws {Error}
   */
  validate(data, rules, messages = {}) {
    const v = new ValidatorJSClass(data, rules, messages);

    if (v.fails()) {
      const error = new Error(Message.UNPROCESSABLE_ENTITY);
      error.errors = {};
      for (const field in rules) {
        if (v.errors.has(field)) {
          error.errors[field] = v.errors.first(field);
        }
      }
      throw error;
    }
  }
}

module.exports = new Validator();
