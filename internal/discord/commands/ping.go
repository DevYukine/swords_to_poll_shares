package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type PingCommand struct {
	logger *zap.Logger
}

func NewPingCommand(logger *zap.Logger) *PingCommand {
	return &PingCommand{logger: logger}
}

func (c PingCommand) GetName() string {
	return "ping"
}

func (c PingCommand) GetDescription() string {
	return "Replies with Pong and Latencies of the Bot!"
}

func (c PingCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	wsPing := s.HeartbeatLatency().Milliseconds()

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong! Getting rest latency...",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		return err
	}

	sent, err := s.InteractionResponse(i.Interaction)
	if err != nil {
		return err
	}

	restPing := time.Since(sent.Timestamp).Milliseconds()

	editedContent := fmt.Sprintf("Pong!\nWebSocket Latency: %dms\nREST Latency: %dms", wsPing, restPing)

	_, err = s.InteractionResponseEdit(i.Interaction, &discordgo.WebhookEdit{
		Content: &editedContent,
	})

	return err
}
