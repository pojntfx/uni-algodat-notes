const { getChecksumNumeric } = require("../../pkg/algorithms/checksum_numeric");

const inputs = [5, 234, 33, 399];

inputs.forEach((input) => {
  const output = getChecksumNumeric(input, 0);

  console.log(`in=${input} out=${output}`);
});
