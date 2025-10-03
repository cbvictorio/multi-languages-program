export function convertStringToArray(str: string): string[] {
  return str.split(" ");
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
