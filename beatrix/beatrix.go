package beatrix

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"sync"
)

type DS struct {
	Mutex   sync.Mutex
	Discord *discordgo.Session
}

var (
	Token     string
	Issuer    string
	ChannelID string
	ErrorMode bool

	Discord DS
)

// Function to fire up discord bot, issuer will be in the heading of message, token is bot token and channelId is channelID
func Init(issuer, token, channelID string) {
	Token = token
	Issuer = issuer
	ChannelID = channelID
	ErrorMode = false
	var err error
	Discord.Mutex.Lock()
	Discord.Discord, err = discordgo.New("Bot " + Token)
	if err != nil {
		// Failed to init Beatrix
		log.Panic(err)
	}
	err = Discord.Discord.Open()
	if err != nil {
		// Failed to init Beatrix
		log.Panic(err)
	}
	_, err = Discord.Discord.ChannelMessageSend(ChannelID, "["+Issuer+" / INIT]")
	if err != nil {
		log.Panic(err)
	}
	Discord.Mutex.Unlock()
}

func Reinit() {
	var err error
	Discord.Mutex.Lock()
	Discord.Discord, err = discordgo.New("Bot " + Token)
	if err != nil {
		// Failed to init Beatrix
		ErrorMode = true
		log.Println(err)
		return
	}
	err = Discord.Discord.Open()
	Discord.Mutex.Unlock()
	if err != nil {
		// Failed to init Beatrix
		ErrorMode = true
		log.Println(err)
		return
	}
	ErrorMode = false
	return
}

// Simply send a message to main channel
func Message(message string) {
	message = "[" + Issuer + "]\n" + message
	if ErrorMode {
		log.Println(message)
		Reinit()
		return
	}
	Discord.Mutex.Lock()
	_, err := Discord.Discord.ChannelMessageSend(ChannelID, message)
	Discord.Mutex.Unlock()
	if err != nil {
		// Since we have goroutine, we don't have to return or something
		// Better re-init discord
		log.Println(err)
		Reinit()
	}
}

// Send error message to channel
func SendError(message, localIssuer string) {
	message = "[" + Issuer + " | " + localIssuer + "]\n" + message
	if ErrorMode {
		log.Println(message)
		Reinit()
		return
	}
	Discord.Mutex.Lock()
	_, err := Discord.Discord.ChannelMessageSend(ChannelID, message)
	Discord.Mutex.Unlock()
	if err != nil {
		log.Println(err)
		Reinit()
	}
	return
}

func Panic(message string) {
	m := "[" + Issuer + " / PANIC]\n@everyone\n\n" + message
	if ErrorMode {
		log.Println(message)
		Reinit()
		return
	}
	Discord.Mutex.Lock()
	_, err := Discord.Discord.ChannelMessageSend(ChannelID, m)
	Discord.Mutex.Unlock()
	if err != nil {
		log.Println(err)
		Reinit()
	}
	return
}
