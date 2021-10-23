# Beatrix
Beatrix is a go binding to a bot, used by meanOs team to get notifications from services

Simple go package

Usage:
```golang
import "github.com/meanOs/Beatrix"

func main(){
  beatrix.Init(issuer, token, channelID)
  
  go beatrix.Message("Test message")
  go beatrix.SendError("Test error")
```
