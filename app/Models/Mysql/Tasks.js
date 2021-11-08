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
    content: {
      type: DataTypes.TEXT,
      allowNull: true
    },
    user_id: {
      type: DataTypes.INTEGER.UNSIGNED,
      allowNull: false,
      references: {
        model: 'users',
        key: 'user_id'
      },
      unique: "FK_tasks_users"
    },
    created_date: {
      type: DataTypes.DATE,
      allowNull: true,
      defaultValue: Sequelize.literal('CURRENT_TIMESTAMP')
    }
  }, {
    sequelize,
    tableName: 'tasks',
    timestamps: false,
    indexes: [
      {
        name: "PRIMARY",
        unique: true,
        using: "BTREE",
        fields: [
          { name: "task_id" },
        ]
      },
      {
        name: "FK_tasks_users",
        using: "BTREE",
        fields: [
          { name: "user_id" },
        ]
      },
    ]
  });
  return Tasks;
  }
}
