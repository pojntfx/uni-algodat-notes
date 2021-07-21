const getAverageImperative = (input) => {
  let count = 0;

  for (let i = 0; i < input.length; i++) {
    count += input[i];
  }

  return count / input.length;
};

const getAverageFunctional = (input, count, i) => {
  if (i >= input.length) {
    return count / input.length;
  }

  return getAverageFunctional(input, count + input[i], i + 1);
};

module.exports = {
  getAverageImperative,
  getAverageFunctional,
};
