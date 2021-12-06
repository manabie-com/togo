const userModel = require('../models/user.model');
const bcrypt = require("bcrypt");

module.exports = {
  validateExistAccount: async function (userName) {
    try {
      let account = await userModel.findOne({userName: userName});
      if(account) 
        return true;
      return false;
    } catch (err) {
      throw new Error (err.message);
    }
  },

  createAccount: async function createAccount(userName, password) {
    try {
      const hashPassword = await bcrypt.hash(password, 10);
      const newUser = await userModel.create({
        userName:userName,
        password: hashPassword
      });
      return newUser;
    } catch (err) {
      throw new Error(err.message);
    }
  },

  passwordCompare: async function (userName, password) {
    try {
      const account = await userModel.findOne({userName: userName});
      let passwordMatch = await bcrypt.compare(password, account.password);
      if (!passwordMatch) {
        throw new Error("Wrong username or password");
      }
      return account._id
    } catch (err) {
      throw new Error(err.message);
    }
  }
}