package models

const (
	ConfigPath = "./configs/configs.json"
	HelpText   = `
List of available commands:
1. !help - Displays a list of all available commands and their descriptions.
2. !weather <city> - Get current weather information for the specified city.
3. !language <target_language> <message> - Translate the message into the specified language.

Unique features:
- Weather: The !weather command allows you to check the current weather in the specified location.
	Usage example:
	 !weather Moscow

- Language Translation: The !language command translates your message into the specified language.
	Usage example:
	 !language es Hello, how are you?   or
	 !language english Hello, how are you? (Translation to English)
	`
	WeatherHelp = `
!weather <city> - Get current weather information for the specified city.
	Usage example:
	 !weather Moscow
	`
)
