const {
  getAverageImperative,
  getAverageFunctional,
} = require("../../pkg/algorithms/average");

const inputs = [[7, 13, 2, 1, 6, 5]];

inputs.forEach((input) => {
  const outputImperative = getAverageImperative(input);
  const outputFunctional = getAverageFunctional(input, 0, 0);

  console.log(`in=${input} out=${outputImperative}=${outputFunctional}`);
});
