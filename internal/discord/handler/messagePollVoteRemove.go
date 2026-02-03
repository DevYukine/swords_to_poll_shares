package handler

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type MessagePollVoteRemoveHandler struct {
	logger *zap.Logger
}

func NewMessagePollVoteRemoveHandler(logger *zap.Logger) *MessagePollVoteRemoveHandler {
	return &MessagePollVoteRemoveHandler{
		logger: logger,
	}
}

func (h *MessagePollVoteRemoveHandler) Handle(s *discordgo.Session, r *discordgo.MessagePollVoteRemove) {
	member, err := s.State.Member(r.GuildID, r.UserID)

	if err != nil {
		h.logger.Error("Failed to get guild member", zap.Error(err))
	}

	h.logger.Debug("Received MessagePollVoteRemove Event", zap.Any("member", member))
}

func (h *MessagePollVoteRemoveHandler) GetHandlerFunc() interface{} {
	return h.Handle
}
