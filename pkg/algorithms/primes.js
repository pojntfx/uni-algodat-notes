const primes = (i, res) => {
  if (i < 2) {
    return res;
  }

  if (isPrime(i, i - 1)) {
    return primes(i - 1, [i, ...res]);
  }

  return primes(i - 1, res);
};

const isPrime = (number, i) => {
  if (i < 2) {
    return true;
  }
  if (number % i == 0) {
    return false;
  }

  return isPrime(number, i - 1);
};

module.exports = {
  primes,
  isPrime,
};
