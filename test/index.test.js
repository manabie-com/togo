/**
 * @author Nguyen Minh Tam / ngmitamit@gmail.com
 */

//During the test the env variable is set to test
process.env.NODE_ENV = 'test';

//Require the dev-dependencies
let chai = require('chai');
let chaiHttp = require('chai-http');
let server = require('../server');
let should = chai.should();

chai.use(chaiHttp);


describe('Togo',()=>{
    describe("Main", ()=>{
        it('it should GET hello message', (done) => {
            chai.request(server)
                .get('/')
                .end((err, res) => {
                    res.should.have.status(200);
                    res.text.should.be.a('string');
                    res.text.should.be.eql("Hello world! This is Togo backend server!");
                    done();
                });
        });

    })
})