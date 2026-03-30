export type JSONMap = Record<string, unknown>;
export type ClaudeHandler = (event: JSONMap) => JSONMap | void;
export type CodexHandler = (event: JSONMap) => number | void;

export declare const CLAUDE_STABLE_HOOKS: readonly string[];
export declare const CLAUDE_EXTENDED_HOOKS: readonly string[];

export declare function allow(): Record<string, never>;
export declare function continue_(): number;

export declare class ClaudeApp {
  constructor(options: { allowedHooks: string[] | readonly string[]; usage: string });
  on(hookName: string, handler: ClaudeHandler): this;
  onStop(handler: ClaudeHandler): this;
  onPreToolUse(handler: ClaudeHandler): this;
  onUserPromptSubmit(handler: ClaudeHandler): this;
  run(): number;
}

export declare class CodexApp {
  onNotify(handler: CodexHandler): this;
  run(): number;
}
