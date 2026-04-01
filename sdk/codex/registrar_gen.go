package codex

func generatedRegistrarMarker() {}

// OnNotify registers a handler for the Codex Notify.
func (r *Registrar) OnNotify(fn func(*NotifyEvent) *Response) {
	r.backend.Register("codex", "Notify", wrapNotify(fn))
}
