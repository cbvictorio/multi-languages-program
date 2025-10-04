import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import { readContentFromFile } from "./file-io";
import {
  convertTextToArray,
  createWordsLengthMap,
  createWordsDictionary,
} from "./utils";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const textFileRelativePath = "../../python/text-analysis-tool/text_file.txt";
const filePath = resolve(__dirname, textFileRelativePath);

async function main() {
  const fileContent = await readContentFromFile(filePath);
  const fileContentToArray = convertTextToArray(fileContent);
  const wordsDictionary = createWordsDictionary(fileContentToArray);
  const wordsLengthMap = createWordsLengthMap(fileContentToArray);

  print("\noutput:");
  print(wordsDictionary);
}

function print(param: any) {
  console.log(param);
}

main();
