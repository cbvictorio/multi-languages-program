from math import floor


class MinHeap:
    def __init__(self) -> None:
        self.array = []
        pass

    def get_left_child(self, index: int) -> int:
        return (2 * index) + 1

    def get_right_child(self, index: int) -> int:
        return (2 * index) + 2

    def get_parent(self, index: int) -> int:
        return floor((index - 1) / 2)

    def is_empty(self) -> bool:
        return len(self.array) == 0

    def size(self) -> int:
        return len(self.array)

    def has_element_left_child(self, index: int) -> bool:
        left_child_index = self.get_left_child(index)
        return len(self.array) > left_child_index

    def has_element_right_child(self, index: int) -> bool:
        right_child_index = self.get_right_child(index)
        return len(self.array) > right_child_index

    def heapify_up(self, index: int) -> None:
        if index == 0:
            return

        parent_index = self.get_parent(index)
        parent_value = self.array[parent_index]
        element_value = self.array[index]

        if element_value < parent_value:
            self.array[parent_index] = element_value
            self.array[index] = parent_value
            self.heapify_up(parent_index)

    def heapify_down(self, index: int) -> None:
        if index == len(self.array) - 1:
            return

        element_value = self.array[index]

        if self.has_element_left_child(index):
            left_child_index = self.get_left_child(index)
            left_child_value = self.array[left_child_index]
            if left_child_value < element_value:
                self.array[index] = left_child_value
                self.array[left_child_index] = element_value
                self.heapify_down(left_child_index)
                return

        if self.has_element_right_child(index):
            right_child_index = self.get_right_child(index)
            right_child_value = self.array[right_child_index]
            if right_child_value < element_value:
                self.array[index] = right_child_value
                self.array[right_child_index] = element_value
                self.heapify_down(right_child_index)

    def extract_min(self) -> int:
        if self.is_empty():
            print("the heap is empty")
            return

        min_value = self.array[0]

        # extract last element index and value
        last_element_index = len(self.array) - 1
        last_element_value = self.array[last_element_index]

        # replace root position with the last element value
        self.array[0] = last_element_value

        # remove the last element
        self.array = self.array[:last_element_index]

        self.heapify_down(0)

        return min_value

    def print_values(self) -> None:
        if self.is_empty():
            print("the heap is empty")
            return

        print("\n")
        s = ""
        for element in self.array:
            s = f"{s} [{element}]"

        print(s)

    def insert(self, value: int) -> None:
        self.array.append(value)

        if len(self.array) < 2:
            return

        recently_added_index = len(self.array) - 1
        self.heapify_up(recently_added_index)


def main():
    min_heap = MinHeap()
    min_heap.insert(24)
    min_heap.insert(7)
    min_heap.insert(11)
    min_heap.insert(2)
    min_heap.insert(41)
    min_heap.insert(30)
    min_heap.insert(20)
    min_heap.print_values()

    min_heap_minimum_value = min_heap.extract_min()
    print(f"the minimum value of the heap is: {min_heap_minimum_value}")
    min_heap.print_values()


main()
