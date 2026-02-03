package discord

import (
	"fmt"
	app "swords_to_poll_shares/internal"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

func ProvideDiscordBotSession(config *app.Config, logger *zap.Logger) *discordgo.Session {
	overrideDiscordGoLogger(logger)

	session, _ := discordgo.New("Bot " + config.DiscordBotToken)

	session.Identify.Intents |= discordgo.IntentGuildMessagePolls |
		discordgo.IntentGuildMembers |
		discordgo.IntentGuilds

	session.LogLevel = discordgo.LogDebug

	return session
}

func overrideDiscordGoLogger(logger *zap.Logger) {
	discordgo.Logger = func(level, caller int, format string, args ...interface{}) {
		msg := fmt.Sprintf(format, args...)
		log := logger.WithOptions(
			zap.AddCallerSkip(caller),
			zap.AddStacktrace(zap.ErrorLevel),
		)

		switch level {
		case discordgo.LogDebug:
			log.Debug(msg)
		case discordgo.LogInformational:
			// discordgo informational contains debug logs so we map it to debug
			log.Debug(msg)
		case discordgo.LogWarning:
			log.Warn(msg)
		case discordgo.LogError:
			log.Error(msg)
		default:
			log.Error("Unknown log level", zap.Int("level", level), zap.String("message", msg))
		}
	}
}
