function handleNotify(): number {
  const payload = process.argv[3];
  if (!payload) {
    process.stderr.write("missing notify payload\n");
    return 1;
  }
  const event = JSON.parse(payload) as Record<string, unknown>;
  void event;
  return 0;
}

function main(): number {
  const hookName = process.argv[2];
  if (hookName !== "notify") {
    process.stderr.write("usage: main.ts notify <json-payload>\n");
    return 1;
  }
  return handleNotify();
}

process.exit(main());
