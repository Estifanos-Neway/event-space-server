package commons

import "time"

var MinPasswordLength int = 6

// HTTP response messages
const No_Content = "No_Content"
const Ok = "OK"
const Invalid_Input = "Invalid_Input"
const Not_Found = "Not_Found"

// jwt tokens
const AccessTokenExpiresAfter time.Duration = 20 * time.Minute
const RefreshTokenExpiresAfter time.Duration = 1440 * time.Hour
