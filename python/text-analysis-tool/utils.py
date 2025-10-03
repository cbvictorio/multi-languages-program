from dataclasses import dataclass


@dataclass
class WordData:
    count: int
    words: set[str]
    unique_words_count: int


def convert_string_to_array(string: str) -> list[str]:
    return string.split()


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
