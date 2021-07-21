const {
  getChecksumReducedImperative,
  getChecksumReducedFunctional,
} = require("../../pkg/algorithms/checksum_reduced");

const inputs = [
  [1, 4, 6],
  [1, 4, 6, 123, 838],
];

inputs.forEach((input) => {
  const outputImperative = getChecksumReducedImperative(input);
  const outputFunctional = getChecksumReducedFunctional(input, 0, 0);

  console.log(`in=${input} out=${outputImperative}=${outputFunctional}`);
});
