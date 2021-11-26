Cypress.on('uncaught:exception', () => false)

const env = Cypress.env()
const { baseUrlConsole: baseUrl, user } = env
describe('Login', function () {
  it('should login successful when userId & password right', function () {
    cy.visit(`${baseUrl}`)
    cy.get('input[name=userId]').type(user.userName)
    cy.get('input[name=password]').type(user.password)
    cy.get('form').submit()
    cy.contains('Clear all todos')
  })
  it('should login failed when userId & password wrong', function () {
    cy.visit(`${baseUrl}`)
    cy.get('input[name=userId]').type(user.userName)
    cy.get('input[name=password]').type('123123123')
    cy.get('form').submit()
    cy.contains('Sign in')
  })

  it('should login failed when userId & password empty', function () {
    cy.visit(`${baseUrl}`)
    cy.get('form').submit()
    cy.contains('Sign in')
  })

  it('should login failed when userId & password is space', function () {
    cy.visit(`${baseUrl}`)
    cy.get('input[name=userId]').type('           ')
    cy.get('input[name=password]').type('           ')
    cy.get('form').submit()
    cy.contains('Sign in')
  })

  it('should redirect login page when visit todo page without authorized', function () {
    cy.visit(`${baseUrl}/todo`)
    cy.contains('Sign in')
  })
})
