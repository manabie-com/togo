import {Sequelize} from 'sequelize'
import configs from '../Configs'

export default new Sequelize(configs.mysql_master.db_master_db, null, null, {
    dialect: configs.mysql_master.db_master_dialect,
    logging: false,
    replication: {
        read: [
            {
                host: configs.mysql_slave.db_slave_host,
                port: configs.mysql_slave.db_slave_port,
                username: configs.mysql_slave.db_slave_user,
                password: configs.mysql_slave.db_slave_pass
            }
        ],
        write: {
            host: configs.mysql_master.db_master_host,
            port: configs.mysql_master.db_master_port,
            username: configs.mysql_master.db_master_user,
            password: configs.mysql_master.db_master_pass
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
