package teler

import (
	"fmt"
	"strings"

	"net/http"

	"github.com/kitabisa/teler-waf/request"
	"github.com/kitabisa/teler-waf/threat"
	"gitlab.com/golang-commonmark/mdurl"
)

// inThreatIndex checks if the given substring is in specific threat datasets
func (t *Teler) inThreatIndex(kind threat.Threat, substr string) bool {
	if i := strings.Index(t.threat.data[kind], substr); i >= 0 {
		return true
	}

	return false
}

// inWhitelist checks if the given substring is in whitelist patterns
func (t *Teler) inWhitelist(substr string) bool {
	substr = toURLDecode(substr)

	for _, pattern := range t.whitelistRegexes {
		if pattern.MatchString(substr) {
			return true
		}
	}

	return false
}

// headersToRawString converts a map of http.Header to
// multiline string, example:
// from,
//
//	Header = map[string][]string{
//		"Accept-Encoding": {"gzip, deflate"},
//		"Accept-Language": {"en-us"},
//		"Foo": {"Bar", "two"},
//	}
//
// to
//
//	Host: example.com
//	accept-encoding: gzip, deflate
//	Accept-Language: en-us
//	fOO: Bar
//	foo: two
func headersToRawString(headers http.Header) string {
	var h strings.Builder

	// Iterate over the request headers and append each key-value pair to the builder
	for key, values := range headers {
		for _, value := range values {
			h.WriteString(
				fmt.Sprintf("%s: %s\n", toURLDecode(key), toURLDecode(value)),
			)
		}
	}

	// Returns the accumulated string of builder
	return h.String()
}

// toURLDecode decode URL-decoded characters string using mdurl package
func toURLDecode(s string) string {
	return mdurl.Decode(s)
}

// isValidMethod check if the given request.Method is valid
func isValidMethod(method request.Method) bool {
	switch method {
	case request.GET, request.HEAD, request.POST, request.PUT, request.PATCH:
	case request.DELETE, request.CONNECT, request.OPTIONS, request.TRACE, request.ALL:
	case "":
		return true
	}

	return false
}

// normalizeRawStringReader trim double-quotes of HTTP raw string,
// replace double-escape of CR and LF, and double it in the end, and
// returning as pointer of strings.Reader
func normalizeRawStringReader(raw string) *strings.Reader {
	raw = strings.Trim(raw, `"`)
	raw = strings.ReplaceAll(raw, "\\n", "\n")
	raw = strings.ReplaceAll(raw, "\\r", "\r")
	raw += "\r\n\r\n"

	return strings.NewReader(raw)
}

// getClientIP to get client IP address from request
func getClientIP(r *http.Request) string {
	// Get the client's IP address from the X-Real-Ip header field
	clientIP := r.Header.Get("X-Real-Ip")

	// If the X-Real-Ip header field is not present, try the X-Forwarded-For header field
	if clientIP == "" {
		clientIP = r.Header.Get("X-Forwarded-For")
	}

	// If the X-Forwarded-For header field is not present, use the RemoteAddr field
	if clientIP == "" {
		clientIP = r.RemoteAddr
	}

	// Returning client IP address
	return clientIP
}
