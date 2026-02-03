package app

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

func CreateWeeklyCommanderPoll() *discordgo.Poll {
	now := time.Now()

	var answers []discordgo.PollAnswer
	germanWeekdays := map[time.Weekday]string{
		time.Monday:    "Montag",
		time.Tuesday:   "Dienstag",
		time.Wednesday: "Mittwoch",
		time.Thursday:  "Donnerstag",
		time.Friday:    "Freitag",
		time.Saturday:  "Samstag",
		time.Sunday:    "Sonntag",
	}

	for i := 0; i < 7; i++ {
		day := now.AddDate(0, 0, i)
		weekdayGerman := germanWeekdays[day.Weekday()]
		dateStr := day.Format("02.01.2006")

		answers = append(answers, discordgo.PollAnswer{
			Media: &discordgo.PollMedia{
				Text: weekdayGerman + " " + dateStr,
			},
		})
	}

	firstDay := now.Format("02.01")
	lastDay := now.AddDate(0, 0, 6).Format("02.01")

	poll := &discordgo.Poll{
		Question: discordgo.PollMedia{
			Text: "Commander Termine in der Woche vom " + firstDay + " - " + lastDay,
		},
		Answers:          answers,
		Duration:         160,
		AllowMultiselect: true,
	}

	return poll
}
