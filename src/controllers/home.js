'use strict!';


const { version } = require('../../package.json');

// GET: /api/public/home
module.exports.home = async (ctx) => {
  return ctx.showResult({
    message: `REST API VERSION ${version}.`,
    date: new Date()
  })
};
