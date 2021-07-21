const getNumAsArray = (num, arr) => {
  if (num == 0) {
    return arr;
  }

  return getNumAsArray(Math.floor(num / 10), [num % 10, ...arr]);
};

module.exports = {
  getNumAsArray,
};
