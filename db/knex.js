// Create a database connection and a database
const fs = require('fs');
const sqlite3 = require('sqlite3').verbose();
const dbPath = process.env.NODE_ENV === 'test' ? './db/test.db' : './db/data.db';
const db = new sqlite3.Database(dbPath);

const tasksSQL = fs.readFileSync('./db/models/tasks.sql').toString();
const usersSQL = fs.readFileSync('./db/models/users.sql').toString();
const taskCountSQL = fs.readFileSync('./db/models/user_tasks.sql').toString();
db.serialize(() => {
    db.run(tasksSQL, (err) => {
        if (err) throw err;
        console.log('Table tasks created successfully!');
    });

    db.run(usersSQL, (err) => {
        if (err) throw err;
        console.log('Table users created successfully!');
    });

    db.run(taskCountSQL, (err) => {
        if (err) throw err;
        console.log('Table user_tasks created successfully!');
    });
});


const options = {
    client: 'sqlite3',
    connection: {
        filename: dbPath,
    },
    useNullAsDefault: true,
};
const knex = require('knex')(options);

module.exports = knex;