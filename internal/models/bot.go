package models

const (
	ConfigPath = "./configs/configs.json"
	HelpText   = `
List of available commands:
1. !help - Displays a list of all available commands and their descriptions.
2. !weather <city> - Get current weather information for the specified city.
	Usage example:
***	 !weather Almaty ***
3. !remindme <time> - Set a reminder for yourself at a specific time. 
The <time> should be in the format "15:04" (24-hour clock). 
	Usage example: 
***	 !remindme go to football 18:30 ***
	`
	WeatherHelp = `
!weather <city> - Get current weather information for the specified city.
	Usage example:
	 !weather Almaty
`
	ReminderHelp = `
!remindme <reason for reminding> <time> - Set a reminder for yourself at a specific time. 
The <time> should be in the format "15:04" (24-hour clock). 
	Usage example: 
***	 !remindme go to football 18:30 ***
`
)
