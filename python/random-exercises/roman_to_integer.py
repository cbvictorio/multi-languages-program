roman_numbers_map = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}


def roman_to_integer(text: str) -> int:
    i = 0
    total_sum = 0

    while i <= len(text) - 1:
        curr_char = text[i]
        current_value = roman_numbers_map[curr_char]

        if i == len(text) - 1:
            return total_sum + current_value

        next_char = text[i + 1]
        next_value = roman_numbers_map[next_char]

        if current_value < next_value:
            total_sum += next_value - current_value
            i += 2
            continue

        total_sum += current_value
        i += 1

    return total_sum


print(roman_to_integer("III"))
