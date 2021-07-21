const { getMedian } = require("../../pkg/algorithms/median");

const inputs = [[7, 13, 2, 1, 6, 5]];

inputs.forEach((input) => {
  const output = getMedian(input);

  console.log(`in=${input} out=${output}`);
});
