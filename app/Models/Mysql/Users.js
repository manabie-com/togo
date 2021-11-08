/* jshint indent: 2 */

import { Model } from 'sequelize';

export default class Users extends Model {
  static init(sequelize, DataTypes) {
  super.init({
    user_id: {
      autoIncrement: true,
      type: DataTypes.INTEGER.UNSIGNED,
      allowNull: false,
      primaryKey: true
    },
    username: {
      type: DataTypes.STRING(50),
      allowNull: false,
      unique: "username"
    },
    password: {
      type: DataTypes.STRING(32),
      allowNull: false
    },
    max_todo: {
      type: DataTypes.INTEGER.UNSIGNED,
      allowNull: true,
      defaultValue: 5
    }
  }, {
    sequelize,
    tableName: 'users',
    timestamps: false,
    indexes: [
      {
        name: "PRIMARY",
        unique: true,
        using: "BTREE",
        fields: [
          { name: "user_id" },
        ]
      },
      {
        name: "username",
        unique: true,
        using: "BTREE",
        fields: [
          { name: "username" },
        ]
      },
    ]
  });
  return Users;
  }
}
