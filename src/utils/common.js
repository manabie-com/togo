/**
 * Get start of current day
 * @returns {string}
 */
const getStartOfDay = () => {
  const currentDate = new Date();
  return currentDate.setHours(0, 0, 0, 0);
};

/**
 * Generate object id from date
 * @param {number} date
 * @returns {string}
 */
const objectIdFromDate = (date) => {
  return Math.floor(date / 1000).toString(16) + "0000000000000000";
};

module.exports = {
  getStartOfDay,
  objectIdFromDate,
};
