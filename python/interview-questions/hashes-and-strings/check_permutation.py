def check_permutation(first_text: str, second_text: str) -> bool:
    sorted_first_text: str = "".join(sorted(first_text))
    loop_limit: int = len(second_text) - len(first_text)

    for i in range(0, loop_limit):
        index_step: int = i + len(first_text)
        sub_string: str = second_text[i:index_step]
        sorted_sub_string: str = "".join(sorted(sub_string))

        if sorted_first_text == sorted_sub_string:
            return True

    return False
