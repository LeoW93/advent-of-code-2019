const fs = require("fs");

function replacePositionWithValue(string, position, value) {
  const array = mapToArray(string);

  return [
    ...array.slice(0, position),
    value,
    ...array.slice(position + 1)
  ].join(',')
}

const mapToArray = string => {
  return string.split(',').map(s => {
    if (s === "0") return 0;
    return parseInt(s);
  })
}

function generateOutput(_input) {
  const input = mapToArray(_input)

  let output = input;
  for (let i = 0; i < input.length; i++) {
    const value = input[i];

    if (value === 1) {
      // adding
      output[input[i + 3]] = input[input[i + 1]] + input[input[i + 2]];
      // skip
      i += 3;
    }

    if (value === 2) {
      // multiplying
      output[input[i + 3]] = input[input[i + 1]] * input[input[i + 2]];
      // skip
      i += 3;
    }

    if (value === 99) {
      break;
    }
  }

  return output;
}

const originalText = fs.readFileSync("./input.txt").toString()

const inputText = replacePositionWithValue(
  replacePositionWithValue(originalText, 1, "12"),
  2,
  "2"
);


console.log(generateOutput(inputText)[0]);

const result = 19690720;

for (let i = 0;i <= 99; i++ ) {
  for (let j = 0;  j <= 99; j++) {
    const test = generateOutput(replacePositionWithValue(
      replacePositionWithValue(
        originalText, 1, `${i}`
      ),
      2,
      `${j}`
    ))[0];
    if (test === result) {
      console.log({
        i,
        j,
        result: (100* i) + j
      })
      break;
    } 
  }
}


