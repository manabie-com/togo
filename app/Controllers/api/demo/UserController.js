import Controller from '../Controller'
import UserRepositories from "../../../Repositories/UserRepositories";
import jwt from "jsonwebtoken";
import configs from "../../../Configs";
import logger from "../../../Helpers/logger";

class UserController extends Controller {

    constructor() {
        super()
    }

    /**
     * @api {post} /login Login user
     * @apiName Login
     * @apiGroup User
     * @apiHeaderExample {json} Header-Example:
     *     {
     *       "Content-type" : "application\json"
     *     }

     * @apiParam {string} username account login (required)
     * @apiParam {string} password your password  (required)
     *
     * @apiSuccess {Object} result
     * @apiSuccessExample Success-Response:
     *     {
     *        "status": 200
     *        "message": "Thành công!",
     *        "data": {
     *            "token": "..."
     *        }
     *     }
     *
     * @apiError {Object} error.
     * @apiErrorExample Error-Response:
     *     {
     *       "message": "...",
     *       "status": !=(200)
     *     }
     */
    login = async (req, res) => {
        let params = {...req.body};
        if (!params.username || !params.password) {
            return res.json({
                status: 422,
                message: "Biến truyền vào không đúng"
            });
        }
        let userData = await UserRepositories.getUser({
            username: params.username,
            password: params.password
        })
        if (userData != null && userData) {
            jwt.sign({...userData.dataValues}, configs.tokenKey, {expiresIn: configs.tokenTimeout}, function (err, token) {
                if (err) {
                    logger.error(err.message)
                    res.json({status: 500, message: "Server Error", data: {token: token}})
                } else {
                    res.json({status: 200, message: "Thành công!", data: {token: token}})
                }
            })
        } else {
            return res.json({
                status: 400,
                message: "Tài khoản không tồn tại"
            });
        }
    }
}

export const objClass = new UserController();
