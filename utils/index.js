/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

module.exports = {
  /**
   *
   * @returns today with type string by format dd-mm-yyyy
   */
  getTodayString: () => {
    const today = new Date();
    const todayString = `${today.getDate()}-${
      today.getMonth() + 1
    }-${today.getFullYear()}`;

    return todayString;
  },
};
