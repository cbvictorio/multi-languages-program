from file_io import read_content_from_file
from utils import convert_string_to_array
import os


def main():
    script_dir = os.path.dirname(os.path.abspath(__file__))
    file_path = os.path.join(script_dir, "../../files/text_file.txt")
    file_content: str = read_content_from_file(file_path)
    file_to_array: list[str] = convert_string_to_array(file_content)

    print("output:")
    print(file_to_array)


if __name__ == "__main__":
    main()
