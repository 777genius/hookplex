import { execFileSync } from "node:child_process";
import { repoRoot } from "../config/site.mjs";

try {
  execFileSync("git", ["diff", "--quiet", "--", "website/generated"], {
    cwd: repoRoot,
    stdio: "inherit"
  });
} catch {
  console.error("Generated docs drift detected under website/generated.");
  process.exit(1);
}
