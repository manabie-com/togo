const { app } = require('./app');
const dbConn = require('./utils/db');

async function start() {
    // First wait for DB connection ready 
    await dbConn.init();
    const port = process.env.PORT || 9100;
    const httpServer = app.listen(port, () => {
        console.log(`\n${new Date()}. Server listening on port ${port}\n`);
    });
    return httpServer;
}

start();