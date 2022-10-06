package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	token := flag.String("token", "", "steam api token")
	steamid := flag.String("steamid", "", "steamid")
	flag.Parse()
	fmt.Println(*token)
	//response, err := http.Get("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" + *token + "&steamids=" + *steamid)
	response, err := http.Get("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=" + *token + "&steamid=" + *steamid + "&format=json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	bodyBytes, err := io.ReadAll(response.Body)
	fmt.Println(string(bodyBytes))

}
