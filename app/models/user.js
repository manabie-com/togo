const constant = require('../common/constant')

module.exports = (sequelize, Sequelize) => {
  const User = sequelize.define('user', {
    userId: {
      allowNull: false,
      autoIncrement: true,
      primaryKey: true,
      type: Sequelize.INTEGER
    },
    name: {
      type: Sequelize.STRING
    },
    email: {
      type: Sequelize.STRING
    },
    tasksPerDay: {
      defaultValue: 0,
      type: Sequelize.INTEGER
    },
    timezone: {
      type: Sequelize.STRING
    },
  })

  return User
}
