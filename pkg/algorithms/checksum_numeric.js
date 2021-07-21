const getChecksumNumeric = (input, out) => {
  if (input == 0) {
    return out;
  }

  return getChecksumNumeric(Math.floor(input / 10), out + (input % 10));
};

module.exports = {
  getChecksumNumeric,
};
