import { Error } from '../error'

/* 
 * DTO - Data Transfer Object
 * Handles data passing from controller to useCase
 * Also Handles input validation
 */
export interface AuthRequest {
	username: string;
	password: string;
	validate(): void;
}

export const authRequest: AuthRequest = {
	username: '',
	password: '',
	validate() {
		if (!this.username || this.username == '') Error.exec(`Username is invalid: ${this.username}`, 400)
		if (!this.password || this.password == '') Error.exec('Password is invalid', 400)
	}
}
