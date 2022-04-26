package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/shomali11/slacker"
)

func main() {
	os.Setenv("APP_TOKEN", "xapp-1-A03CEEVUUKW-3432788188689-e9202fccebac348698c302660d390da2998d98c3770c12eef85a81f055ddc5f7")
	os.Setenv("BOT_TOKEN", "xoxb-3412834228678-3432830473249-rXy4te77AImr6NlArjMsOOjb")

	bot := slacker.NewClient(os.Getenv("BOT_TOKEN"), os.Getenv("APP_TOKEN"))

	bot.Command("<YOB>", &slacker.CommandDefinition{
		Description: "Age calculator",
		Example:     "1996 OR 25",
		Handler: func(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
			enteredNumber := req.Param("YOB")
			ageOrYear, err := strconv.Atoi(enteredNumber)
			if err != nil {
				fmt.Print("Error Occured!")
				res.Reply("Some error occured!")
			} else if ageOrYear <= 2022 && ageOrYear >= 1900 {
				age := 2022 - ageOrYear
				response := fmt.Sprintf("Your age is %d", age)
				res.Reply(response)
			} else if ageOrYear <= 100 && ageOrYear >= 0 {
				year := 2022 - ageOrYear
				response := fmt.Sprintf("You born in %d", year)
				res.Reply(response)
			} else {
				res.Reply("Your entered wrong number")
			}

		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
