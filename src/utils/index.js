function validateObject(schema, object) {
  const result = {
    valid: true,
    message: ''
  };

  try {
    const messages = schema.validate(object);
    if (messages?.error?.message) {
      result.valid = false;
      result.message = messages?.error?.message;
    }
  } catch (error) {
    result.message = error.message;
  }

  return result;
}

async function sleep(time) {
  return await new Promise((solve) => {
    setTimeout(() => {
      solve(true);
    }, time);
  })
}

module.exports = {
  validateObject,
  sleep
}