const common = require('./ormconfig.common');

module.exports = Object.assign(common, {
  seeds: ['src/database/seeds/*.seed.ts'],
  entities: ['src/modules/**/*.entity{.ts,.js}', 'src/modules/**/entity/*.entity{.ts,.js}'],
  migrations: ['src/database/migrations/mssql/*{.ts,.js}'],
  logging: false,
});
