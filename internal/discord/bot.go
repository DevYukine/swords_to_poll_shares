package discord

import (
	app "swords_to_poll_shares/internal"

	"github.com/bwmarrin/discordgo"
)

func ProvideDiscordBotSession(config *app.Config) *discordgo.Session {
	session, _ := discordgo.New("Bot " + config.DiscordBotToken)

	session.Identify.Intents |= discordgo.IntentGuildMessagePolls |
		discordgo.IntentGuildMembers |
		discordgo.IntentGuilds

	session.LogLevel = discordgo.LogDebug

	return session
}
