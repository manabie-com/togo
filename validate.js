const generateTask = (name, userId) => {
  return {
    name: name.trim(),
    userId: +userId,
    completed: false,
  }
}

const validateInput = (name, userId) => {
  if (!name || !userId || name.trim().length === 0 || +userId === NaN) {
    return false
  }
  return true
}

exports.checkAndGenerate = (name, userId) => {
  if (!validateInput(name, userId)) {
    return false
  }
  return generateTask(name, userId)
}

exports.generateTask = generateTask
exports.validateInput = validateInput
