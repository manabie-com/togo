import postgres from "./postgres"
import mysql from "./mysql"
import dotenv from 'dotenv'
import path from "path";

dotenv.config(path.resolve(process.cwd(), '.env'))
let database = postgres
if (process.env.DATABASE_USE === "mysql") {
    database = mysql
}

export default database