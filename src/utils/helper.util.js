const crypto = require("crypto-js");
const secretKey = require("../configs/authen.config").jwt.secretKey;

const matches = (plainText, encryptedText) => {
  return decrypt(String(encryptedText)) === plainText;
};

const encrypt = (plainText) => {
  return crypto.AES.encrypt(plainText, secretKey).toString();
};

const decrypt = (encryptedText) => {
  const bytes = crypto.AES.decrypt(String(encryptedText), secretKey);
  return bytes.toString(crypto.enc.Utf8);
};

const mandatory = (param) => {
  if (param === "" || param === null || param === undefined) {
    throw new Error("Parameter is required");
  }
  throw new Error(`${param} is required`);
};


module.exports = { matches, encrypt, decrypt, mandatory };
