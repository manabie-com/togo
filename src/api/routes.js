module.exports = function(app) {
  let tasksCtrl = require('./handlers/tasksHandler');
  // todoList Routes
  app.route('/tasks')
    .post(tasksCtrl.create)
};