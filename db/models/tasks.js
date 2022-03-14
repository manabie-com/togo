'use strict';
const {
  Model
} = require('sequelize');
module.exports = (sequelize, DataTypes) => {
  class tasks extends Model {
    /**
     * Helper method for defining associations.
     * This method is not a part of Sequelize lifecycle.
     * The `models/index` file will call this method automatically.
     */
    static associate(models) {
      // define association here
    }
  }
  tasks.init({
    name: DataTypes.STRING,
    description: DataTypes.STRING,
    status: DataTypes.STRING,
    estimated_time: DataTypes.INTEGER,
    due_date: DataTypes.INTEGER,
    user_id: DataTypes.INTEGER,
    created_at: DataTypes.INTEGER,
    updated_at: DataTypes.INTEGER
  }, {
    sequelize,
    modelName: 'tasks',
  });
  return tasks;
};