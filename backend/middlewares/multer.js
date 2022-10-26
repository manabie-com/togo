const multer = require('multer');
const path = require('path');
const filePath = path.join(__dirname, '../../excel');
const logger = require('../config/logger');
const httpStatus = require('http-status');
const ApiError = require('../utils/ApiError');
const excelFilter = (req, file, cb) => {
  if (
    // file.mimetype.includes("excel")
    file.mimetype === 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
  ) {
    cb(null, true);
  } else {
    cb(new ApiError(httpStatus.BAD_REQUEST, 'Unsupported file format'), false);
  }
};
const storage = multer.diskStorage({
  destination: (req, file, cb) => {
    // logger.info('filePath: ', filePath);
    // cb(null, __basedir + "/resources/static/assets/uploads/");
    cb(null, `${filePath}`);
  },
  filename: (req, file, cb) => {
    // logger.info('filename', file.originalname);
    cb(null, `data_upload.xlsx`);
  },
});
const upload = multer({ storage: storage, fileFilter: excelFilter });
module.exports = upload;
