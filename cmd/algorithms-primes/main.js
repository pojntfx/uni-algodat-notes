const { search } = require("../../pkg/algorithms/text_search");

const inputs = [
  ["Hello, world!", "orl"],
  ["Hello, world!", "asdf"],
];

inputs.forEach((input) => {
  const output = search(input[0], input[1], 0);

  console.log(`in=${input} out=${output}`);
});
