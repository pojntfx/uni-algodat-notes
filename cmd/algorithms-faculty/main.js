const { getFaculty } = require("../../pkg/algorithms/faculty");

const inputs = [1, 4, 15, 20];

inputs.forEach((input) => {
  const output = getFaculty(input, 1);

  console.log(`in=${input} out=${output}`);
});
