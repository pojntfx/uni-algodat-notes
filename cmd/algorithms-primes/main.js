const { primes } = require("../../pkg/algorithms/primes");

const inputs = [1, 4, 15, 20];

inputs.forEach((input) => {
  const output = primes(input, []);

  console.log(`in=${input} out=${output}`);
});
