const fs = require('fs');

const constants = {};
fs.readdirSync(`${process.cwd()}/constants`).forEach((fileName) => {
  fileName = fileName.slice(0, -3);
  constants[fileName] = require(`./${fileName}`);
});

module.exports = constants;
