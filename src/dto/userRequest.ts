import { Error } from '../error'

export interface UserRequest {
	username: string;
	password: string;
	type: string;
	validate(): void;
}

export const userRequest: UserRequest = {
	username: '',
	password: '',
	type: 'basic',
	validate() {
		if (!this.username || this.username == '') Error.exec(`Username is invalid: ${this.username}`, 400)
		if (!this.password || this.password == '') Error.exec('Password is invalid', 400)
		if (!this.type || this.type == '') Error.exec('Type is invalid it should be "basic" or "premium"', 400)
	}
}
