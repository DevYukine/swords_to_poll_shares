package main

import (
	"context"
	"fmt"
	app "swords_to_poll_shares/internal"
	"swords_to_poll_shares/internal/discord"
	"swords_to_poll_shares/internal/discord/handler"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	appFx := fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			logger := fxevent.ZapLogger{Logger: log}

			logger.UseLogLevel(zap.DebugLevel)

			return &logger
		}),
		fx.Provide(
			app.ProvideLogger,
			app.ProvideConfig,
			app.ProvideHTTPClient,
			discord.ProvideDiscordBotSession,
			AsHandler(handler.NewReadyHandler),
			AsHandler(handler.NewMessagePollVoteAddHandler),
			AsHandler(handler.NewMessagePollVoteRemoveHandler),
		),
		fx.Invoke(func(session *discordgo.Session, logger *zap.Logger, params struct {
			fx.In
			Handlers []handler.Handler `group:"handlers"`
		}) {
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

			for _, h := range params.Handlers {
				session.AddHandler(h.GetHandlerFunc())
			}
		}),
		fx.Invoke(func(lc fx.Lifecycle, cfg *app.Config, logger *zap.Logger, discordSession *discordgo.Session) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					logger.Info("Connecting to discord websocket")

					err := discordSession.Open()
					if err != nil {
						return err
					}

					logger.Info("Connected to discord websocket")

					return nil
				},
				OnStop: func(ctx context.Context) error {
					logger.Info("Disconnecting from discord websocket")

					err := discordSession.Close()
					if err != nil {
						return err
					}

					logger.Info("Disconnected from discord websocket")

					return nil
				},
			})
		}))

	appFx.Run()
}

func AsHandler(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(handler.Handler)),
		fx.ResultTags(`group:"handlers"`),
	)
}
