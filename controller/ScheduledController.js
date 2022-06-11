const TaskClass = require('../classes/TaskClass');
const cron = require('node-cron');

module.exports =
{
    async scheduled(){
    const job = cron.schedule('0 8 * * *', async () => {
        console.log('Reset Limit Every 8AM');
        let task_class = new TaskClass();
	    await task_class.resetDailyLimit();
      });
    job.start();
    }
}



