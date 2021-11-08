import md5 from 'md5'
import dotenv from 'dotenv'
import path from 'path'

dotenv.config(path.resolve(process.cwd(), '.env'))

const config = {
    tokenKey: md5(process.env.JSON_TOKEN ?? "token_test"),
    tokenTimeout: 30 * 24 * 60 * 60,//30 ngày
}

export default config


