const search = (string, query, i) => {
  if (i >= string.length) {
    return -1;
  }

  if (subString(string, i, query.length, 0, "") == query) {
    return i + 1;
  }

  return search(string, query, i + 1);
};

const subString = (string, position, length, i, res) => {
  if (i >= length || position + i >= string.length) {
    return res;
  }

  return subString(string, position, length, i + 1, res + string[position + i]);
};

module.exports = {
  search,
  subString,
};
