const { expect } = require("chai");
const { matches, encrypt, decrypt, mandatory } = require("../../src/utils/helper.util");

describe("Test helper.util.js", () => {
  describe("Test function matches", () => {
    it("Should return true when plainText and encryptedText are equal", () => {
      const plainText = "Hello";
      const encryptedText = encrypt(plainText);
      expect(matches(plainText, encryptedText)).to.be.true;
    });
    it("Should return false when plainText and encryptedText are not equal", () => {
      const plainText = "Hello";
      const encryptedText = encrypt(plainText);
      expect(matches("hello", encryptedText)).to.be.false;
    });
  });
  describe("Test function encrypt", () => {
    it("Should return encrypted text", () => {
      const plainText = "Hello";
      const encryptedText = encrypt(plainText);
      expect(encryptedText).to.be.ok;
    });
  });
  describe("Test function decrypt", () => {
    it("Should return decrypted text", () => {
      const plainText = "Hello";
      const encryptedText = encrypt(plainText);
      expect(decrypt(encryptedText)).to.equal(plainText);
    });
  });
  describe("Test function mandatory", () => {
    it("Should throw error when param is undefined", () => {
      expect(() => {
        mandatory();
      }).to.throw("Parameter is required");
    });
    it("Should throw error when param is null", () => {
      expect(() => {
        mandatory(null);
      }).to.throw("Parameter is required");
    });
    it("Should throw error when param is empty string", () => {
      expect(() => {
        mandatory("");
      }).to.throw("Parameter is required");
    });
    it("Should not throw error when param is string", () => {
      expect(() => {
        mandatory("Hello");
      }).not.to.throw("Parameter is required");
    });
  });
});
