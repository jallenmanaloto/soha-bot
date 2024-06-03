package constants

// Commands
const (
	Command = "command"
	Default = "default"
	Fetch   = "fetch"
	Hello   = "hello"
	Look    = "look"
	Prefix  = "!soha"
	Tricks  = "tricks"
	Watch   = "watch"
)

// Command messages
const (
	MessageDefault          = "I don't know that trick! Just give me treats, you piece of shit!"
	MessageFetch            = "`!soha fetch:` Soha will fetch all the titles he is watching for you."
	MessageHello            = "arf arf! What is it, you degenerate!"
	MessageLook             = "`!soha look <title>:` Soha to look for a manhwa with that title."
	MessageManhwaNotExist   = "I can't find that manhwa. You're wasting my energy you shit!"
	MessageTricks           = "Soha's tricks or command displays things he can do.\nYou can call out to Soha with `!soha` followed by your command\n\n%s\n\n%s\n\n%s"
	MessageTricksEmbedTitle = "**Soha's tricks and quirks**"
	MessageWatch            = "`!soha watch <ID>:` Soha will watch out for new chapters for the title."
	MessageWatchSuccess     = "arf! I am now watching that for you, you 'lil shit."
)

// Message embed
const (
	EmbedManhwaDesc = "ID: %s\n%s"
)