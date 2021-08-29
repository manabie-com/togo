import * as chai from 'chai';
import 'mocha';
import { SERVICES, UserService } from "../../service";
import { AppContainer } from "../../core";
import chaiHttp = require('chai-http');

chai.use(chaiHttp);
const should = chai.should();
const container = AppContainer.Load();

describe('[UNIT] Test Service', () => {

    describe('UserService', () => {
        const userService = container.get<UserService>(SERVICES.UserService);
        it('it should return a token', (done) => {
            userService.createToken('firstUser').then(res => {
                res.should.be.a('string');
                done();
            }).catch(done);
        });
    });
});
