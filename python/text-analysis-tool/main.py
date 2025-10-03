from file_io import read_content_from_file
import os


def main():
    filePath: str = os.path.abspath("../../files/text_file.txt")
    fileContent: str = read_content_from_file(filePath)
    print("the current file content is: ")
    print(fileContent)


if __name__ == "__main__":
    main()
