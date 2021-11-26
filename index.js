const express = require("express");
const app = express();
const server = require("http").Server(app);

const router = require("./routes");

app.use(router);

const port = process.env.PORT || 5050;
server.listen(port, () => {
  console.log(`Server has started with port ${port}`);
});
