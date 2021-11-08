/* jshint indent: 2 */

import { Model, Sequelize } from 'sequelize';

export default class Users extends Model {
  static init(sequelize, DataTypes) {
  super.init({
    user_id: {
      autoIncrement: true,
      type: DataTypes.INTEGER,
      allowNull: false,
      primaryKey: true
    },
    username: {
      type: DataTypes.STRING,
      allowNull: false,
      defaultValue: "",
      unique: "username"
    },
    password: {
      type: DataTypes.STRING,
      allowNull: false,
      defaultValue: ""
    },
    max_todo: {
      type: DataTypes.INTEGER,
      allowNull: false,
      defaultValue: 5
    }
  }, {
    sequelize,
    tableName: 'users',
    schema: 'public',
    timestamps: false,
    indexes: [
      {
        name: "users_pkey",
        unique: true,
        fields: [
          { name: "user_id" },
        ]
      },
      {
        name: "username",
        unique: true,
        fields: [
          { name: "username" },
        ]
      },
    ]
  });
  return Users;
  }
}
