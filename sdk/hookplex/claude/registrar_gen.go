package claude

func generatedRegistrarMarker() {}

func (r *Registrar) OnPreToolUse(fn func(*PreToolUseEvent) *PreToolResponse) {
	r.backend.Register("claude", "PreToolUse", wrapPreToolUse(fn))
}

func (r *Registrar) OnStop(fn func(*StopEvent) *Response) {
	r.backend.Register("claude", "Stop", wrapStop(fn))
}

func (r *Registrar) OnUserPromptSubmit(fn func(*UserPromptEvent) *UserPromptResponse) {
	r.backend.Register("claude", "UserPromptSubmit", wrapUserPromptSubmit(fn))
}
