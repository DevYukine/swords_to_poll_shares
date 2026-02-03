package handler

import (
	"github.com/DevYukine/swords_to_poll_shares/internal/discord/commands"
	"github.com/bwmarrin/discordgo"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type InteractionCreateHandler struct {
	logger   *zap.Logger
	commands []commands.Command
}

func NewInteractionCreateHandler(logger *zap.Logger, params struct {
	fx.In
	Commands []commands.Command `group:"commands"`
}) *InteractionCreateHandler {

	return &InteractionCreateHandler{
		logger:   logger,
		commands: params.Commands,
	}
}

func (h *InteractionCreateHandler) Handle(s *discordgo.Session, r *discordgo.InteractionCreate) {
	if r.Type != discordgo.InteractionApplicationCommand {
		return
	}

	data := r.ApplicationCommandData()

	for _, cmd := range h.commands {
		if data.Name == cmd.GetName() {
			err := cmd.Execute(s, r)

			if err != nil {
				h.logger.Error("Failed to execute command", zap.Error(err))
			}

			return
		}
	}

	h.logger.Warn("Received unknown command", zap.String("command_name", data.Name))
}

func (h *InteractionCreateHandler) GetHandlerFunc() interface{} {
	return h.Handle
}
