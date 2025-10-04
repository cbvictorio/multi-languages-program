import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import { readContentFromFile } from "./file-io";
import { ConvertTextToArray, createWordsLengthMap } from "./utils";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const textFileRelativePath = "../../python/text-analysis-tool/text_file.txt";
const filePath = resolve(__dirname, textFileRelativePath);

async function main() {
  const fileContent = await readContentFromFile(filePath);
  const fileContentToArray = ConvertTextToArray(fileContent);
  const wordsLengthMap = createWordsLengthMap(fileContentToArray);

  print("\noutput:");
  print(wordsLengthMap);
}

function print(param: any) {
  console.log(param);
}

main();
