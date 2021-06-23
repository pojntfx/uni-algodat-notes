const binarySearchFunctional = (haystack, needle, low, high) => {
  // Needle not in haystack
  if (low > high) {
    return -1;
  }

  /*
  To make this fully side-effect free, replace all references
  to the following with these expressions
  */
  const focus = Math.floor((low + high) / 2);
  const center = haystack[focus];

  // Needle at focus
  if (needle == center) {
    return focus;
  }

  // Take left part
  if (needle < center) {
    return binarySearchFunctional(haystack, needle, low, focus - 1);
  }

  // Take right part
  return binarySearchFunctional(haystack, needle, focus + 1, high);
};

const binarySearchImperative = (haystack, needle) => {
  let low = 0;
  let high = haystack.length - 1;

  while (true) {
    // Needle not in haystack
    if (low > high) {
      return -1;
    }

    const focus = Math.floor((low + high) / 2);
    const center = haystack[focus];

    // Needle at focus
    if (needle == center) {
      return focus;
    }

    // Take left part
    if (needle < center) {
      high = focus - 1;
    } else {
      // Take right part
      low = focus + 1;
    }
  }
};

module.exports = {
  binarySearchFunctional,
  binarySearchImperative,
};
