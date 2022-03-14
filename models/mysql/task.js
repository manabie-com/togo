'use strict';
const Utils = require(`${process.cwd()}/libs/Utils`);

module.exports = (sequelize, DataTypes) => {
  const Task = sequelize.define('tasks', {
    name: {type: DataTypes.STRING},
    description: {type: DataTypes.STRING},
    status: {type: DataTypes.ENUM('todo','in-process','in-review', 'done')},
    estimated_time: {
      type: DataTypes.INTEGER,
      allowNull: true,
      defaultValue: Utils.timestamps(),
    },
    due_date: {
      type: DataTypes.INTEGER,
      allowNull: true,
      defaultValue: Utils.timestamps(),
    },
    user_id: {type: DataTypes.INTEGER},
    created_at: {
      type: DataTypes.INTEGER,
      allowNull: true,
      defaultValue: Utils.timestamps(),
    },
    updated_at: {
      type: DataTypes.INTEGER,
      allowNull: true,
      defaultValue: Utils.timestamps(),
    },
  },{
    timestamps: false
  });

  return Task;
};