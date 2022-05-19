'use strict'

const Client = require('pg').Client

module.exports.up = async function (next) {
  const client = new Client()
  await client.connect()
  try {
    await client.query('BEGIN')
    for (let i = 0; i < 10000; i++) {
      await client.query(`
        INSERT INTO "user" (
          id,
          name,
          task_limit
        ) VALUES (
          $1,
          $2,
          $3
        )
      `, [i, `user_${i + 1}`, i + 1])
    }
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
    let response = await client.query(`
      SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES 
        WHERE TABLE_NAME = 'user'
    `)
    let count = response.rows[0].count 
    if (count > 0) {
      for (let i = 0; i < 10000; i++) {
        await client.query(`
          DELETE FROM "user" where id=$1
        `, [i])
      }
    }
    await client.query('COMMIT')
  } catch (exception) {
    await client.query('ROLLBACK')
    throw exception
  } finally {
    client.end()
  }
}
