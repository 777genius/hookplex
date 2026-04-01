import fs from "node:fs/promises";
import path from "node:path";

export async function ensureDir(dirPath) {
  await fs.mkdir(dirPath, { recursive: true });
}

export async function rimraf(dirPath) {
  await fs.rm(dirPath, { force: true, recursive: true });
}

export async function writeFile(filePath, content) {
  await ensureDir(path.dirname(filePath));
  await fs.writeFile(filePath, content, "utf8");
}

export async function writeJson(filePath, value) {
  await writeFile(filePath, JSON.stringify(value, null, 2) + "\n");
}

export async function copyTree(sourceDir, targetDir) {
  await ensureDir(targetDir);
  const entries = await fs.readdir(sourceDir, { withFileTypes: true });
  for (const entry of entries) {
    const sourcePath = path.join(sourceDir, entry.name);
    const targetPath = path.join(targetDir, entry.name);
    if (entry.isDirectory()) {
      await copyTree(sourcePath, targetPath);
      continue;
    }
    await ensureDir(path.dirname(targetPath));
    await fs.copyFile(sourcePath, targetPath);
  }
}

export async function listMarkdownFiles(rootDir) {
  const out = [];
  await walk(rootDir, out);
  return out.sort();
}

async function walk(currentDir, out) {
  const entries = await fs.readdir(currentDir, { withFileTypes: true });
  for (const entry of entries) {
    const currentPath = path.join(currentDir, entry.name);
    if (entry.isDirectory()) {
      await walk(currentPath, out);
      continue;
    }
    if (entry.name.endsWith(".md")) {
      out.push(currentPath);
    }
  }
}
