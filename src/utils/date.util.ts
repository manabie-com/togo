export class DateUtil {
    static getCurrentDate() {
        return new Date(Date.now()).toISOString().split('T')[0];
    }
}