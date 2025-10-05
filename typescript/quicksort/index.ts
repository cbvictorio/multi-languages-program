function quicksort(numbers: number[]): number[] {
  const numberOfElements = numbers.length;

  if (numberOfElements <= 1) {
    return numbers;
  }

  const pivotIndex = numbers.length - 1;
  const pivotValue = numbers[pivotIndex];
  const leftSideNumbers: number[] = [];
  const rightSideNumbers: number[] = [];

  for (const number of numbers) {
    if (number < pivotValue) {
      leftSideNumbers.push(number);
      continue;
    }

    if (number > pivotValue) {
      rightSideNumbers.push(number);
      continue;
    }
  }

  const sortedLeftSide = quicksort(leftSideNumbers);
  const sortedRightSide = quicksort(rightSideNumbers);

  return [...sortedLeftSide, pivotValue, ...sortedRightSide];
}

function main() {
  const array = [2, 3, 1, 0];
  const result = quicksort(array);
  console.log("result: ", result);
}

main();
