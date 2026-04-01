import path from "node:path";
import { sourceRoot } from "../config/site.mjs";
import { listMarkdownFiles } from "../lib/fs.mjs";

const enFiles = await listMarkdownFiles(path.join(sourceRoot, "en"));
const ruFiles = await listMarkdownFiles(path.join(sourceRoot, "ru"));

const normalized = (root, files) => new Set(files.map((filePath) => path.relative(root, filePath)));
const enSet = normalized(path.join(sourceRoot, "en"), enFiles);
const ruSet = normalized(path.join(sourceRoot, "ru"), ruFiles);

const missingInRu = [...enSet].filter((filePath) => !ruSet.has(filePath));
const missingInEn = [...ruSet].filter((filePath) => !enSet.has(filePath));

if (missingInRu.length || missingInEn.length) {
  console.error("Locale parity check failed.");
  if (missingInRu.length) {
    console.error("Missing in ru:", missingInRu.join(", "));
  }
  if (missingInEn.length) {
    console.error("Missing in en:", missingInEn.join(", "));
  }
  process.exit(1);
}
