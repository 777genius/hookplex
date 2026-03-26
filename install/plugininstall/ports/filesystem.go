package ports

import (
	"context"
	"io"
)

// FileSystem writes installed binaries and inspects install paths.
type FileSystem interface {
	MkdirAll(ctx context.Context, path string, perm uint32) error
	WriteFileAtomic(ctx context.Context, dir, name string, r io.Reader, size int64, perm uint32) error
	// PathExists is true if path exists; false with nil error if absent; other errors (e.g. permission) are returned.
	PathExists(ctx context.Context, path string) (bool, error)
	// RemoveBestEffort removes path if present; ignores not-exist (best effort before overwrite on Windows).
	RemoveBestEffort(ctx context.Context, path string) error
}
