import { MaxHeap } from "./max-heap";

function main() {
  // [5, 3, 8, 1] ====> [8, 3, 5, 1]
  const array = [15, 12, 10, 8, 2, 6, 3, 1];
  const heap = new MaxHeap();
  const secondHeap = new MaxHeap();
  array.forEach((n) => {
    heap.insert(n);
    secondHeap.insert(n);
  });

  // [8, 3, 5, 1] ====> [5, 3, 1]
  const extractedValue = heap.extractMax();
  console.log({
    originalHeap: secondHeap.printValues(),
    heap: heap.printValues(),
    extractedValue,
  });
}

main();
