export class MaxHeap {
  heap: number[];

  constructor() {
    this.heap = [];
  }

  isEmpty(): boolean {
    return this.heap.length === 0;
  }

  getLeftChild(nodeIndex: number): number {
    return 2 * nodeIndex + 1;
  }

  getRightChild(nodeIndex: number): number {
    return 2 * nodeIndex + 2;
  }

  getParent(nodeIndex: number): number {
    return Math.floor((nodeIndex - 1) / 2);
  }

  peek(): number | undefined {
    if (this.isEmpty()) return;
    return this.heap[0];
  }

  swapElements(firstIndex: number, secondIndex: number) {
    const firstNodeValue = this.heap[firstIndex];
    const secondNodeValue = this.heap[secondIndex];
    this.heap[secondIndex] = secondNodeValue;
    this.heap[secondIndex] = firstNodeValue;
  }

  heapifyUp(nodeIndex: number) {
    if (nodeIndex === 0) return;

    const valueToHeap = this.heap[nodeIndex];
    const parentIndex = this.getParent(nodeIndex);
    const parentValue = this.heap[parentIndex];

    if (valueToHeap > parentValue) {
      this.swapElements(nodeIndex, parentIndex);
      this.heapifyUp(parentIndex);
    }
  }

  heapifyDown(nodeIndex: number) {
    if (nodeIndex === this.heap.length - 1) return;

    const nodeValue = this.heap[nodeIndex];
    const leftChildIndex = this.getLeftChild(nodeIndex);
    const rightChildIndex = this.getRightChild(nodeIndex);
    const leftChildValue = this.heap[leftChildIndex];
    const rightChildValue = this.heap[rightChildIndex];

    if (leftChildValue && leftChildValue > nodeValue) {
      this.swapElements(nodeIndex, leftChildIndex);
      return this.heapifyDown(leftChildIndex);
    }

    if (rightChildValue && rightChildValue > nodeValue) {
      this.swapElements(nodeIndex, rightChildIndex);
      this.heapifyDown(rightChildIndex);
    }
  }

  extractMax(): number | undefined {
    if (this.heap.length === 1) {
      return this.heap.pop();
    }

    const largerNodeValue = this.heap[0];
    const lastNodeIndex = this.heap.length - 1;

    this.heap[0] = this.heap[lastNodeIndex];
    this.heap.pop();

    this.heapifyDown(0);
    return largerNodeValue;
  }

  insert(n: number) {
    if (this.isEmpty()) {
      return this.heap.push(n);
    }

    this.heap.push(n);
    const newNodeIndex = this.heap.length - 1;
    this.heapifyUp(newNodeIndex);
  }

  printValues() {
    return this.heap;
  }
}
