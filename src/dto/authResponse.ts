/* 
 * DTO - Data Transfer Object
 * Handles data passing from controller to useCase
 * Also Handles input validation
 */
export interface AuthResponse {
	token: string;
}

export interface AuthUserResponse {
	id: number;
	username: string;
	password: string;
	type: string;
}

export const authResponse: AuthResponse = {
	token: '',
}

export const authUserResponse: AuthUserResponse = {
	id: 0,
	username: '',
	password: '',
	type: '',
}
