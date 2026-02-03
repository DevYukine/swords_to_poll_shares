package main

import (
	"context"

	app "github.com/DevYukine/swords_to_poll_shares/internal"
	"github.com/DevYukine/swords_to_poll_shares/internal/discord"
	"github.com/DevYukine/swords_to_poll_shares/internal/discord/commands"
	"github.com/DevYukine/swords_to_poll_shares/internal/discord/handler"

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
			AsHandler(handler.NewInteractionCreateHandler),
			AsCommand(commands.NewPingCommand),
		),
		fx.Invoke(func(session *discordgo.Session, params struct {
			fx.In
			Handlers []handler.Handler `group:"handlers"`
		}) {
			for _, h := range params.Handlers {
				session.AddHandler(h.GetHandlerFunc())
			}
		}),
		fx.Invoke(func(lc fx.Lifecycle, cfg *app.Config, logger *zap.Logger, discordSession *discordgo.Session, params struct {
			fx.In
			Commands []commands.Command `group:"commands"`
		}) {
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					logger.Info("Connecting to discord websocket")

					err := discordSession.Open()
					if err != nil {
						return err
					}

					logger.Info("Connected to discord websocket")

					err = commands.RegisterCommands(discordSession, logger, params.Commands)
					if err != nil {
						logger.Error("Failed to register commands", zap.Error(err))
						return err
					}

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

func AsCommand(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(commands.Command)),
		fx.ResultTags(`group:"commands"`),
	)
}
