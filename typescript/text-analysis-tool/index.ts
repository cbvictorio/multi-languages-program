import { resolve, dirname } from "path";
import { fileURLToPath } from "url";
import { readContentFromFile } from "./file-io";

const __filename = fileURLToPath(import.meta.url);
const __dirname = dirname(__filename);
const textFileRelativePath = "../../python/text-analysis-tool/text_file.txt";

async function main() {
  const filePath = resolve(__dirname, textFileRelativePath);
  const content = await readContentFromFile(filePath);

  console.log("=== CONTENT ===");
  console.log(content);
  console.log("=== CONTENT ===");
}

main();
