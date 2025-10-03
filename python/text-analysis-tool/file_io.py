import os


def read_content_from_file(filePath: str) -> str:
    try:
        with open(filePath, "r") as file:
            content = file.read()
            return content
    except Exception as e:
        print(f"Error: {e}")
