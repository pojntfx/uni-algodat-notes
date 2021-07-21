const {
  getLeastCommonMultiple,
} = require("../../pkg/algorithms/least_common_multiple");

const inputs = [
  [1, 4],
  [4, 6],
  [15, 6],
  [6, 123],
  [234, 231],
];

inputs.forEach((input) => {
  const output = getLeastCommonMultiple(input[0], input[1], 1);

  console.log(`in=${input} out=${output}`);
});
