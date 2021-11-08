import dotenv from 'dotenv'
import path from 'path'
import fs from 'fs'

dotenv.config(path.resolve(process.cwd(), '.env'))
let folder = ""
switch (process.env.NODE_ENV) {
    case "staging":
        folder = "staging"
        break
    case "production":
        folder = "product"
        break
}
let fullPathConfig = "./"
if (folder !== "") {
    fullPathConfig =  "app/Configs/" + folder + "/"
}
//neu co khai bao them file config khác thì đặt tên no vào mảng dưới
let arr_config_file = ["auth", "constant", "database", "queue", "mail", "log", "session", "cdn"]
let configs = {}
let obj = {}
let path_temp
for (let x in arr_config_file) {
    path_temp = fullPathConfig + arr_config_file[x] + ".js"
    if (fs.existsSync(path_temp)) {
        obj = require("./" + folder + "/" + arr_config_file[x] + ".js")
    } else {
        obj = require("./" + arr_config_file[x] + ".js")
    }
    configs = {...configs, ...obj.default}

}
export default configs

