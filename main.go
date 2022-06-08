package main

import(
	"log"
	"time"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main(){
	b, err := tb.NewBot(tb.Settings{
		Token: "YOUR-API-TOKEN",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})


	if err != nil {
		log.Fatal(err)
		return
	}


	b.Handle("/sayhello", func(m *tb.Message){
		b.Send(m.Sender, "Hi Guyz")
	})

	b.Start()
}