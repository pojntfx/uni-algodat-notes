const {
  getAsBinaryImperative,
  getAsBinaryFunctional,
} = require("../../pkg/algorithms/binary");

const inputs = [5, 6, 9, 15, 31, 66, 242];

inputs.forEach((input) => {
  const outputImperative = getAsBinaryImperative(input);
  const outputFunctional = getAsBinaryFunctional(input, "");

  console.log(
    `in=${String(input).padEnd(3, " ")} out=${String(outputImperative).padStart(
      8,
      0
    )}=${String(outputFunctional).padStart(8, 0)}`
  );
});
