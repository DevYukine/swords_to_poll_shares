package commands

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Command interface {
	GetName() string
	GetDescription() string
	Execute(s *discordgo.Session, i *discordgo.InteractionCreate) error
}

func RegisterCommands(session *discordgo.Session, logger *zap.Logger, commands []Command) error {
	foundCommands := make([]*discordgo.ApplicationCommand, 0, len(commands))

	for _, cmd := range commands {
		command := &discordgo.ApplicationCommand{
			Name:        cmd.GetName(),
			Description: cmd.GetDescription(),
		}
		foundCommands = append(foundCommands, command)
	}

	_, err := session.ApplicationCommandBulkOverwrite(session.State.User.ID, "", foundCommands)
	if err != nil {
		return err
	}

	logger.Info("Registered commands", zap.Int("count", len(foundCommands)))

	return nil
}
