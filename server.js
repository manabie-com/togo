/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

const http = require("http");
const express = require("express");
const cors = require("cors");
const bodyParser = require("body-parser");
const cookieParser = require("cookie-parser");

const router = require("./router");

const app = express();
const server = http.createServer(app);

app.use(cors());
app.use(express.json());
app.use(bodyParser.urlencoded({ extended: true }));
app.use(cookieParser());

app.use((err, req, res, _next) => {
  res.json({
    error_message: "Body should be a JSON",
  });
});

app.use(
  express.urlencoded({
    extended: true,
  })
);

app.get("/", function (req, res) {
  res.send("Hello world! This is Togo backend server!");
});

router(app);

app.use((err, req, res, next) => {
  res.status(200).json({
    error_code: err.error_code || err.message,
    error_message: err.message,
    error_data: err.error_data,
  });
});

module.exports = server;
