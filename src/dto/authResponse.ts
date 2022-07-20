export interface AuthResponse {
	token: string;
}

export interface AuthUserResponse {
	id: number;
	username: string;
	password: string;
	type: string;
}

const authResponse: AuthResponse = {
	token: '',
}

export const authUserResponse: AuthUserResponse = {
	id: 0,
	username: '',
	password: '',
	type: '',
}

export default authResponse;
