package commons

import "time"

var MinPasswordLength int = 6

// HTTP response messages
const InternalError string = "Internal_Error"
const No_Content string = "No_Content"
const Ok string = "OK"

// jwt tokens
const AccessTokenExpiresAfter time.Duration = 20 * time.Minute

const Invalid_Token string = "Invalid_Token"
