package gemini

import (
	"encoding/json"

	internalgemini "github.com/777genius/plugin-kit-ai/sdk/internal/platforms/gemini"
	"github.com/777genius/plugin-kit-ai/sdk/internal/runtime"
)

// SessionStartEvent is the Gemini SessionStart hook input.
type SessionStartEvent = internalgemini.SessionStartInput

// SessionEndEvent is the Gemini SessionEnd hook input.
type SessionEndEvent = internalgemini.SessionEndInput

// BeforeToolEvent is the Gemini BeforeTool hook input.
type BeforeToolEvent = internalgemini.BeforeToolInput

// AfterToolEvent is the Gemini AfterTool hook input.
type AfterToolEvent = internalgemini.AfterToolInput

// CommonResponse contains fields shared by Gemini's synchronous hook envelope.
type CommonResponse struct {
	Continue       *bool
	SuppressOutput bool
	StopReason     string
	Decision       string
	Reason         string
	SystemMessage  string
}

// SessionStartResponse is the SessionStart response type.
type SessionStartResponse struct {
	CommonResponse
	AdditionalContext string
}

// SessionEndResponse is the SessionEnd response type.
type SessionEndResponse = CommonResponse

// BeforeToolResponse is the BeforeTool response type.
type BeforeToolResponse struct {
	CommonResponse
	ToolInput json.RawMessage
}

// AfterToolResponse is the AfterTool response type.
type AfterToolResponse = CommonResponse

// AllowTool returns an explicit allow decision for BeforeTool or AfterTool.
func AllowTool() *CommonResponse {
	return &CommonResponse{Decision: "allow"}
}

// DenyTool returns a deny decision with a reason for BeforeTool or AfterTool.
func DenyTool(reason string) *CommonResponse {
	return &CommonResponse{Decision: "deny", Reason: reason}
}

func commonOutcomeFromResponse(r *CommonResponse) internalgemini.CommonOutcome {
	if r == nil {
		return internalgemini.CommonOutcome{}
	}
	return internalgemini.CommonOutcome{
		Continue:       r.Continue,
		SuppressOutput: r.SuppressOutput,
		StopReason:     r.StopReason,
		Decision:       r.Decision,
		Reason:         r.Reason,
		SystemMessage:  r.SystemMessage,
	}
}

func sessionStartOutcomeFromResponse(r *SessionStartResponse) internalgemini.SessionStartOutcome {
	if r == nil {
		return internalgemini.SessionStartOutcome{}
	}
	return internalgemini.SessionStartOutcome{
		CommonOutcome:     commonOutcomeFromResponse(&r.CommonResponse),
		AdditionalContext: r.AdditionalContext,
	}
}

func beforeToolOutcomeFromResponse(r *BeforeToolResponse) internalgemini.BeforeToolOutcome {
	if r == nil {
		return internalgemini.BeforeToolOutcome{}
	}
	return internalgemini.BeforeToolOutcome{
		CommonOutcome: commonOutcomeFromResponse(&r.CommonResponse),
		ToolInput:     r.ToolInput,
	}
}

func sessionEndOutcomeFromResponse(r *SessionEndResponse) internalgemini.SessionEndOutcome {
	return internalgemini.SessionEndOutcome{CommonOutcome: commonOutcomeFromResponse(r)}
}

func afterToolOutcomeFromResponse(r *AfterToolResponse) internalgemini.AfterToolOutcome {
	return internalgemini.AfterToolOutcome{CommonOutcome: commonOutcomeFromResponse(r)}
}

func wrapSessionStart(fn func(*SessionStartEvent) *SessionStartResponse) runtime.TypedHandler {
	return wrapGeminiHandler("SessionStart", fn, func(r *SessionStartResponse) any {
		return sessionStartOutcomeFromResponse(r)
	})
}

func wrapSessionEnd(fn func(*SessionEndEvent) *SessionEndResponse) runtime.TypedHandler {
	return wrapGeminiHandler("SessionEnd", fn, func(r *SessionEndResponse) any {
		return sessionEndOutcomeFromResponse(r)
	})
}

func wrapBeforeTool(fn func(*BeforeToolEvent) *BeforeToolResponse) runtime.TypedHandler {
	return wrapGeminiHandler("BeforeTool", fn, func(r *BeforeToolResponse) any {
		return beforeToolOutcomeFromResponse(r)
	})
}

func wrapAfterTool(fn func(*AfterToolEvent) *AfterToolResponse) runtime.TypedHandler {
	return wrapGeminiHandler("AfterTool", fn, func(r *AfterToolResponse) any {
		return afterToolOutcomeFromResponse(r)
	})
}

func internalgeminiTypeMismatch(name string) error {
	return runtime.InternalHookTypeMismatch(name)
}
