// Copyright Dwi Siswanto and/or licensed to Dwi Siswanto under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
// See the LICENSE-ELv2 file in the project root for more information.

package request

import "net/http"

// Method is a type alias for string, used to represent HTTP methods.
//
// It is defined as a type alias to allow for custom methods to be added
// in the future, while still maintaining type-safety.
type Method string

// Constants representing common HTTP methods.
//
// These constants are of type Method and are assigned the values of the
// corresponding HTTP methods from the net/http package. Using these
// constants allows users of the request package to specify HTTP methods
// in a type-safe manner, rather than using raw strings.
const (
	GET       Method = http.MethodGet     // GET is the HTTP GET method.
	HEAD             = http.MethodHead    // HEAD is the HTTP HEAD method.
	POST             = http.MethodPost    // POST is the HTTP POST method.
	PUT              = http.MethodPut     // PUT is the HTTP PUT method.
	PATCH            = http.MethodPatch   // PATCH is the HTTP PATCH method.
	DELETE           = http.MethodDelete  // DELETE is the HTTP DELETE method.
	CONNECT          = http.MethodConnect // CONNECT is the HTTP CONNECT method.
	OPTIONS          = http.MethodOptions // OPTIONS is the HTTP OPTIONS method.
	TRACE            = http.MethodTrace   // TRACE is the HTTP TRACE method.
	ALL              = "ALL"              // ALL is representing as any HTTP method.
	UNDEFINED        = ""                 // UNDEFINED is undefined HTTP method.
)
