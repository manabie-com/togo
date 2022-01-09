'use strict!';

const mongoose = require('mongoose');
const Event = require('events');


class DBConn extends Event {
  constructor() {
    super();
    this._mongoURI = process.env.MONGO_URI || 'mongodb://localhost:27017/togo';
    this._conn = false;
    this._retryTime = 1000;
  }

  async init() {
    // Already initialized
    if (this._initialized === true) return;
    this._initialized = true;
    await this.connect();
  }

  async connect() {

    while (!this._conn) {
      console.log(`[INFO]: ${new Date()} Wait for database connection...`);

      await mongoose.connect(this._mongoURI).then(res => {
        this._conn = true;
        console.log(`Connect DB Success!`);
      }).catch(err => {
        console.log(err.message);
      });

      if (!this._conn) {
        await new Promise((solve) => {
          setTimeout(() => {
            console.log(`${new Date()}: Init db connection fail. Retry after ${this._retryTime} ms.`);
            solve(true);
          }, 1000);
        })
      }
    }
    // TODO: retry connect db on loss connection
  }

  static get Instance() {
    return this._instance || (this._instance = new this());
  }
}

module.exports = DBConn.Instance;