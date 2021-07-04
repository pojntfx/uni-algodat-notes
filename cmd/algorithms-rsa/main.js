const {
  getRSAKeysImperative,
  rsaEncrypt,
  rsaDecrypt,
  getRSAKeysFunctional,
} = require("../../pkg/algorithms/rsa");

const inputs = [
  [22, 5, 7, false], // m p q "false -> can't fail, true -> can fail"
  [2, 3, 7, false], // m p q "false -> can't fail, true -> can fail"
  [69, 5, 7, true], // m p q "false -> can't fail, true -> can fail"
];

inputs.forEach((input) => {
  {
    const [m, p, q, canFail] = input;

    const [n, e, d] = getRSAKeysImperative(p, q);

    let c = 0;
    try {
      c = rsaEncrypt(m, e, n);
    } catch (e) {
      if (canFail) {
        console.log(
          `imp: (p=${p} q=${q}) -> (n=${n} e=${e} d=${d}) -> skipped because keyspace to short`
        );

        return;
      }

      throw e;
    }
    const m2 = rsaDecrypt(c, d, n);

    console.log(
      `imp: (p=${p} q=${q}) -> (n=${n} e=${e} d=${d}) -> (m=${m}) -> (c=${c}) -> (m2=${m2})`
    );
  }

  {
    const [m, p, q, canFail] = input;

    const [n, e, d] = getRSAKeysFunctional(p, q, 1, 2);

    let c = 0;
    try {
      c = rsaEncrypt(m, e, n);
    } catch (e) {
      if (canFail) {
        console.log(
          `imp: (p=${p} q=${q}) -> (n=${n} e=${e} d=${d}) -> skipped because keyspace to short`
        );

        return;
      }

      throw e;
    }
    const m2 = rsaDecrypt(c, d, n);

    console.log(
      `imp: (p=${p} q=${q}) -> (n=${n} e=${e} d=${d}) -> (m=${m}) -> (c=${c}) -> (m2=${m2})`
    );
  }
});
