const constant = require('../common/constant')

module.exports = (sequelize, Sequelize) => {
  const Task = sequelize.define('task', {
    taskId: {
      allowNull: false,
      autoIncrement: true,
      primaryKey: true,
      type: Sequelize.INTEGER
    },
    name: {
      type: Sequelize.STRING
    },
    status: {
      type: Sequelize.STRING,
      defaultValue: constant.STATUS_NEW
    },
    localDate: {
      type: Sequelize.STRING,
      allowNull: false,
    },
  })

  Task.associate = models => {
    Task.belongsTo(models.user,
      {
        foreignKey: 'userId',
        as: 'taskAssigned'
      }
    )
  }

  return Task
}
