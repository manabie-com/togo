'use strict';
const Utils = require(`${process.cwd()}/libs/Utils`);

module.exports = (sequelize, DataTypes) => {
  const Users = sequelize.define('users', {
    name: {type: DataTypes.STRING},
    email: {type: DataTypes.STRING},
    password: {type: DataTypes.STRING},
    status: {type: DataTypes.STRING},
    role: {type: DataTypes.ENUM('admin','manager','user')},
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

  return Users;
};