const sequelize = require('sequelize');
const fs = require('fs');
const path = require('path');
const configDB = require('../../config/database');

const modelDir = path.join(process.cwd(), 'src/infrastructure/models');
const db = {};
const {
  database: dbName,
  username: dbUsername,
  password: dbPassword,
  host: dbHost,
  port: dbPort,
} = configDB[process.env.NODE_ENV];
const sequelizeInstance = new sequelize.Sequelize(dbName, dbUsername, dbPassword, {
  host: dbHost,
  dialect: 'postgres',
  dialectOptions: {
    decimalNumbers: true,
  },
  port: dbPort,
  omitNull: true,
  define: {
    underscored: true,
    createdAt: 'createdAt',
    updatedAt: 'updatedAt',
  },
  logging: (a, b) => { console.log(a, b.bind) },
});

fs.readdirSync(modelDir)
  .filter((file) => file.indexOf('.') !== 0 && file !== 'index.js' && /\.js$/.test(file))
  .forEach((file) => {
    const model = require(path.join(modelDir, file))(sequelizeInstance, sequelize.DataTypes);

    model.prototype.toJSON = function toJSON() {
      const values = this.toObject();

      return values;
    };

    model.prototype.toObject = function toObject() {
      return this.get({ plain: true });
    };

    db[model.name] = model;
  });

Object.keys(db).forEach((modelName) => {
  if ('associate' in db[modelName]) {
    db[modelName].associate(db);
  }
});

db.sequelize = sequelizeInstance;
db.Sequelize = sequelize.Sequelize;

module.exports = db;
