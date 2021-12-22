const UserModel = require('../../../src/models/user')
const {createUserService} =require('../../../src/services/user')

describe('User: Register', () => {
    it("Check create user Successful", async()=>{
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
        UserModel.create = jest.fn().mockResolvedValue(userModel)
        UserModel.findOne = jest.fn().mockResolvedValue(null);
        UserModel.findOneAndUpdate = jest.fn().mockResolvedValue(userModel)
        const newUser = await createUserService({name: data.name, email: data.email, password: "123456"})
        expect(newUser.user).toStrictEqual(userModel)
    })
    it("Check create fail user exist", async () => {
        try {
          const name = "Nhan Phan";
          const password = "1234567";
          const email="hieunhan2000@gmail.com";
          UserModel.findOne = jest.fn().mockResolvedValue({
              name, password, email
          });
           await createUserService({name: name, email: email, password: password})
        } catch (err) {
          expect(err.message).toStrictEqual('Email is used');
        }
      });
});