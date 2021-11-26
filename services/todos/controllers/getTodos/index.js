const fs = require("fs");

const getTodos = (modelUrl) => {
  try {
    const data = fs.readFileSync(modelUrl, "utf8");
    const todos = JSON.parse(data);
    return todos;
  } catch (err) {
    return err;
  }
};

module.exports = {
  getTodos,
};
