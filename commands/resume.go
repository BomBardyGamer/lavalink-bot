package commands

import (
	"context"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

func (c *Commands) Resume(e *handler.CommandEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	player := c.Lavalink.ExistingPlayer(*e.GuildID())

	if err := player.Update(ctx, lavalink.WithPaused(false)); err != nil {
		return e.CreateMessage(discord.MessageCreate{
			Content: "Failed to resume player",
		})
	}

	return e.CreateMessage(discord.MessageCreate{
		Content: "▶ Resumed player",
	})
}
