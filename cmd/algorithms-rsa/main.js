const {
  getRSAKeysImperative,
  rsaEncrypt,
  rsaDecrypt,
  getRSAKeysFunctional,
} = require("../../pkg/algorithms/rsa");

const inputs = [
  [22, 5, 7], // m p q
  [2, 3, 7],
];

inputs.forEach((input) => {
  {
    const [m, p, q] = input;

    const [n, e, d] = getRSAKeysImperative(p, q);

    const c = rsaEncrypt(m, e, n);
    const m2 = rsaDecrypt(c, d, n);

    console.log(
      `imp: (p=${p} q=${q}) -> (n=${n} e=${e} d=${d}) -> (m=${m}) -> (c=${c}) -> (m2=${m2})`
    );
  }

  {
    const [m, p, q] = input;

    const [n, e, d] = getRSAKeysFunctional(p, q, 1, 2);

    const c = rsaEncrypt(m, e, n);
    const m2 = rsaDecrypt(c, d, n);

    console.log(
      `imp: (p=${p} q=${q}) -> (n=${n} e=${e} d=${d}) -> (m=${m}) -> (c=${c}) -> (m2=${m2})`
    );
  }
});
