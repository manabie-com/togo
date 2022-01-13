'use strict';
const redis = require("redis");
const { asyncCreateClient, asyncSet, asyncGet, asyncHSet, asyncHGetAll, asyncQuit, asyncExpire, asyncHDel } =  require("./async_redis.js");


const createTask = async (event) => {
	let body = validation(event.body, createTask)

	if(!body)
		return {statusCode: 400,body: JSON.stringify({message: 'Invalid payload' },null,2 ),};

	if(!event.headers || !event.headers.Authorization || !event.headers.Authorization===process.env.AUTH)
		return {statusCode: 400,body: JSON.stringify({message: 'Missing or invalid AUTHORIZATION header'},null,2 ),};
	
	try{
		let redisClient = await asyncCreateClient(process.env.PORT, process.env.REDIS_HOST);

        let user = await asyncHGetAll(redisClient, body.userId)
        if (!user)
            return {statusCode: 400,body: JSON.stringify({message: 'User not found' },null,2 ),};
		if (parseInt(user.remain) === 0)
			return {statusCode: 400,body: JSON.stringify({message: 'Create task limit reached' },null,2 ),};

		var taskKey = body.userId + "_tasks"
        var taskId = (Date.now() + Math.floor(Math.random() * 10)).toString() + "_task";
		await asyncHSet(redisClient, taskKey, taskId, body.description);

		let remain = parseInt(user.remain) - 1
		await asyncHSet(redisClient, body.userId, "remain", remain)
		
		await asyncQuit(redisClient);
	}catch(e){
		console.log(e);
		return {statusCode: 500,body: JSON.stringify({message: 'Internal error' },null,2 ),};
	}

	return {statusCode: 200,body: JSON.stringify({message: 'SUCCESS',taskId: taskId, description: body.description },null,2 ),};

};


const deleteTask = async (event) => {
	let body = validation(event.body, deleteTask)
    if(!body)
		return {statusCode: 400,body: JSON.stringify({message: 'Invalid payload' },null,2 ),};
		
	if(!event.headers || !event.headers.Authorization || !event.headers.Authorization===process.env.AUTH)
		return {statusCode: 400,body: JSON.stringify({message: 'Missing or invalid AUTHORIZATION header'},null,2 ),};

	try{
		let redisClient = await asyncCreateClient(process.env.PORT, process.env.REDIS_HOST);
		let user = await asyncHGetAll(redisClient, body.userId);
        if (!user)
            return {statusCode: 400,body: JSON.stringify({message: 'User not found' },null,2 ),};

        let key = body.userId + "_tasks";
		console.log("key: " + key);
		console.log("task: " + body.taskId);

        await asyncHDel(redisClient, key, body.taskId);
		await asyncQuit(redisClient);
		return {statusCode: 200,body: JSON.stringify({message: 'Success delete'},null,2 ),};
	}catch(e){
		console.log(e);
		return {statusCode: 500,body: JSON.stringify({message: 'Internal error' },null,2 ),};
	}
}


const validation = function (body, type) {
	if(!body)
		return
	
	let result = JSON.parse(body)

    if(type === "createTask") {
        if(!result.userId || !result.description )
		    return
    }else if (type === "deleteTask"){
        if(!result.userId || !result.taskId )
		    return
    }
	
	return result
}

module.exports = {
	createTask: createTask,
	deleteTask: deleteTask
}