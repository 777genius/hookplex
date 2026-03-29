import fs from "node:fs";

function readStdin(): string {
  return fs.readFileSync(0, "utf8");
}

function handleClaude(hookName: string): number {
  switch (hookName) {
    case "Stop":
    case "PreToolUse":
    case "UserPromptSubmit":
      break;
    default:
      process.stderr.write("usage: main.ts <Stop|PreToolUse|UserPromptSubmit>\n");
      return 1;
  }
  const event = JSON.parse(readStdin()) as Record<string, unknown>;
  void event;
  process.stdout.write("{}");
  return 0;
}

function main(): number {
  const hookName = process.argv[2];
  if (!hookName) {
    process.stderr.write("usage: main.ts <Stop|PreToolUse|UserPromptSubmit>\n");
    return 1;
  }
  return handleClaude(hookName);
}

process.exit(main());
