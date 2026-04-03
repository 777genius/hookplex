package gemini

import (
	"encoding/json"
	"fmt"
	"strings"

	internalgemini "github.com/777genius/plugin-kit-ai/sdk/internal/platforms/gemini"
	"github.com/777genius/plugin-kit-ai/sdk/internal/runtime"
)

// SessionStartEvent is the Gemini SessionStart hook input.
type SessionStartEvent = internalgemini.SessionStartInput

// SessionEndEvent is the Gemini SessionEnd hook input.
type SessionEndEvent = internalgemini.SessionEndInput

// BeforeAgentEvent is the Gemini BeforeAgent hook input.
type BeforeAgentEvent = internalgemini.BeforeAgentInput

// AfterAgentEvent is the Gemini AfterAgent hook input.
type AfterAgentEvent = internalgemini.AfterAgentInput

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

// BeforeAgentResponse is the BeforeAgent response type.
type BeforeAgentResponse struct {
	CommonResponse
	AdditionalContext string
}

// AfterAgentResponse is the AfterAgent response type.
type AfterAgentResponse struct {
	CommonResponse
	ClearContext bool
}

// BeforeToolResponse is the BeforeTool response type.
type BeforeToolResponse struct {
	CommonResponse
	ToolInput json.RawMessage
}

// TailToolCallRequest requests an immediate follow-up tool execution from an
// AfterTool hook.
type TailToolCallRequest struct {
	Name string
	Args json.RawMessage
}

// AfterToolResponse is the AfterTool response type.
type AfterToolResponse struct {
	CommonResponse
	AdditionalContext   string
	TailToolCallRequest *TailToolCallRequest
}

// SessionStartContinue returns an explicit no-op SessionStart response.
func SessionStartContinue() *SessionStartResponse {
	return &SessionStartResponse{}
}

// SessionStartAddContext appends additional context during SessionStart.
func SessionStartAddContext(context string) *SessionStartResponse {
	return &SessionStartResponse{AdditionalContext: context}
}

// SessionEndContinue returns an explicit no-op SessionEnd response.
func SessionEndContinue() *SessionEndResponse {
	return &SessionEndResponse{}
}

// BeforeAgentContinue returns an explicit no-op BeforeAgent response.
func BeforeAgentContinue() *BeforeAgentResponse {
	return &BeforeAgentResponse{}
}

// BeforeAgentAddContext appends additional context to the current turn prompt.
func BeforeAgentAddContext(context string) *BeforeAgentResponse {
	return &BeforeAgentResponse{AdditionalContext: context}
}

// BeforeAgentAllow returns an explicit allow decision for BeforeAgent.
func BeforeAgentAllow() *BeforeAgentResponse {
	return &BeforeAgentResponse{CommonResponse: CommonResponse{Decision: "allow"}}
}

// BeforeAgentDeny blocks the turn and discards the user's prompt from history.
func BeforeAgentDeny(reason string) *BeforeAgentResponse {
	return &BeforeAgentResponse{CommonResponse: CommonResponse{Decision: "deny", Reason: reason}}
}

// AfterAgentContinue returns an explicit no-op AfterAgent response.
func AfterAgentContinue() *AfterAgentResponse {
	return &AfterAgentResponse{}
}

// AfterAgentAllow returns an explicit allow decision for AfterAgent.
func AfterAgentAllow() *AfterAgentResponse {
	return &AfterAgentResponse{CommonResponse: CommonResponse{Decision: "allow"}}
}

// AfterAgentDeny rejects the response and requests a retry.
func AfterAgentDeny(reason string) *AfterAgentResponse {
	return &AfterAgentResponse{CommonResponse: CommonResponse{Decision: "deny", Reason: reason}}
}

// AfterAgentClearContext clears LLM conversation memory while preserving the
// UI display.
func AfterAgentClearContext() *AfterAgentResponse {
	return &AfterAgentResponse{ClearContext: true}
}

// BeforeToolContinue returns an explicit no-op BeforeTool response.
func BeforeToolContinue() *BeforeToolResponse {
	return &BeforeToolResponse{}
}

// BeforeToolAllow returns an explicit allow decision for BeforeTool.
func BeforeToolAllow() *BeforeToolResponse {
	return &BeforeToolResponse{CommonResponse: CommonResponse{Decision: "allow"}}
}

// BeforeToolDeny blocks the tool invocation with a deny decision.
func BeforeToolDeny(reason string) *BeforeToolResponse {
	return &BeforeToolResponse{CommonResponse: CommonResponse{Decision: "deny", Reason: reason}}
}

// BeforeToolRewriteInput continues with a rewritten tool_input payload.
func BeforeToolRewriteInput(input json.RawMessage) *BeforeToolResponse {
	return &BeforeToolResponse{ToolInput: input}
}

// BeforeToolRewriteInputValue marshals a replacement tool_input object for
// Gemini BeforeTool hooks. Gemini expects hookSpecificOutput.tool_input to be a
// JSON object, so non-object values return an error.
func BeforeToolRewriteInputValue(v any) (*BeforeToolResponse, error) {
	body, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("marshal Gemini tool_input rewrite: %w", err)
	}
	if !looksLikeJSONObject(body) {
		return nil, fmt.Errorf("marshal Gemini tool_input rewrite: expected JSON object")
	}
	return BeforeToolRewriteInput(body), nil
}

// AfterToolContinue returns an explicit no-op AfterTool response.
func AfterToolContinue() *AfterToolResponse {
	return &AfterToolResponse{}
}

// AfterToolAddContext appends additional text to the tool result sent back to
// the agent.
func AfterToolAddContext(context string) *AfterToolResponse {
	return &AfterToolResponse{AdditionalContext: context}
}

// AfterToolAllow returns an explicit allow decision for AfterTool.
func AfterToolAllow() *AfterToolResponse {
	return &AfterToolResponse{CommonResponse: CommonResponse{Decision: "allow"}}
}

// AfterToolDeny blocks the follow-up path with a deny decision.
func AfterToolDeny(reason string) *AfterToolResponse {
	return &AfterToolResponse{CommonResponse: CommonResponse{Decision: "deny", Reason: reason}}
}

// AfterToolTailCall requests an immediate follow-up tool invocation.
func AfterToolTailCall(name string, args json.RawMessage) *AfterToolResponse {
	return &AfterToolResponse{
		TailToolCallRequest: &TailToolCallRequest{
			Name: name,
			Args: args,
		},
	}
}

// AfterToolTailCallValue marshals a typed follow-up tool request. Gemini
// expects tailToolCallRequest.args to be a JSON object, so non-object values
// return an error.
func AfterToolTailCallValue(name string, args any) (*AfterToolResponse, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, fmt.Errorf("marshal Gemini tail tool call args: %w", err)
	}
	if !looksLikeJSONObject(body) {
		return nil, fmt.Errorf("marshal Gemini tail tool call args: expected JSON object")
	}
	return AfterToolTailCall(name, body), nil
}

// Deprecated: prefer BeforeToolAllow, BeforeToolContinue, AfterToolAllow, or
// AfterToolContinue. Gemini handlers return typed response structs, and these
// CommonResponse helpers are kept only for backward compatibility.
//
// AllowTool returns an explicit allow decision for BeforeTool or AfterTool.
func AllowTool() *CommonResponse {
	return &CommonResponse{Decision: "allow"}
}

// Deprecated: prefer BeforeToolDeny or AfterToolDeny. Gemini handlers return
// typed response structs, and this CommonResponse helper is kept only for
// backward compatibility.
//
// DenyTool returns a deny decision with a reason for BeforeTool or AfterTool.
func DenyTool(reason string) *CommonResponse {
	return &CommonResponse{Decision: "deny", Reason: reason}
}

func looksLikeJSONObject(body []byte) bool {
	return strings.HasPrefix(strings.TrimSpace(string(body)), "{")
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

func lifecycleOutcomeFromResponse(r *CommonResponse) internalgemini.CommonOutcome {
	out := commonOutcomeFromResponse(r)
	out.Continue = nil
	out.StopReason = ""
	out.Decision = ""
	out.Reason = ""
	return out
}

func sessionStartOutcomeFromResponse(r *SessionStartResponse) internalgemini.SessionStartOutcome {
	if r == nil {
		return internalgemini.SessionStartOutcome{}
	}
	return internalgemini.SessionStartOutcome{
		CommonOutcome:     lifecycleOutcomeFromResponse(&r.CommonResponse),
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
	return internalgemini.SessionEndOutcome{CommonOutcome: lifecycleOutcomeFromResponse(r)}
}

func beforeAgentOutcomeFromResponse(r *BeforeAgentResponse) internalgemini.BeforeAgentOutcome {
	if r == nil {
		return internalgemini.BeforeAgentOutcome{}
	}
	return internalgemini.BeforeAgentOutcome{
		CommonOutcome:     commonOutcomeFromResponse(&r.CommonResponse),
		AdditionalContext: r.AdditionalContext,
	}
}

func afterAgentOutcomeFromResponse(r *AfterAgentResponse) internalgemini.AfterAgentOutcome {
	if r == nil {
		return internalgemini.AfterAgentOutcome{}
	}
	return internalgemini.AfterAgentOutcome{
		CommonOutcome: commonOutcomeFromResponse(&r.CommonResponse),
		ClearContext:  r.ClearContext,
	}
}

func afterToolOutcomeFromResponse(r *AfterToolResponse) internalgemini.AfterToolOutcome {
	if r == nil {
		return internalgemini.AfterToolOutcome{}
	}
	out := internalgemini.AfterToolOutcome{
		CommonOutcome:     commonOutcomeFromResponse(&r.CommonResponse),
		AdditionalContext: r.AdditionalContext,
	}
	if r.TailToolCallRequest != nil {
		out.TailToolCallRequest = &internalgemini.TailToolCallRequest{
			Name: r.TailToolCallRequest.Name,
			Args: r.TailToolCallRequest.Args,
		}
	}
	return out
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

func wrapBeforeAgent(fn func(*BeforeAgentEvent) *BeforeAgentResponse) runtime.TypedHandler {
	return wrapGeminiHandler("BeforeAgent", fn, func(r *BeforeAgentResponse) any {
		return beforeAgentOutcomeFromResponse(r)
	})
}

func wrapAfterAgent(fn func(*AfterAgentEvent) *AfterAgentResponse) runtime.TypedHandler {
	return wrapGeminiHandler("AfterAgent", fn, func(r *AfterAgentResponse) any {
		return afterAgentOutcomeFromResponse(r)
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
