'use strict'

const Client = require('pg').Client

module.exports.up = async function (next) {
  const client = new Client()
  await client.connect()
  try {
    await client.query('BEGIN')
    await client.query(`
      CREATE TABLE IF NOT EXISTS "user" (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        task_limit INT NOT NULL CHECK (task_limit > 0)
      )
    `)

    await client.query(`
      CREATE TABLE IF NOT EXISTS "tasks" (
        id SERIAL PRIMARY KEY,
        user_id INT REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE CASCADE NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL 
      )
    `)
    await client.query('COMMIT')
  } catch (exception) {
    await client.query('ROLLBACK')
    throw exception
  } finally {
    client.end()
  }
}

module.exports.down = async function (next) {
  const client = new Client()
  await client.connect()
  try {
    await client.query('BEGIN')
    await client.query(`
      CREATE TABLE IF NOT EXISTS "tasks"
    `)
    await client.query(`
      DROP TABLE IF EXISTS "user"`)
    await client.query('COMMIT')
  } catch (exception) {
    await client.query('ROLLBACK')
    throw exception
  } finally {
    client.end()
  }
}
