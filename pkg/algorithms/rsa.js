const getRSAKeysImperative = (p, q) => {
  const n = p * q;

  const phi = (p - 1) * (q - 1);

  let e = 2;
  while (true) {
    if (phi % e != 0) {
      break;
    }

    e++;
  }

  let k = 1;
  while (true) {
    if ((k * phi + 1) % e == 0) {
      break;
    }

    k++;
  }

  const d = Math.floor((k * phi + 1) / e);

  return [n, e, d];
};

const rsaEncrypt = (m, e, n) => {
  if (m > n) {
    throw new Error("keyspace to short");
  }

  return Math.pow(m, e) % n;
};

const rsaDecrypt = (c, d, n) => Math.pow(c, d) % n;

const getRSAKeysFunctional = (p, q, k, e) => {
  /*
  	To make this fully side-effect free, replace all references
  	to the following with these expressions
  */
  const n = p * q;
  const phi = (p - 1) * (q - 1);

  if (phi % e != 0) {
    if ((k * phi + 1) % e == 0) {
      return [n, e, Math.floor((k * phi + 1) / e)];
    }

    return getRSAKeysFunctional(p, q, k + 1, e);
  }

  return getRSAKeysFunctional(p, q, k, e + 1);
};

module.exports = {
  getRSAKeysImperative,
  rsaEncrypt,
  rsaDecrypt,
  getRSAKeysFunctional,
};
