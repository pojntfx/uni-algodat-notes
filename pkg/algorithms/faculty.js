const getFaculty = (i, res) => {
  if (i < 1) {
    return res;
  }

  return getFaculty(i - 1, res * i);
};

module.exports = { getFaculty };
