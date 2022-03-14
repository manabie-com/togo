'use strict';
const Hash = require('../../libs/Hash');

module.exports = {
  async up (queryInterface, Sequelize) {
    /**
     * Add seed commands here.
     *
     * Example:
     * await queryInterface.bulkInsert('People', [{
     *   name: 'John Doe',
     *   isBetaMember: false
     * }], {});
    */
    return queryInterface.bulkInsert('users', [{
        name: 'Admin',
        email: 'admin@gmail.com',
        password: Hash.make('admin123'),
        status: 'active',
        role: 'admin',
        created_at: 1646873640,
        updated_at: 1646873640
    }]);
  },

  async down (queryInterface, Sequelize) {
    /**
     * Add commands to revert seed here.
     *
     * Example:
     * await queryInterface.bulkDelete('People', null, {});
     */
    return queryInterface.bulkDelete('users', null, {});
  }
};
