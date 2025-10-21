from is_unique import is_unique
from check_permutation import check_permutation


def main():
    check_text_uniqueness = is_unique("aas")
    check_string_permutation = check_permutation("ab", "eidbaooo")
    print(check_text_uniqueness)
    print(check_string_permutation)


main()
