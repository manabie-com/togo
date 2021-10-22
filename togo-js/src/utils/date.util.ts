import * as moment from 'moment';

// This should be a value object such as Date
// this came to me as i already implemented the start/end time function

export function getStartTime(dateString: string) {
  return moment(dateString)
    .set({ hours: 0, minutes: 0, seconds: 0 })
    .toISOString();
}
export function getEndTime(dateString: string) {
  return moment(dateString).set({ hours: 23, minutes: 59, seconds: 59 })
    .toISOString();
}
