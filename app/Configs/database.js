import dotenv from 'dotenv'
import path from 'path'

dotenv.config(path.resolve(process.cwd(), '.env'))

const config = {
    mysql_master: {
        db_master_dialect: process.env.DB_MASTER_DIALECT || 'mysql',
        db_master_host: process.env.DB_MASTER_HOST || 'localhost',
        db_master_port: process.env.DB_MASTER_PORT || 3306,
        db_master_user: process.env.DB_MASTER_USERNAME || 'root',
        db_master_pass: process.env.DB_MASTER_PASSWORD || '',
        db_master_db: process.env.DB_MASTER_DATABASE || '',
    },
    mysql_slave: {
        db_slave_host: process.env.DB_SLAVE_HOST || process.env.DB_MASTER_HOST,
        db_slave_port: process.env.DB_SLAVE_PORT || process.env.DB_MASTER_PORT,
        db_slave_user: process.env.DB_SLAVE_USERNAME || process.env.DB_MASTER_USERNAME,
        db_slave_pass: process.env.DB_SLAVE_PASSWORD || process.env.DB_MASTER_PASSWORD,
        db_slave_db: process.env.DB_SLAVE_DATABASE || process.env.DB_MASTER_DATABASE,
    },
    postgres: {
        db_postgres_dialect: process.env.DB_POSTGRES_DIALECT || 'postgres',
        db_postgres_host: process.env.DB_POSTGRES_HOST || "localhost",
        db_postgres_port: process.env.DB_POSTGRES_PORT || 5432,
        db_postgres_user: process.env.DB_POSTGRES_USERNAME || 'postgres',
        db_postgres_pass: process.env.DB_POSTGRES_PASSWORD || "",
        db_postgres_db: process.env.DB_POSTGRES_DATABASE || "",
    },
    mssql: {
        db_dialect: process.env.MSSQL_DIALECT || 'mssql',
        db_host: process.env.MSSQL_HOST || "localhost",
        db_port: process.env.MSSQL_PORT || '1433',
        db_user: process.env.MSSQL_USERNAME || '',
        db_pass: process.env.MSSQL_PASSWORD || '',
        db_name: process.env.MSSQL_DATABASE || '',
    },
    redis: {
        host: process.env.REDIS_HOST || 'localhost',
        port: process.env.REDIS_PORT || 6379,
        password: process.env.REDIS_PASS || '',
        db: process.env.REDIS_DB || null,
        prefix: '',
    },
    mongodb: {
        host: process.env.MONGO_MASTER_HOST || 'localhost',
        port: process.env.MONGO_MASTER_PORT || 27017,
        user: process.env.MONGO_MASTER_USERNAME || '',
        pass: process.env.MONGO_MASTER_PASSWORD || '',
        database: process.env.MONGO_MASTER_DATABASE || 'tat_gateway',
    },

    elasticSearch: {
        host: process.env.ELASTIC_SEARCH_HOST || 'localhost',
        port: process.env.ELASTIC_SEARCH_PORT || 27017,
        version: process.env.ELASTIC_SEARCH_VERSION || ''
    }

}

export default config
