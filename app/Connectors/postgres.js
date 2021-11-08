import {Sequelize} from 'sequelize'
import configs from '../Configs'

export default new Sequelize(configs.postgres.db_postgres_db, null, null, {
    dialect: configs.postgres.db_postgres_dialect,
    logging: true,
    replication: {
        read: [
            {
                host: configs.postgres.db_postgres_host,
                port: configs.postgres.db_postgres_port,
                username: configs.postgres.db_postgres_user,
                password: configs.postgres.db_postgres_pass
            }
        ],
        write: {
            host: configs.postgres.db_postgres_host,
            port: configs.postgres.db_postgres_port,
            username: configs.postgres.db_postgres_user,
            password: configs.postgres.db_postgres_pass
        }
    },
    timezone: "+07:00",
    dialectOptions: {
        connectTimeout: 10000
    },
    pool: {
        max: 20,
        idle: 30000,
    }
})
