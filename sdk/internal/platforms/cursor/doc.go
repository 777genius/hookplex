// Package cursor reserves the Cursor target identity in the internal platform tree.
//
// Cursor is currently scaffold/generate/import/validate only. It intentionally has
// no runtime event implementation in the SDK yet, but it still needs a distinct
// internal package so descriptor metadata does not alias existing platform internals.
package cursor
