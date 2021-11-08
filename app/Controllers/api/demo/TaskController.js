import Controller from '../Controller'
import TaskRepositories from "../../../Repositories/TaskRepositories";

class TaskController extends Controller {

    constructor() {
        super()
    }

    /**
     * @api {post} /tasks Create task
     * @apiName Create task
     * @apiGroup Task
     * @apiHeaderExample {json} Header-Example:
     *     {
     *       "Authentication": "Bearer (token)",
     *       "Content-type" : "application\json"
     *     }
     * @apiParam {String} content Content of task. (required)
     *
     * @apiSuccess {Object} result
     * @apiSuccessExample Success-Response:
     *     {
     *       "message": "Thực thi task thành công!",
     *       "status": 200
     *     }
     *
     * @apiError {Object} error.
     * @apiErrorExample Error-Response:
     *     {
     *       "message": "...",
     *       "status": !=(200)
     *     }
     */
    tasks = async (req, res) => {
        let task = await TaskRepositories.updateTask({
            user_id: req.userData.user_id,
            content: req.body.content,
            created_date: new Date()
        })
        if(!task){
            return res.json({status: 500, message: "Server Error"})
        }
        //todo .....
        ////////////
        res.json({status: 200, message: "Thực thi task thành công!"})
    }

}

export const objClass = new TaskController();
