import configs from '../Configs'
import jwt from 'jsonwebtoken'
import dotenv from 'dotenv'
import UserRepositories from "../Repositories/UserRepositories";
import TaskRepositories from "../Repositories/TaskRepositories";
import {convertCurrentDate} from "./function";

dotenv.config()

export function validToken(req, res, next) {
    let token = req.get('authorization')

    if (!token || token.length < 64) {
        res.json({
            status: 403,
            message: "authorization không chính xác"
        })
    } else {
        token = token.split(' ').pop()
        jwt.verify(token, configs.tokenKey, async function (err, decode) {
            if (err) {
                return res.json({status: 403, message: "authorization không chính xác"})
            }
            //valid user from token
            let userData = await UserRepositories.getUser({
                userId: decode.user_id
            })
            if (!userData) {
                return res.json({status: 422, message: "Token không chính xác"})
            }
            if (decode.username === userData.username) {
                let date = convertCurrentDate()
                //valid max_todo
                let totalTask = await TaskRepositories.getTasks({
                    onlyCount: true,
                    userId: [userData.user_id],
                    dateFrom: `${date} 00:00:00`,
                    dateTo: `${date} 23:59:59`
                })

                if(totalTask === null){
                    return res.json({status: 500, message: "Server Error"})
                }

                if (userData.max_todo <= totalTask) {
                    return res.json({status: 400, message: "Đã quá số lượng yêu cầu api"})
                }

                req.userData = userData
                next()
            }
        })
    }
}

export function defaultRequest(app, Controllers, apiPath) {
    for (let k in Controllers) {
        let slug = apiPath + "/" + k.toLowerCase()
        if (typeof Controllers[k].index === "function") {
            app.get(slug, (req, res) => Controllers[k].index(req, res))
        }
        if (typeof Controllers[k].create === "function") {
            app.post(slug, [validToken], (req, res) => Controllers[k].create(req, res))
        }
        if (typeof Controllers[k].update === "function") {
            app.post(slug + "/:id?", [validToken], (req, res) => Controllers[k].update(req, res))
            app.put(slug + "/:id?", [validToken], (req, res) => Controllers[k].update(req, res))
        }
        if (typeof Controllers[k].delete === "function") {
            app.post(slug + "/:id?", [validToken], (req, res) => Controllers[k].delete(req, res))
            app.delete(slug + "/:id?", [validToken], (req, res) => Controllers[k].delete(req, res))
        }
    }
}
