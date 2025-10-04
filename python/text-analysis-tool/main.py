import os
from file_io import read_content_from_file
from utils import (
    convert_text_to_array,
    create_words_length_map,
    create_words_dictionary,
    WordData,
)


def main():
    script_dir = os.path.dirname(os.path.abspath(__file__))
    file_path = os.path.join(script_dir, "../../files/text_file.txt")
    file_content: str = read_content_from_file(file_path)
    file_to_array: list[str] = convert_text_to_array(file_content)
    words_dictionary: dict[string, int] = create_words_dictionary(file_to_array)
    words_length_map: dict[int, WordData] = create_words_length_map(file_to_array)

    print("output:")
    print(words_dictionary)
    # for key in words_length_map:
    #     print({"key": key, "value": words_length_map[key]}, end="\n\n")


if __name__ == "__main__":
    main()
