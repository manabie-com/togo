const fs = require('fs');

const libs = {};
fs.readdirSync(`${process.cwd()}/libs`).forEach((fileName) => {
  fileName = fileName.slice(0, -3);
  libs[fileName] = require(`./${fileName}`);
});

module.exports = libs;
