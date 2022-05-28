const common = require('./ormconfig.common');

module.exports = Object.assign(common, {
  seeds: ['dist/database/seeds/*.seed{.ts,.js}'],
  entities: ['dist/modules/**/*.entity{.ts,.js}', 'dist/modules/**/entity/*.entity{.ts,.js}'],
  migrations: ['dist/database/migrations/mssql/*{.ts,.js}'],
  logging: process.env.ENABLED_RDS_QUERY_LOG === 'true',
});
