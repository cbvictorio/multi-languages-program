export function convertStringToArray(str: string): string[] {
  return str
    .replace(/—/g, "-") // Normalize em-dash to hyphen
    .replace(/['"`,;:!?]/g, "")
    .replace(/[.\n\r]/g, " ")
    .replace(/[^a-zA-Z0-9\s-]/g, " ")
    .split(/\s+/)
    .filter((word) => word.length > 0);
}

type WordLengthDictionary = {
  totalCount: number;
  uniqueWordsCount: number;
  words: string[];
};

export function createWordsLengthMap(
  words: string[]
): Record<number, WordLengthDictionary> {
  const wordsMap: Record<number, WordLengthDictionary> = {};

  for (const word of words) {
    const wordLength = word.length;

    const hasWordBeenRecorded = wordsMap[wordLength];

    const newCountValue = hasWordBeenRecorded
      ? wordsMap[wordLength].totalCount + 1
      : 1;

    const newWordsArray = hasWordBeenRecorded
      ? [...wordsMap[wordLength].words, word]
      : [word];

    const filteredUniqueWords = [...new Set(newWordsArray)];

    wordsMap[wordLength] = {
      totalCount: newCountValue,
      uniqueWordsCount: filteredUniqueWords.length,
      words: filteredUniqueWords,
    };
  }

  return wordsMap;
}
