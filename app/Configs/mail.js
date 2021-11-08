import dotenv from 'dotenv'
import path from 'path'

dotenv.config(path.resolve(process.cwd(), '.env'))
const config = {
    mail: {
        key: process.env.SENDGRID_API_KEY || "",
        from_tatmart: "sales@tatmart.com"
    }
}
export default config