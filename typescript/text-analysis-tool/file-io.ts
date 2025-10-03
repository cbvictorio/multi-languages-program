import { readFile } from "fs/promises";

export async function readContentFromFile(filePath: string): Promise<string> {
  if (!filePath) {
    throw new Error("the filePath param is missing");
  }

  try {
    return await readFile(filePath, "utf-8");
  } catch (error) {
    throw error;
  }
}
