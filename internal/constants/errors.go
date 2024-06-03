package constants

// Discord Errors
const (
	DiscordUnexpectedHandler = "ERROR An unexpected error occured. Try again later."
	ErrorDiscordMessage      = "ERROR unable to send message to channel."
	ErrorDiscordMessageSend  = "arf! Something unexpected happened. Give me treats for now you 'lil shit."
)

// Dynamo DB Errors
const (
	ErrorBuildExpression = "ERROR unable to build new scan expression: %v\n"
	ErrorScan            = "ERROR failed to search manhwa via Scan: %v\n"
	ErrorUnmarshalItem   = "ERROR failed to unmarshal item: %v\n"
)
