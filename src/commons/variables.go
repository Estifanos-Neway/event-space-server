package commons

import "time"

var MinPasswordLength int = 6

// HTTP response messages
const NoContent = "No_Content"
const Ok = "OK"
const InvalidInput = "Invalid_Input"
const NotFound = "Not_Found"
const InternalError = "Internal_Error"

// jwt tokens
const AccessTokenExpiresAfter time.Duration = 45 * time.Minute
const RefreshTokenExpiresAfter time.Duration = 1440 * time.Hour
const EmailVerificationExpiresAfter time.Duration = 30 * time.Minute
