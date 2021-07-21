const { getNumAsArray } = require("../../pkg/algorithms/num_to_array");

const inputs = [5, 234, 33, 399];

inputs.forEach((input) => {
  const output = getNumAsArray(input, []);

  console.log(`in=${input} out=${output}`);
});
