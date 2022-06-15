const {check} = require('express-validator');

let validateTodo = () => {
  return [ 
    check('title', 'title does not Empty').not().isEmpty(),
    check('description', 'description does not Empty').not().isEmpty(),
    check('user_id', 'user_id does not Empty').not().isEmpty()
  ]; 
}

let validate = {
    validateTodo: validateTodo
};

module.exports = {validate};