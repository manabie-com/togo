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

    const dd = `${today.getDate()}`.padStart(2, 0);
    const mm = `${today.getMonth() + 1}`.padStart(2, 0);
    const yyyy = today.getFullYear();

    const todayString = `${dd}-${mm}-${yyyy}`;

    return todayString;
  },
};
