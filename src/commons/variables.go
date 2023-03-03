package commons

import "time"

var MinPasswordLength int = 6

// HTTP response messages
const No_Content = "No_Content"
const Ok = "OK"
const Invalid_Input = "Invalid_Input"

// jwt tokens
const AccessTokenExpiresAfter time.Duration = 20 * time.Minute
