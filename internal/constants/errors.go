package constants

// Discord Errors
const (
	DiscordUnexpectedHandler = "ERROR An unexpected error occured. Try again later."
	ErrorDiscordMessage      = "ERROR unable to send message to channel."
	ErrorDiscordMessageSend  = "arf! arf! Something unexpected happened. Give me treats for now you 'lil shit."
	WatchAlreadyExist        = "arf! arf! I am already watching that title, dummy!"
)

// Dynamo DB Errors
const (
	ErrorBuildExpression = "ERROR unable to build new scan expression: %v\n"
	ErrorMarshalItem     = "ERROR failed to marshal item: %v\n"
	ErrorScan            = "ERROR failed to search manhwa via Scan: %v\n"
	ErrorUnmarshalItem   = "ERROR failed to unmarshal item: %v\n"
	ErrorPutItem         = "ERROR failed to put item: %v\n"
	ErrorQuery           = "ERROR failed to query items: %v\n"
)

// Utils package Errors
const (
	ErrorGenerateId = "ERROR unable to generate a random ID: %v\n"
)

// HTTP Errors
const (
	ErrorJsonDecode = "ERROR failed to decode json body request: %v\n"
	ErrorMarshalRes = "ERROR failed to marshal http response: %v\n"
	Unauthorized    = "Unauthorized attempt. We cannot verify the signature key provided."
)
