'use strict';
const redis = require("redis");
const { asyncKeys, asyncCreateClient, asyncSet, asyncGet, asyncHSet, asyncHGetAll, asyncQuit, asyncExpire } =  require("./async_redis.js");


const createUser = async (event) => {
	let body = validation(event.body)

	if(!body)
		return {statusCode: 400,body: JSON.stringify({message: 'Invalid payload' },null,2 ),};

	if(!event.headers || !event.headers.Authorization || !event.headers.Authorization===process.env.AUTH)
		return {statusCode: 400,body: JSON.stringify({message: 'Missing or invalid AUTHORIZATION header'},null,2 ),};
	
	try{
		var redisClient = await asyncCreateClient(process.env.PORT, process.env.REDIS_HOST);
		var id = (Date.now() + Math.floor(Math.random() * 10)).toString();
		await asyncHSet(redisClient, id, "limit", body.limit, "remain", body.limit);
		await asyncQuit(redisClient);
	}catch(e){
		console.log(e);
		return {statusCode: 500,body: JSON.stringify({message: 'Internal error' },null,2 ),};
	}

	return {statusCode: 200,body: JSON.stringify({message: 'SUCCESS',id: id, limit: body.limit },null,2 ),};

};


const getUser = async (event) => {
	if (event.body)
		var id = JSON.parse(event.body).id;
	else
		var id = ''
		
	if(!event.headers || !event.headers.Authorization || !event.headers.Authorization===process.env.AUTH)
		return {statusCode: 400,body: JSON.stringify({message: 'Missing or invalid AUTHORIZATION header'},null,2 ),};

	try{
		let redisClient = await asyncCreateClient(process.env.PORT, process.env.REDIS_HOST);
		let user = await asyncHGetAll(redisClient, id);

		let taskId = id + "_tasks"
		let task = await asyncHGetAll(redisClient, taskId) || {}
		
		await asyncQuit(redisClient);
		return {statusCode: 200,body: JSON.stringify({user: user, task: task},null,2 ),};
	}catch(e){
		console.log(e);
		return {statusCode: 500,body: JSON.stringify({message: 'Internal error' },null,2 ),};
	}
}


const resetLimit = async (event) => {
	if(!event.headers || !event.headers.Authorization || !event.headers.Authorization===process.env.AUTH)
		return {statusCode: 400,body: JSON.stringify({message: 'Missing or invalid AUTHORIZATION header'},null,2 ),};

	try{
		let redisClient = await asyncCreateClient(process.env.PORT, process.env.REDIS_HOST);
		let users = await asyncKeys(redisClient, "*");

		for(let i = 0; i < users.length; i++){
			let user = await asyncHGetAll(redisClient, users[i]);
			if (typeof user.remain !== 'undefined' && user.limit !== 'undefined')
				await asyncHSet(redisClient, users[i], "remain", user.limit)
		}
		
		await asyncQuit(redisClient);
		return {statusCode: 200,body: JSON.stringify({message: "Reset limit successful"},null,2 ),};
	}catch(e){
		console.log(e);
		return {statusCode: 500,body: JSON.stringify({message: 'Internal error' },null,2 ),};
	}
}


const validation = function (body) {
	if(!body)
		return
	
	let result = JSON.parse(body)

	if(!result.limit)
		return

	return result
}

module.exports = {
	createUser: createUser,
	getUser: getUser,
	resetLimit: resetLimit
}