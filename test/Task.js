const assert     = require("chai").assert;
const TaskClass   	 = require('../classes/TaskClass');

// describe('Add New Task', async() => {

//     it('It should add new task', async() => {	
// 		let taskInfo = {
// 			userName        :  'pergent100',
//             task            :  'swimming',
// 		} 
// 		let task = new TaskClass(taskInfo);
// 		const res = await task.createTask();
// 		assert.isObject(res);
//         assert.isNotNull(res.data.userName);
//         assert.isNotNull(res.data.task);
//         assert.isString(res.data.userName);
// 		assert.isString(res.data.task);
        
// 	});

// });

describe('Reset Limit Every 8am', async() => {

    it('It should reset all the limit of users', async() => {	

		let task = new TaskClass();
		const res = await task.resetDailyLimit();
		assert(res.status !== 'error', 'Successfuly update limit')

        
	});	

});