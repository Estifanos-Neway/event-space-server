package commons

import "time"

var MinPasswordLength int = 6

// HTTP response messages
var InternalError string = "Internal_Error"
var No_Content string = "No_Content"

// jwt tokens
var AccessTokenExpiresAfter time.Duration = 20 * time.Minute
