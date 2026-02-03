package handler

import (
	app "swords_to_poll_shares/internal"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type ReadyHandler struct {
	logger *zap.Logger
}

func NewReadyHandler(logger *zap.Logger) *ReadyHandler {
	return &ReadyHandler{
		logger: logger,
	}
}

func (h *ReadyHandler) Handle(s *discordgo.Session, r *discordgo.Ready) {
	h.logger.Info("Bot Logged in", zap.String("user", r.User.String()))

	for _, guild := range r.Guilds {
		err := s.RequestGuildMembers(guild.ID, "", 0, "false", false)
		if err != nil {
			h.logger.Error("Failed to Request Guild Members", zap.Error(err))
		}
	}

	_, err := s.ChannelMessageSendComplex("1461914574663061574", &discordgo.MessageSend{
		Poll: app.CreateWeeklyCommanderPoll(),
	})

	if err != nil {
		h.logger.Error("Failed to send weekly poll message", zap.Error(err))
	}
}

func (h *ReadyHandler) GetHandlerFunc() interface{} {
	return h.Handle
}
