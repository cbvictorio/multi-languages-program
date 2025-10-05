def quicksort(numbers: list[int]) -> list[int]:
    result: list[int] = []
    number_of_elements: int = len(numbers)

    if number_of_elements <= 1:
        return numbers

    pivot_index = number_of_elements - 1
    pivot_value = numbers[pivot_index]
    left_side_elements: list[int] = []
    right_side_elements: list[int] = []

    for number in numbers:
        if number < pivot_value:
            left_side_elements.append(number)
            continue

        if number > pivot_value:
            right_side_elements.append(number)
            continue

    sorted_left_side = quicksort(left_side_elements)
    sorted_right_side = quicksort(right_side_elements)
    result = [*sorted_left_side, pivot_value, *sorted_right_side]

    return result


def main():
    array = [2, 3, 4, 1, 0, 2, 2, 3, 0, 4]
    result = quicksort(array)
    print("output:")
    print(result)


if __name__ == "__main__":
    main()
