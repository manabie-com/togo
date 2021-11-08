import dotenv from 'dotenv'
import path from "path"
import {defaultRequest, validToken} from '../Helpers/middlewares'

dotenv.config(path.resolve(process.cwd(), '.env'))
export const apiPath = "/api/v1"

export default (app, Controllers) => {
    //set default methods for all API restful
    defaultRequest(app, Controllers, apiPath)

    /* controller API*/
    //-- user Controller
    app.post(apiPath + '/login', (req, res) => Controllers["User"].login(req, res))
    //--tasks Controller
    app.post(apiPath + '/tasks', [validToken], (req, res) => Controllers["Task"].tasks(req, res))


}
// sequelize-auto -o "./app/Models/Mysql" -d manabie -h localhost -u root -p 3306 -x  -e mariadb --lang esm --caseModel p -t tasks
