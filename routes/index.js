const fs = require('fs');

const routers = {};
fs.readdirSync(`${process.cwd()}/routes`).forEach((fileName) => {
  fileName = fileName.slice(0, -3);
  if (fileName === 'index') {
    return;
  }
  routers[fileName] = require(`./${fileName}`);
});

module.exports = routers;
