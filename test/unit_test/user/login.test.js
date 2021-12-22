const UserModel = require('../../../src/models/user')
const {loginService} =require('../../../src/services/user')

describe('User: Register', () => {
    it("Check login Successful", async()=>{
        const data={
            name: "Nhan Phan",
            email: "hieunhan2000@gmail.com",
            password: "$2a$10$EvqgG3/AEdIadKpxG7u6oeoSZy.fJw3S55xdV9Q5opAuvuY9GowYu",
            tokens: [
                {
                    token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MWJmNGE0ZjZmODVmY2U5ZmVjYWZlZmYiLCJpYXQiOjE2Mzk5MjYzNTJ9.nQBn83TbxQe-MALLy82Sw4s_o5TVHZLY_QKnEk0yNgk"
                }
            ],
        }
        const userModel = new UserModel(data)
        UserModel.findOne = jest.fn().mockResolvedValue(userModel)
        UserModel.findOneAndUpdate = jest.fn().mockResolvedValue(data)
        const newUser = await loginService(data.email, "1234567")
        expect(newUser.user).toStrictEqual(userModel)
    })
    it("Check login username not exist", async () => {
        try {
            UserModel.findOne = jest.fn().mockResolvedValue(null)
            await loginService("hieunhan123@gmail.com", "1234567")
        } catch (err) {
            expect(err.message).toStrictEqual('User is not exist')
        }
      });
      it("Check login wrong password", async()=>{
          try{
            const data={
                name: "Nhan Phan",
                email: "hieunhan2000@gmail.com",
                password: "$2a$10$EvqgG3/AEdIadKpxG7u6oeoSZy.fJw3S55xdV9Q5opAuvuY9GowYu",
                tokens: [
                    {
                        token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2MWJmNGE0ZjZmODVmY2U5ZmVjYWZlZmYiLCJpYXQiOjE2Mzk5MjYzNTJ9.nQBn83TbxQe-MALLy82Sw4s_o5TVHZLY_QKnEk0yNgk"
                    }
                ],
            }
            const userModel = new UserModel(data)
            UserModel.findOne = jest.fn().mockResolvedValue(userModel)
            const newUser = await loginService(data.email, "12345673242")
        }
        catch (err) {
            expect(err.message).toStrictEqual('Password is wrong');
        }
    })
});