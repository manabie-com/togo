
describe('Login test case', () => {
  it('verify prevent unauthorized request', () => {
    cy.request({
      method: 'GET',
      url: "http://localhost:3001/todos",
      body: {},
      failOnStatusCode: false,
      headers: {
        'Content-Type': 'application/json'
      }
    }).then((res)=> {
      expect(res.status).to.eq(401)
    })
  })
  it('verify authorized request', () => {
    cy.request({
      method: 'GET',
      url: "http://localhost:3001/todos",
      body: {},
      failOnStatusCode: false,
      headers: {
        'Content-Type': 'application/json',
        'Authorization' : 'testabc.xyz.ahk'
      }
    }).then((res)=> {
      expect(res.status).to.eq(200)
    })
  })
  it('verify login flow', () => {
    cy.request({
      method: 'GET',
      url: "http://localhost:3001/login",
      body: {},
      failOnStatusCode: false,
      headers: {
        'Content-Type': 'application/json'
      }
    }).then((res)=> {
      expect(res.status).to.eq(401)
    })
  })
  it('verify login flow with username pwd', () => {
    cy.request({
      method: 'GET',
      url: "http://localhost:3001/login",
      body: {
        username: "firstUser",
        password: "example"
      },
      failOnStatusCode: false,
      headers: {
        'Content-Type': 'application/json'
      }
    }).then((res)=> {
      expect(res.status).to.eq(200)
      expect(res.body.token).to.eq('testabc.xyz.ahk')
      expect(res.body.message).to.eq('Success')
    })
  })
})
