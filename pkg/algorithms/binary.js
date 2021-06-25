const getAsBinaryImperative = (num) => {
  let output = "";

  for (let i = num; i > 0; i = Math.floor(i / 2)) {
    output = (i % 2) + output;
  }

  return output;
};

const getAsBinaryFunctional = (num, output) => {
  if (num == 0) {
    return output;
  }

  return getAsBinaryFunctional(Math.floor(num / 2), (num % 2) + output);
};

module.exports = {
  getAsBinaryImperative,
  getAsBinaryFunctional,
};
