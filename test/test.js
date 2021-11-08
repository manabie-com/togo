import chai from "chai"
import chaiHttp from "chai-http"
import server from "../app/app"
import {apiPath} from "../app/Routes/api"
chai.use(chaiHttp)
let should = chai.should()

describe('Manabie Mini testing', () => {
    beforeEach((done) => {
        //Before each test we empty the database in your case
        done();
    });

    let token = '';

    it('Login account get Token', (done) => {
        let account = {
            username: "test",
            password: "123456"
        };
        chai.request(server)
            .post(apiPath + '/login')
            .send(account)
            .end((err, res) => {
                res.should.have.status(200);
                res.body.should.be.a('object');
                res.body.should.have.property('message').eql('Thành công!');
                res.body.should.have.property('status').eql(200);
                token = res.body?.data?.token
                done();
            });
    });

    it('Call Task', (done) => {
        let data = {
            content: "test"
        };
        chai.request(server)
            .post(apiPath + '/tasks')
            .set('authorization', 'Bearer '+ token)
            .send(data)
            .end((err, res) => {
                if(res.body?.status === 400){
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('message').eql('Đã quá số lượng yêu cầu api');
                    res.body.should.have.property('status').eql(400);
                }else{
                    res.should.have.status(200);
                    res.body.should.be.a('object');
                    res.body.should.have.property('message').eql('Thực thi task thành công!');
                    res.body.should.have.property('status').eql(200);
                }
                done();
            });
    });
})