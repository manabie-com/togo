const bcrypt = require('bcrypt');
const config = require('../configs');
const saltRounds = config.getENV('SALT_ROUNDS');

module.exports = class Hash {
  /**
   * constructor
  */
  constructor() {
    this._saltRounds = parseInt(saltRounds);
  }

  /**
   * make hash
   * @param {string} value
   * @return {string}
   */
  make(value) {
    const salt = bcrypt.genSaltSync(this._saltRounds);
    const hash = bcrypt.hashSync(value, salt);
    return hash;
  }

  /**
   * check
   * @param {string} value
   * @param {string} hashedValue
   * @return {string}
   */
  check(value, hashedValue) {
    return bcrypt.compareSync(value, hashedValue);
  }
};
