import { Error } from '../error'

export interface AuthRequest {
	username: string;
	password: string;
	validate(): void;
}

const authRequest: AuthRequest = {
	username: '',
	password: '',
	validate() {
		if (!this.username || this.username == '') Error.exec(`Username is invalid: ${this.username}`, 400)
		if (!this.password || this.password == '') Error.exec('Password is invalid', 400)
	}
}

export default authRequest;
