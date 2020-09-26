function getGroupsOfAdjacentDigits(number) {
  const str = number.toString();
  return str.split("").reduce((groups, digit) => {
    const finder = (g) => g.some((elm) => elm === digit);
    const group = groups.find(finder);
    const index = groups.findIndex(finder);

    if (group) {
      return [
        ...groups.slice(0, index),
        [...group, digit],
        ...groups.slice(index + 1),
      ];
    }

    return [...groups, [digit]];
  }, []);
}

function hasTwoAdjacentDigits(number) {
  return getGroupsOfAdjacentDigits(number).some((group) => group.length === 2);
}

function digitsNeverDecrease(number) {
  const str = number.toString();

  for (let i = 0; i < str.length - 1; i++) {
    if (str[i] > str[i + 1]) {
      return false;
    }
  }
  return true;
}

function main() {
  const min = 245318;
  const max = 765747;

  let matches = [];

  for (i = min; i <= max; i++) {
    if (hasTwoAdjacentDigits(i) && digitsNeverDecrease(i)) {
      matches.push(i);
      console.log(i);
    }
  }

  console.log("n = ", matches.length);
}

main();
