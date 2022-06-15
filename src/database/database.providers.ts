import { Sequelize } from 'sequelize-typescript';
import { ENV } from 'src/constance/variable';
import { LimitTask } from 'src/model/limit-task/schema/limitTask.entity';
import { Task } from 'src/model/task/schema/task.entity';
import { User } from 'src/model/user/schema/user.entity';
import { databaseConfig } from './database.config';

export const databaseProviders = [
  {
    provide: 'SEQUELIZE',
    useFactory: async () => {
      const sequelize = sequelizeSingleTon.getInstance();
      await sequelize.sync();
      return sequelize;
    },
  },
];

const sequelizeSingleTon = (function () {
  let instance;
  process.env.NODE_ENV = '';
  function createInstance() {
    let config;
    switch (process.env.NODE_ENV) {
      case ENV.DEVELOPMENT:
        config = databaseConfig.development;
        break;
      case ENV.PRODUCTION:
        config = databaseConfig.production;
        break;
      default:
        config = databaseConfig.development;
    }
    const sequelize = new Sequelize(config);
    sequelize.addModels([Task, User, LimitTask]);
    Object.assign(sequelize.options, { logging: false, timestamps: false });
    return sequelize;
  }

  return {
    getInstance: function () {
      if (!instance) instance = createInstance();
      return instance;
    },
  };
})();
