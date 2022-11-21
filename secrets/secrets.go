package secrets

import "os"

// Session exports a variable containing a session token for AoC
var Session = "session=" + os.Getenv("ADVENT_OF_CODE_SESSION_TOKEN")