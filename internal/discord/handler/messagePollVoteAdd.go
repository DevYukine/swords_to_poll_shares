package handler

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type MessagePollVoteAddHandler struct {
	logger *zap.Logger
}

func NewMessagePollVoteAddHandler(logger *zap.Logger) *MessagePollVoteAddHandler {
	return &MessagePollVoteAddHandler{
		logger: logger,
	}
}

func (h *MessagePollVoteAddHandler) Handle(s *discordgo.Session, r *discordgo.MessagePollVoteAdd) {
	member, err := s.State.Member(r.GuildID, r.UserID)

	if err != nil {
		h.logger.Error("Failed to get guild member", zap.Error(err))
	}

	h.logger.Debug("Received MessagePollVoteAdd Event", zap.Any("member", member))
}

func (h *MessagePollVoteAddHandler) GetHandlerFunc() interface{} {
	return h.Handle
}
