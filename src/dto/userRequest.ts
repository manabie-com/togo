/* 
 * DTO - Data Transfer Object
 * Handles data passing from controller to useCase
 * Also Handles input validation
 */
import { Error } from '../error'
import { CONSTANTS } from '../util/constant';

export interface UserRequest {
	id: number;
	username: string;
	password: string;
	type: string;
	validate(): void;
}

export const userRequest: UserRequest = {
	id: 0,
	username: '',
	password: '',
	type: CONSTANTS.ACCOUNT_TYPE.BASIC,
	validate() {
		if (!this.username || this.username == '') Error.exec(`Username is invalid: ${this.username}`, 400)
		if (!this.password || this.password == '') Error.exec('Password is invalid', 400)
		if (!this.type || this.type == '') Error.exec('Type is invalid it should be "basic" or "premium"', 400)
	}
}
