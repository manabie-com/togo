/* jshint indent: 2 */

import { Model, Sequelize } from 'sequelize';

export default class Tasks extends Model {
  static init(sequelize, DataTypes) {
  super.init({
    task_id: {
      autoIncrement: true,
      type: DataTypes.INTEGER,
      allowNull: false,
      primaryKey: true
    },
    user_id: {
      type: DataTypes.INTEGER,
      allowNull: false,
      defaultValue: 0,
      references: {
        model: 'users',
        key: 'user_id'
      }
    },
    content: {
      type: DataTypes.STRING,
      allowNull: false
    },
    created_date: {
      type: DataTypes.DATE,
      allowNull: true,
      defaultValue: Sequelize.literal('CURRENT_TIMESTAMP')
    }
  }, {
    sequelize,
    tableName: 'tasks',
    schema: 'public',
    timestamps: false,
    indexes: [
      {
        name: "tasks_pkey",
        unique: true,
        fields: [
          { name: "task_id" },
        ]
      },
    ]
  });
  return Tasks;
  }
}
