package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type response struct {
  ID        string  `json:"id"`
  Prompt    string  `json:"prompt"`
  Response  string  `json:"response"`
}

type npc struct {
  Name      string     `json:"name"`
  Responses []response `json:"responses"`
}

var npcs = []npc {
  {Name: "TavernKeep", Responses: []response {
    {ID: "0", Prompt: "bathroom", Response: "Oh its around the corner and on the left"},
    {ID: "1", Prompt: "drinks", Response: "We got a fine selection of teas including Lady Grey and Chamomile"},
    {ID: "2", Prompt: "food", Response: "We're well known for our grilled cheese"},
    },
  },
}

func getNPCs(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, npcs)
}

func main() {
  router := gin.Default()
  router.GET("/npcs", getNPCs)

  router.Run("localhost:8080")
}
