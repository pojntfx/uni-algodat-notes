const getChecksumReducedImperative = (input) => {
  let check = 0;

  for (let i = 0; i < input.length; i++) {
    check += parseInt(input[i]);
  }

  if (check > 10) {
    return getChecksumReducedImperative(check.toString().split(""));
  }

  return check;
};

const getChecksumReducedFunctional = (input, check, i) => {
  if (i >= input.length) {
    if (check > 10) {
      return getChecksumReducedFunctional(check.toString().split(""), 0, 0);
    }

    return check;
  }

  return getChecksumReducedFunctional(input, check + parseInt(input[i]), i + 1);
};

module.exports = {
  getChecksumReducedImperative,
  getChecksumReducedFunctional,
};
