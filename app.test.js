const request = require('supertest')
const app = require('./app')
const { generateTask, validateInput, checkAndGenerate } = require('./validate')

/* UNIT TEST*/

test('should output a task', () => {
  const task = generateTask('do something', 1)
  expect(task).toEqual({
    name: 'do something',
    userId: 1,
    completed: false,
  })
})

/* INTEGRATION TEST*/

// test post successfully
it('POST /tasks', function () {
  return request(app)
    .post('/tasks')
    .send({ name: 'do something', userId: 1 })
    .set('Accept', 'application/json')
    .expect('Content-Type', /json/)
    .expect(201)
    .then((response) => {
      expect(response.body).toEqual(
        expect.objectContaining({
          id: expect.any(Number),
          name: expect.any(String),
          userId: expect.any(Number),
          completed: expect.any(Boolean),
        })
      )
    })
})

// test error 400 with userId is ''
it('POST /tasks', function () {
  return request(app)
    .post('/tasks')
    .send({ name: 'do something', userId: '' })
    .set('Accept', 'application/json')
    .expect('Content-Type', /json/)
    .expect(400)
    .then((response) => {
      expect(response.body).toEqual(
        expect.objectContaining({
          message:
            'Name and userId are required. Name should be a string and userId should be a number.',
        })
      )
    })
})

// test error 400 with name is ''
it('POST /tasks', function () {
  return request(app)
    .post('/tasks')
    .send({ name: '', userId: 1 })
    .set('Accept', 'application/json')
    .expect('Content-Type', /json/)
    .expect(400)
    .then((response) => {
      expect(response.body).toEqual(
        expect.objectContaining({
          message:
            'Name and userId are required. Name should be a string and userId should be a number.',
        })
      )
    })
})

// test error 400 with user does not exist
it('POST /tasks', function () {
  return request(app)
    .post('/tasks')
    .send({ name: 'do something', userId: 2 })
    .set('Accept', 'application/json')
    .expect('Content-Type', /json/)
    .expect(400)
    .then((response) => {
      expect(response.body).toEqual(
        expect.objectContaining({
          message: 'User does not exist.',
        })
      )
    })
})

// test error 400 with user's tasks reach limit
it('POST /tasks', function () {
  return request(app)
    .post('/tasks')
    .send({ name: 'do something', userId: 1 })
    .send({ name: 'do another', userId: 1 })
    .set('Accept', 'application/json')
    .expect('Content-Type', /json/)
    .expect(400)
    .then((response) => {
      expect(response.body).toEqual(
        expect.objectContaining({
          message:
            'This user already reached a maximum limit of tasks per day.',
        })
      )
    })
})
