'use strict';

module.exports = {
    async up(queryInterface, Sequelize) {
        await queryInterface.bulkInsert('Role', [{
            code: 'Admin',
            status: true
        }], {});
    },

    async down(queryInterface, Sequelize) {
        await queryInterface.bulkDelete('Role', null, {});
    }
};