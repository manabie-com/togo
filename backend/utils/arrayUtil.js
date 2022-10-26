/**
 * ex: getArrayKeyExclude( { a:1, b:2, c:3 }, ['a']) => ['b', 'c']
 */
const getArrayKeyExclude = (inputObj, excludeArr = []) => {
  const arrKey = Object.keys(inputObj);
  return arrKey.filter((item) => !excludeArr.includes(item));
};

/**
 * ex: validateArray( ['x', 'y', 'z'], ['x', 'z', 'a']) => { missingElements: ['a'], falseElements: ['z'] }
 */
const validateArray = (inputArr, trueArr = []) => {
  const { missingElements, falseElements } = trueArr.reduce(
    (acc, cur, index) => {
      if (!inputArr.includes(cur)) {
        acc.missingElements.push(cur);
      } else {
        if (inputArr[index] !== cur) acc.falseElements.push(cur);
      }
      return acc;
    },
    {
      missingElements: [],
      falseElements: [],
    }
  );
  return { missingElements, falseElements };
};
module.exports = {
  getArrayKeyExclude,
  validateArray,
};
