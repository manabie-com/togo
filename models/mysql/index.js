'use strict';

const fs = require('fs');
const path = require('path');
const Sequelize = require('sequelize');
const basename= path.basename(__filename);
const {mysql} = require(`${process.cwd()}/configs`);
const db = {};
let sequelize = null;

sequelize = new Sequelize(mysql.database,
    mysql.username,
    mysql.password,
    mysql);

fs
    .readdirSync(__dirname)
    .filter((file) => {
      return (file.indexOf('.') !== 0)
            && (file !== basename)
            && (file.slice(-3) === '.js');
    })
    .forEach((file) => {
      // eslint-disable-next-line max-len
      const model = require(path.join(__dirname, file))(sequelize, Sequelize.DataTypes);
      db[model.name] = model;
    });

Object.keys(db).forEach((modelName) => {
  if (db[modelName].associate) {
    db[modelName].associate(db);
  }
});

db.sequelize = sequelize;
db.Sequelize = Sequelize;
db.Op = Sequelize.Op;

sequelize
    .authenticate()
    .then(() => {
      console.log('Connection MySQL has been established successfully.');
    })
    .catch((err) => {
      console.error('Unable to connect to the database:', err);
    });

module.exports = db;
