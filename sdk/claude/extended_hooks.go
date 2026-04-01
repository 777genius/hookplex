package claude

import (
	"encoding/json"

	internalclaude "github.com/777genius/plugin-kit-ai/sdk/internal/platforms/claude"
	"github.com/777genius/plugin-kit-ai/sdk/internal/runtime"
)

// SessionStartEvent is the Claude SessionStart hook input.
type SessionStartEvent = internalclaude.SessionStartInput

// SessionEndEvent is the Claude SessionEnd hook input.
type SessionEndEvent = internalclaude.SessionEndInput

// NotificationEvent is the Claude Notification hook input.
type NotificationEvent = internalclaude.NotificationInput

// PostToolUseEvent is the Claude PostToolUse hook input.
type PostToolUseEvent = internalclaude.PostToolUseInput

// PostToolUseFailureEvent is the Claude PostToolUseFailure hook input.
type PostToolUseFailureEvent = internalclaude.PostToolUseFailureInput

// PermissionRequestEvent is the Claude PermissionRequest hook input.
type PermissionRequestEvent = internalclaude.PermissionRequestInput

// SubagentStartEvent is the Claude SubagentStart hook input.
type SubagentStartEvent = internalclaude.SubagentStartInput

// SubagentStopEvent is the Claude SubagentStop hook input.
type SubagentStopEvent = internalclaude.SubagentStopInput

// PreCompactEvent is the Claude PreCompact hook input.
type PreCompactEvent = internalclaude.PreCompactInput

// SetupEvent is the Claude Setup hook input.
type SetupEvent = internalclaude.SetupInput

// TeammateIdleEvent is the Claude TeammateIdle hook input.
type TeammateIdleEvent = internalclaude.TeammateIdleInput

// TaskCompletedEvent is the Claude TaskCompleted hook input.
type TaskCompletedEvent = internalclaude.TaskCompletedInput

// ConfigChangeEvent is the Claude ConfigChange hook input.
type ConfigChangeEvent = internalclaude.ConfigChangeInput

// WorktreeCreateEvent is the Claude WorktreeCreate hook input.
type WorktreeCreateEvent = internalclaude.WorktreeCreateInput

// WorktreeRemoveEvent is the Claude WorktreeRemove hook input.
type WorktreeRemoveEvent = internalclaude.WorktreeRemoveInput

// PermissionBehavior enumerates allow or deny decisions for PermissionRequest.
type PermissionBehavior = internalclaude.PermissionBehavior

// PermissionUpdate mirrors a single permission rule update returned to Claude.
type PermissionUpdate = internalclaude.PermissionUpdate

// PermissionRuleValue mirrors the supported rule value payload for permission updates.
type PermissionRuleValue = internalclaude.PermissionRuleValue

const (
	// PermissionAllow approves the pending action.
	PermissionAllow PermissionBehavior = internalclaude.PermissionAllow
	// PermissionDeny rejects the pending action.
	PermissionDeny PermissionBehavior = internalclaude.PermissionDeny
)

// CommonResponse contains fields shared by Claude's common hook response envelope.
type CommonResponse struct {
	// Continue overrides the continue flag for compatible hooks when non-nil.
	Continue *bool
	// SuppressOutput omits the usual hook output message when true.
	SuppressOutput bool
	// StopReason explains why processing should stop.
	StopReason string
	// Decision carries the Claude decision token such as approve or block.
	Decision string
	// Reason carries a human-readable reason for the decision.
	Reason string
	// SystemMessage injects a system message into Claude when supported.
	SystemMessage string
}

// ContextResponse extends CommonResponse with additional context text.
type ContextResponse struct {
	CommonResponse
	// AdditionalContext appends hook-specific context visible to Claude.
	AdditionalContext string
}

// PostToolUseResponse extends CommonResponse with tool output overrides.
type PostToolUseResponse struct {
	CommonResponse
	// AdditionalContext appends hook-specific context visible to Claude.
	AdditionalContext string
	// UpdatedMCPToolOutput replaces the MCP tool output payload on the wire.
	UpdatedMCPToolOutput json.RawMessage
}

// PermissionDecision describes the approval or denial returned from PermissionRequest.
type PermissionDecision struct {
	// Behavior is either PermissionAllow or PermissionDeny.
	Behavior PermissionBehavior
	// UpdatedInput replaces the pending input before Claude continues.
	UpdatedInput json.RawMessage
	// UpdatedPermissions amends stored permission rules when approving.
	UpdatedPermissions []PermissionUpdate
	// Message explains the deny decision to the user.
	Message string
	// Interrupt asks Claude to interrupt the current flow after the message.
	Interrupt bool
}

// PermissionRequestResponse extends CommonResponse with a permission decision.
type PermissionRequestResponse struct {
	CommonResponse
	// Permission holds the allow or deny decision when one is returned.
	Permission *PermissionDecision
}

// SessionStartResponse is the response type for SessionStart.
type SessionStartResponse = ContextResponse

// NotificationResponse is the response type for Notification.
type NotificationResponse = ContextResponse

// PostToolUseFailureResponse is the response type for PostToolUseFailure.
type PostToolUseFailureResponse = ContextResponse

// SessionEndResponse is the response type for SessionEnd.
type SessionEndResponse = CommonResponse

// SubagentStartResponse is the response type for SubagentStart.
type SubagentStartResponse = ContextResponse

// SubagentStopResponse is the response type for SubagentStop.
type SubagentStopResponse = CommonResponse

// PreCompactResponse is the response type for PreCompact.
type PreCompactResponse = CommonResponse

// SetupResponse is the response type for Setup.
type SetupResponse = ContextResponse

// TeammateIdleResponse is the response type for TeammateIdle.
type TeammateIdleResponse = CommonResponse

// TaskCompletedResponse is the response type for TaskCompleted.
type TaskCompletedResponse = CommonResponse

// ConfigChangeResponse is the response type for ConfigChange.
type ConfigChangeResponse = CommonResponse

// WorktreeCreateResponse is the response type for WorktreeCreate.
type WorktreeCreateResponse = CommonResponse

// WorktreeRemoveResponse is the response type for WorktreeRemove.
type WorktreeRemoveResponse = CommonResponse

// PermissionApprove returns a response that approves the requested action.
func PermissionApprove() *PermissionRequestResponse {
	return &PermissionRequestResponse{
		Permission: &PermissionDecision{Behavior: PermissionAllow},
	}
}

// PermissionApproveWithUpdates approves the action and sends updated input or rules.
func PermissionApproveWithUpdates(input json.RawMessage, updates []PermissionUpdate) *PermissionRequestResponse {
	return &PermissionRequestResponse{
		Permission: &PermissionDecision{
			Behavior:           PermissionAllow,
			UpdatedInput:       input,
			UpdatedPermissions: updates,
		},
	}
}

// PermissionBlock rejects the action with a message and interrupt choice.
func PermissionBlock(message string, interrupt bool) *PermissionRequestResponse {
	return &PermissionRequestResponse{
		Permission: &PermissionDecision{
			Behavior:  PermissionDeny,
			Message:   message,
			Interrupt: interrupt,
		},
	}
}

func commonOutcomeFromResponse(r *CommonResponse) internalclaude.CommonOutcome {
	if r == nil {
		return internalclaude.CommonOutcome{}
	}
	return internalclaude.CommonOutcome{
		Continue:       r.Continue,
		SuppressOutput: r.SuppressOutput,
		StopReason:     r.StopReason,
		Decision:       r.Decision,
		Reason:         r.Reason,
		SystemMessage:  r.SystemMessage,
	}
}

func contextOutcomeFromResponse(r *ContextResponse) internalclaude.ContextOutcome {
	if r == nil {
		return internalclaude.ContextOutcome{}
	}
	return internalclaude.ContextOutcome{
		CommonOutcome:     commonOutcomeFromResponse(&r.CommonResponse),
		AdditionalContext: r.AdditionalContext,
	}
}

func postToolUseOutcomeFromResponse(r *PostToolUseResponse) internalclaude.PostToolUseOutcome {
	if r == nil {
		return internalclaude.PostToolUseOutcome{}
	}
	return internalclaude.PostToolUseOutcome{
		CommonOutcome:        commonOutcomeFromResponse(&r.CommonResponse),
		AdditionalContext:    r.AdditionalContext,
		UpdatedMCPToolOutput: r.UpdatedMCPToolOutput,
	}
}

func permissionOutcomeFromResponse(r *PermissionRequestResponse) internalclaude.PermissionRequestOutcome {
	if r == nil {
		return internalclaude.PermissionRequestOutcome{}
	}
	out := internalclaude.PermissionRequestOutcome{
		CommonOutcome: commonOutcomeFromResponse(&r.CommonResponse),
	}
	if r.Permission != nil {
		out.Permission = &internalclaude.PermissionDecision{
			Behavior:           internalclaude.PermissionBehavior(r.Permission.Behavior),
			UpdatedInput:       r.Permission.UpdatedInput,
			UpdatedPermissions: r.Permission.UpdatedPermissions,
			Message:            r.Permission.Message,
			Interrupt:          r.Permission.Interrupt,
		}
	}
	return out
}

func wrapSessionStart(fn func(*SessionStartEvent) *SessionStartResponse) runtime.TypedHandler {
	return wrapClaudeHandler("SessionStart", fn, func(r *SessionStartResponse) any {
		return contextOutcomeFromResponse(r)
	})
}

func wrapSessionEnd(fn func(*SessionEndEvent) *SessionEndResponse) runtime.TypedHandler {
	return wrapClaudeHandler("SessionEnd", fn, func(r *SessionEndResponse) any {
		return commonOutcomeFromResponse(r)
	})
}

func wrapNotification(fn func(*NotificationEvent) *NotificationResponse) runtime.TypedHandler {
	return wrapClaudeHandler("Notification", fn, func(r *NotificationResponse) any {
		return contextOutcomeFromResponse(r)
	})
}

func wrapPostToolUse(fn func(*PostToolUseEvent) *PostToolUseResponse) runtime.TypedHandler {
	return wrapClaudeHandler("PostToolUse", fn, func(r *PostToolUseResponse) any {
		return postToolUseOutcomeFromResponse(r)
	})
}

func wrapPostToolUseFailure(fn func(*PostToolUseFailureEvent) *PostToolUseFailureResponse) runtime.TypedHandler {
	return wrapClaudeHandler("PostToolUseFailure", fn, func(r *PostToolUseFailureResponse) any {
		return contextOutcomeFromResponse(r)
	})
}

func wrapPermissionRequest(fn func(*PermissionRequestEvent) *PermissionRequestResponse) runtime.TypedHandler {
	return wrapClaudeHandler("PermissionRequest", fn, func(r *PermissionRequestResponse) any {
		return permissionOutcomeFromResponse(r)
	})
}

func wrapSubagentStart(fn func(*SubagentStartEvent) *SubagentStartResponse) runtime.TypedHandler {
	return wrapClaudeHandler("SubagentStart", fn, func(r *SubagentStartResponse) any {
		return contextOutcomeFromResponse(r)
	})
}

func wrapSubagentStop(fn func(*SubagentStopEvent) *SubagentStopResponse) runtime.TypedHandler {
	return wrapClaudeHandler("SubagentStop", fn, func(r *SubagentStopResponse) any {
		return commonOutcomeFromResponse(r)
	})
}

func wrapPreCompact(fn func(*PreCompactEvent) *PreCompactResponse) runtime.TypedHandler {
	return wrapClaudeHandler("PreCompact", fn, func(r *PreCompactResponse) any {
		return commonOutcomeFromResponse(r)
	})
}

func wrapSetup(fn func(*SetupEvent) *SetupResponse) runtime.TypedHandler {
	return wrapClaudeHandler("Setup", fn, func(r *SetupResponse) any {
		return contextOutcomeFromResponse(r)
	})
}

func wrapTeammateIdle(fn func(*TeammateIdleEvent) *TeammateIdleResponse) runtime.TypedHandler {
	return wrapClaudeHandler("TeammateIdle", fn, func(r *TeammateIdleResponse) any {
		return commonOutcomeFromResponse(r)
	})
}

func wrapTaskCompleted(fn func(*TaskCompletedEvent) *TaskCompletedResponse) runtime.TypedHandler {
	return wrapClaudeHandler("TaskCompleted", fn, func(r *TaskCompletedResponse) any {
		return commonOutcomeFromResponse(r)
	})
}

func wrapConfigChange(fn func(*ConfigChangeEvent) *ConfigChangeResponse) runtime.TypedHandler {
	return wrapClaudeHandler("ConfigChange", fn, func(r *ConfigChangeResponse) any {
		return commonOutcomeFromResponse(r)
	})
}

func wrapWorktreeCreate(fn func(*WorktreeCreateEvent) *WorktreeCreateResponse) runtime.TypedHandler {
	return wrapClaudeHandler("WorktreeCreate", fn, func(r *WorktreeCreateResponse) any {
		return commonOutcomeFromResponse(r)
	})
}

func wrapWorktreeRemove(fn func(*WorktreeRemoveEvent) *WorktreeRemoveResponse) runtime.TypedHandler {
	return wrapClaudeHandler("WorktreeRemove", fn, func(r *WorktreeRemoveResponse) any {
		return commonOutcomeFromResponse(r)
	})
}
