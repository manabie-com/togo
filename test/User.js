const assert     = require("chai").assert;
const UserClass   	 = require('../classes/UserClass');

describe('Add New User', async() => {

    it('It should validate user input', async() => {	
		let userinfo = {
			fullname        :  'Celine Suba',
            password        :  'Password12345',
            confirmpassword :  'Password12345',
            username        :  'Celine100'
		} 
        //userinfo=null;
		let user = new UserClass(userinfo);
		const res = await user.validate();
		assert.isObject(res);
        assert.isNotNull(res.data.username);
        assert.isNotNull(res.data.fullname);
        assert.isNotNull(res.data.password);
        assert.isNotNull(res.data.confirmpassword);
        assert.isString(res.data.username);
		assert.isString(res.data.fullname);
		assert.isString(res.data.password);
		assert.isString(res.data.confirmpassword);
        
	});	

	it('It should add user', async() => {	
		let userinfo = {
			fullname        :  'Celine Suba',
            password        :  'Password12345',
            username        :  'celine100'
		}
		let user = new UserClass(userinfo);
		const res = await user.create();
		assert.isObject(res);
        assert.isNotNull(res.data.fullname);
        assert.isNotNull(res.data.password);
        assert.isNotNull(res.data.username);
        assert.isString(res.data.username);
		assert.isString(res.data.fullname);
		assert.isString(res.data.password);
	});	   
});

describe('Login User', async() => {
	it('It should verify account to login', async() => {	

        let userinfo = {
            username        :  'celine100',
            password        :  'Password12345'

		}
		let user = new UserClass(userinfo);
		const res = await user.authenticate();
		assert.isObject(res, 'verify account');
        assert.isNotNull(res.data.username);
        assert.isNotNull(res.data.password);
	});
});