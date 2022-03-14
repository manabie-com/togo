process.env.NODE_ENV = 'test';
process.env.MYSQL_DATABASE = 'manabie-test';

const {users} = require('../models').mysql;
let chai = require('chai');
let chaiHttp = require('chai-http');
let server = require('../app');
let should = chai.should();
const bcrypt = require('bcrypt');
const config = require('../configs');
const saltRounds = config.getENV('SALT_ROUNDS');

describe('Auth', () => {
  before((done) => {
      //Before test we empty the database in your case
      done();
  });

  describe('Init User Sample', () => {
    it('Init User Sample', (done) => {
      const salt = bcrypt.genSaltSync(parseInt(saltRounds));
      const hash = bcrypt.hashSync("admin123", salt);
      users.create({
        name: "User Test",
        email: "test@gmail.com",
        password: hash,
        status: "active",
        role: "member"
      });
      done();
    });
  });
});