const config = require('config')

const dbConfig = config.get('dbConfig')

const Sequelize = require('sequelize')

let sequelize

const operatorsAliases = {
  $eq: Sequelize.Op.eq,
  $ne: Sequelize.Op.ne,
  $gte: Sequelize.Op.gte,
  $gt: Sequelize.Op.gt,
  $lte: Sequelize.Op.lte,
  $lt: Sequelize.Op.lt,
  $not: Sequelize.Op.not,
  $in: Sequelize.Op.in,
  $notIn: Sequelize.Op.notIn,
  $is: Sequelize.Op.is,
  $like: Sequelize.Op.like,
  $notLike: Sequelize.Op.notLike,
  $iLike: Sequelize.Op.iLike,
  $notILike: Sequelize.Op.notILike,
  $regexp: Sequelize.Op.regexp,
  $notRegexp: Sequelize.Op.notRegexp,
  $iRegexp: Sequelize.Op.iRegexp,
  $notIRegexp: Sequelize.Op.notIRegexp,
  $between: Sequelize.Op.between,
  $notBetween: Sequelize.Op.notBetween,
  $overlap: Sequelize.Op.overlap,
  $contains: Sequelize.Op.contains,
  $contained: Sequelize.Op.contained,
  $adjacent: Sequelize.Op.adjacent,
  $strictLeft: Sequelize.Op.strictLeft,
  $strictRight: Sequelize.Op.strictRight,
  $noExtendRight: Sequelize.Op.noExtendRight,
  $noExtendLeft: Sequelize.Op.noExtendLeft,
  $and: Sequelize.Op.and,
  $or: Sequelize.Op.or,
  $any: Sequelize.Op.any,
  $all: Sequelize.Op.all,
  $values: Sequelize.Op.values,
  $col: Sequelize.Op.col
}
const DATABASE_URL = process.env.HEROKU_POSTGRESQL_DBNAME_URL || process.env.DATABASE_URL
console.log('DATABASE_URL:' + DATABASE_URL)
if (DATABASE_URL) {
  sequelize = new Sequelize(DATABASE_URL, {
    dialect: 'postgres',
    protocol: 'postgres',
    dialectOptions: {
      ssl: true
    }
  })
} else {
  sequelize = new Sequelize(dbConfig.DB, dbConfig.USER, dbConfig.PASSWORD, {
    host: dbConfig.HOST,
    dialect: dbConfig.dialect,
    dialectOptions: {
      ssl: {
          require: true,
          rejectUnauthorized: false
      }
  },
    pool: {
      max: dbConfig.pool.max,
      min: dbConfig.pool.min,
      acquire: dbConfig.pool.acquire,
      idle: dbConfig.pool.idle
    },
    logging: true,
    operatorsAliases: operatorsAliases
  })
}

const db = {}

db.Sequelize = Sequelize
db.sequelize = sequelize

db.user = require('./user')(sequelize, Sequelize)
db.account = require('./account')(sequelize, Sequelize)
db.task = require('./task')(sequelize, Sequelize)

/* Mapping on relationship between each tables */
Object.keys(db).forEach((modelName) => {
  if (db[modelName].associate) {
    db[modelName].associate(db)
  }
})
module.exports = db
