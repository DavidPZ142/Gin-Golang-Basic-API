package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type player struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Team         string `json:"team"`
	JerseyNumber int    `json:"jerseyNumber"`
}

var players = []player{
	{ID: "1", Name: "Lionel Messi", Team: "Inter Miami", JerseyNumber: 30},
	{ID: "2", Name: "Julian Alvarez", Team: "Manchester City", JerseyNumber: 19},
	{ID: "3", Name: "Luis Diaz", Team: "Liverpool", JerseyNumber: 7},
	{ID: "4", Name: "Luis Suarez", Team: "Gremio", JerseyNumber: 9},
}

func getPlayers(get *gin.Context) {
	get.IndentedJSON(http.StatusOK, players)
}

func getPlayerById(get *gin.Context) {
	id := get.Param("id")
	for _, player := range players {
		if player.ID == id {
			get.IndentedJSON(http.StatusOK, player)
			return
		}
	}
	get.IndentedJSON(http.StatusNotFound, gin.H{"message": "Jugador no existente"})
}

func savePlayer(post *gin.Context) {
	var newPlayer player
	err := post.BindJSON(&newPlayer)
	if err != nil {
		return
	}
	players = append(players, newPlayer)
	post.IndentedJSON(http.StatusCreated, newPlayer)
}

func main() {
	router := gin.Default()
	router.GET("/players", getPlayers)
	router.GET("/player/:id", getPlayerById)
	router.POST("/players", savePlayer)
	router.Run(":8080")
}
