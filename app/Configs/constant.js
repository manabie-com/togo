import dotenv from 'dotenv'
import path from 'path'

dotenv.config(path.resolve(process.cwd(), '.env'))

const config = {
    homeRedis: "homeRedisPost",
    domain: process.env.DOMAIN_NAME || 'localhost',
}

export default config
