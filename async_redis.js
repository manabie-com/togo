'use strict';
const redis = require("redis");


const retry_strategy = function(options) {
  if (options.error && (options.error.code === 'ECONNREFUSED' || options.error.code === 'NR_CLOSED')) {
      // Try reconnecting after 5 seconds
      console.error('The server refused the connection. Retrying connection...');
      return 1000;
  }
  if (options.total_retry_time > 1000 * 60 * 60) {
      // End reconnecting after a specific timeout and flush all commands with an individual error
      return new Error('Retry time exhausted');
  }
  if (options.attempt >3) {
      // End reconnecting with built in error
      return new Error('Retry attempt exhausted');;
  }
  // reconnect after
  return Math.min(options.attempt * 100, 3000);
}


const asyncCreateClient = (port=6379, host='127.0.0.1') => {
  return new Promise((resolve, reject) => {
    const client = redis.createClient(port, host, {retry_strategy: retry_strategy});

    client.on('error', (err)=> {
      console.log('error')
      console.log(err);
      reject();
    });

    client.on('end', ()=>{
      // console.log('connection end')
      reject();
    })

    client.on('connect', (conn)=> {
        // console.log('connection success')
        resolve(client);
      });

  })
}


const asyncSet = (client, key, value) => {
  return new Promise((resolve, reject) => {
    client.set(key, value, (err, reply) => {
      if(err) reject(err);
      resolve(reply);
    })
  })
}


const asyncGet = (client, key) => {
  return new Promise((resolve, reject) => {
    client.get(key, (err, reply) => {
      if(err) reject(err);
      resolve(reply);
    })
  })
}


const asyncHSet = (client, hash, ...args) => {
  return new Promise((resolve, reject) => {
    client.hset(hash, args, (err, reply) => {
      if(err) reject(err);
      resolve(reply);
    })
  })
}


const asyncHGetAll = (client, hash) => {
  return new Promise((resolve, reject) => {
    client.hgetall(hash, (err, reply) => {
      if(err) reject(err);
      resolve(reply);
    })
  })
}


const asyncQuit = (client) => {
  return new Promise((resolve, reject) => {
    client.quit((err, reply) => {
      if(err) reject(err);
      resolve(reply);
    })
  })
}


const asyncFlush = (client) => {
  return new Promise((resolve, reject) => {
    client.flushdb((err, reply) => {
      if(err) reject(err);
      console.log('deleted');
      resolve(reply);
    })
  })
}


const asyncExpire = (client, key, time) => {
  return new Promise((resolve, reject) => {
    client.expire(key, time, (err, reply)=> {
      if(err) reject(err);
      console.log(key + ' expired after ' + time + ' seconds.');
      resolve(reply);
    })
  })
}


const asyncHDel = (client, hash, key) => {
  return new Promise((resolve, reject) => {
    client.hdel(hash, key, (err, reply) => {
      if(err) reject(err);
      resolve(reply);
    })
  })
}

const asyncKeys = (client, hash) => {
  return new Promise((resolve, reject) => {
    client.keys(hash, (err, reply) => {
      if(err) reject(err);
      resolve(reply);
    })
  })
}

module.exports = {
	asyncCreateClient : asyncCreateClient,
	asyncSet: asyncSet,
	asyncGet: asyncGet,
  asyncHSet: asyncHSet,
  asyncHGetAll: asyncHGetAll,
  asyncQuit : asyncQuit,
  asyncFlush: asyncFlush,
  asyncExpire: asyncExpire,
  asyncHDel: asyncHDel,
  asyncKeys: asyncKeys
}