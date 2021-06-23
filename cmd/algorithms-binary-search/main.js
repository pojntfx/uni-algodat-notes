const {
  binarySearchImperative,
  binarySearchFunctional,
} = require("../../pkg/algorithms/binary_search.js");

const inputs = [
  [1, 6, 9, 10, 11],
  [1, 2, 3, 4, 5],
  [3, 4, 5, 7, 12, 50],
  [3, 4, 6, 7, 12, 50],
  [1, 2, 3, 4, 5, 6, 12, 50],
];
const target = 6;

inputs.forEach((input) => {
  const outputImperative = binarySearchImperative(input, target);
  const outputFunctional = binarySearchFunctional(
    input,
    target,
    0,
    input.length - 1
  );

  console.log(`in=${input} out=${outputImperative}=${outputFunctional}`);
});
