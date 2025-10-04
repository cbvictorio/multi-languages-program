from dataclasses import dataclass
import re


@dataclass
class WordData:
    count: int
    words: set[str]
    unique_words_count: int


# Comments for this function can be found in the Go script
# for the same module's name
def convert_text_to_array(text: str) -> list[str]:
    # Normalize em-dash
    text = text.replace("—", "-")

    # Remove quotes and punctuation
    text = re.sub(r"['\"`,;:!?]", "", text)

    # Replace dots and line breaks with space
    text = re.sub(r"[.\n\r]", " ", text)

    # Replace special chars with space (keep alphanumeric, spaces, hyphens)
    text = re.sub(r"[^a-zA-Z0-9\s-]", " ", text)

    # Split on whitespace and filter empty strings
    words = text.split()

    return words


def create_words_length_map(words: list[str]) -> dict[int, int]:
    words_length_map: dict[int, WordData] = {}

    for word in words:
        word_length: int = len(word)
        new_word_count: int = 1
        new_words_list: set[str] = {word}
        unique_words_count: int = 1

        if word_length in words_length_map:
            current_word_data = words_length_map[word_length]
            new_word_count = current_word_data.count + 1
            new_words_list = current_word_data.words | {word}
            unique_words_count = len(new_words_list)

        word_data = WordData(new_word_count, new_words_list, unique_words_count)
        words_length_map[word_length] = word_data

    return words_length_map
