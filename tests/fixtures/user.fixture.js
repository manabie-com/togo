const bcrypt = require('bcryptjs');
const faker = require('faker');
const users = require('../../backend/models');

const password = 'password1';
const salt = bcrypt.genSaltSync(8);
const hashedPassword = bcrypt.hashSync(password, salt);

const userOne = {
  name: faker.name.findName(),
  email: faker.internet.email().toLowerCase(),
  password,
};

const userTwo = {
  name: faker.name.findName(),
  email: faker.internet.email().toLowerCase(),
  password,
};

const insertUsers = async (us) => {
  await users.insertMany(us.map((user) => ({ ...user, password: hashedPassword })));
};

module.exports = {
  userOne,
  userTwo,
  insertUsers,
};
