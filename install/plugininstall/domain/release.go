package domain

// Release is a GitHub release (subset of API).
type Release struct {
	TagName    string
	Draft      bool
	Prerelease bool
	Assets     []Asset
}

// Asset is a release attachment.
type Asset struct {
	Name               string
	BrowserDownloadURL string
	Size               int64
}
