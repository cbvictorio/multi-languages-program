# Implement an algorithm to determine if a string has all unique characters.


def is_unique(text: str) -> bool:
    dictionary: dict[str, bool] = {}

    for char in text:
        if char in dictionary:
            return False

        dictionary[char] = True

    return True
