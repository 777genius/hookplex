import { tool } from "@opencode-ai/plugin"

export default tool({
  description: "Echo a short value from a standalone OpenCode tool file.",
  args: {
    value: tool.schema.string().describe("Value to echo back"),
  },
  async execute(args) {
    return {
      ok: true,
      value: args.value,
    }
  },
})
