import dotenv from 'dotenv'
import path from 'path'

dotenv.config(path.resolve(process.cwd(), '.env'))

const config = {
    rabbitMq: {
        host: process.env.RABBIT_MQ_HOST || 'localhost',
        port: process.env.RABBIT_MQ_PORT || 15672,
        user: process.env.RABBIT_MQ_USERNAME || '',
        password: process.env.RABBIT_MQ_PASSWORD || ''
    }
}
export default config