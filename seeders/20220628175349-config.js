'use strict';

module.exports = {
    async up(queryInterface, Sequelize) {
        await queryInterface.bulkInsert('Config', [{
            role: 'Admin',
            limit: 5
        }], {});
    },

    async down(queryInterface, Sequelize) {
        await queryInterface.bulkDelete('Config', null, {});
    }
};