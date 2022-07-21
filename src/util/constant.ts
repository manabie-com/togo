/*
 * HELPER function and CONSTANT data
 * to drive business logic
 */
export const CONSTANTS = {
	ACCOUNT_TYPE: {
		BASIC: 'BASIC',
		PREMIUM: 'PREMIUM',
		BASICLIMIT: 5,
		PREMIUMLIMIT: 10,
		GETDAILYLIMIT: function(type: string) {
			if (this.BASIC === type) return this.BASICLIMIT;
			if (this.PREMIUM === type) return this.PREMIUMLIMIT;
			return 0;
		}
	},
}
