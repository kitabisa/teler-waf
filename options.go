package teler

import "github.com/kitabisa/teler-waf/threat"

// Options is a struct for specifying configuration options for the teler.Teler middleware.
type Options struct {
	// Excludes is a list of threat types to exclude from the security checks.
	// These threat types are defined in the threat.Threat type.
	Excludes []threat.Threat

	// Whitelists is a list of regular expressions that match request elements
	// that should be excluded from the security checks. The request elements
	// that can be matched are user-agent, request path, HTTP referrer,
	// IP address, and request query values.
	Whitelists []string

	// Customs is a list of custom security rules to apply to incoming requests.
	// These rules can be used to create custom security checks or to override
	// the default security checks provided by teler-waf.
	Customs []Rule

	// LogFile is the file path for the log file to store the security logs.
	// If LogFile is specified, log messages will be written to the specified
	// file in addition to stderr (if NoStderr is false).
	LogFile string

	// TODO:
	// LogRotate specifies whether to rotate the log file when it reaches a new day.
	// LogRotate bool

	// NoStderr is a boolean flag indicating whether or not to suppress log messages
	// from being printed to the standard error (stderr) stream. When set to true, log messages
	// will not be printed to stderr. If set to false, log messages will be printed to stderr.
	// By default, log messages are printed to stderr (false).
	NoStderr bool
}
