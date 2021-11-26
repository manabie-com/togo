import { hri } from 'human-readable-ids'
import 'cypress-wait-until'

Cypress.on('uncaught:exception', () => false)

const env = Cypress.env()
const { baseUrlConsole: baseUrl, user } = env
const newTodoName = hri.random()

describe('Test todo', function () {
  it('should get todos successful', function () {
    cy.visit(`${baseUrl}`)
    cy.login(user)
    cy.get('.ToDo__item').find('input').should('be.have', 4)
  })

  it('should redirect todo page when visit login page when authorized', function () {
    cy.visit(`${baseUrl}`)
    cy.login(user)
    cy.visit(`${baseUrl}`)
    cy.contains('Clear all todos')
  })

  it('should create new todo successful', function () {
    cy.visit(`${baseUrl}`)
    cy.login(user)
    cy.get("[name='new-todo']").type(newTodoName).type('{enter}')
    cy.contains(newTodoName)
    cy.get('.ToDo__item').find('input').should('be.have', 5)
  })

  it('should delete todo by name successful', function () {
    cy.visit(`${baseUrl}`)
    cy.login(user)
    cy.get('.ToDo__item').within(() => cy.get('button').first().click())
    cy.contains(newTodoName).should('not.exist')
    cy.get('.ToDo__item').find('input').should('be.have', 3)
  })

  it('should filter with todo ative successful', function () {
    cy.visit(`${baseUrl}`)
    cy.login(user)
    cy.contains('Active').click()
    cy.get('.ToDo__item').find('input').should('be.have', 2)
  })

  it('should filter with todo completed successful', function () {
    cy.visit(`${baseUrl}`)
    cy.login(user)
    cy.contains('Completed').click()
    cy.get('.ToDo__item').find('input').should('be.have', 2)
  })

  it('should clear all todos successful', function () {
    cy.visit(`${baseUrl}`)
    cy.login(user)
    cy.contains('Clear all todos').click()
    cy.get('.ToDo__item').should('not.exist')
  })
})
