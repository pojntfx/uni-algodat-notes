const getLeastCommonMultiple = (x, y, z) => {
  if ((x * z) % y == 0) {
    return x * z;
  }

  return getLeastCommonMultiple(x, y, z + 1);
};

module.exports = {
  getLeastCommonMultiple,
};
