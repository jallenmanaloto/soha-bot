package constants

// Commands
const (
	Command = "command"
	Bury    = "bury"
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
	MessageAlert            = "**Note:** All manhwa that is in your watch list will be guarded by Soha. Meaning, for any new chapters he find, you will be alerted on this server."
	MessageBury             = "`!soha bury <ID>:` Soha will bury the manhwa and permanently remove from your watch list."
	MessageDefault          = "I don't know that trick! Just give me treats, you piece of shit!"
	MessageDeleteSuccess    = "I successfully bury this title. arf! Give me treats, am good boi!"
	MessageFetch            = "`!soha fetch:` Soha will fetch all the titles he is watching for you."
	MessageHello            = "arf arf! What is it, you degenerate!"
	MessageLook             = "`!soha look <title>:` Soha will look for a manhwa with the provided title."
	MessageManhwaNotExist   = "I can't find that manhwa. You're wasting my energy you shit!"
	MessageEmptyWatchList   = "arf! arf! arf! I am not watching anything yet. Give me treats and I'll watch anything for you dumdum!"
	MessageTricks           = "Soha's tricks or command displays things he can do.\nYou can call out to Soha with `!soha` followed by your command\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s"
	MessageTricksEmbedTitle = "**Soha's tricks and quirks**"
	MessageWatch            = "`!soha watch <ID>:` Soha will watch out for new chapters for the title."
	MessageWatchListTitle   = "arf! arf! These are the titles I am watchin' for you shit."
	MessageWatchSuccess     = "arf! I am now watching that for you, you 'lil shit."
)

// Message embed
const (
	EmbedManhwaDesc      = "ID: %s\n%s"
	EmbedManhwaWatchList = "**%s**ID: %s\n%s\n%s\n\n" // format: Title, ID, Chapter, Url
)
