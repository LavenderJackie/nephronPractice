package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

type response struct {
  ID      string  `json:"id"`
  Prompt  string  `json:"prompt"`
  Phrase  string  `json:"phrase"`
}
type npc struct {
  Name      string      `json:"name"`
  Responses []response  `json:"responses"`
}

var npcs = []npc {
  {Name: "Tavern Keep", Responses: []response{
    {ID: "0", Prompt: "", Phrase: "im sorry i dont understand, could you rephrase??"},
    {ID: "1", Prompt: "bathroom", Phrase: "oh its down the hall to the right of the kitchen"},
    {ID: "2", Prompt: "eat", Phrase: "we dont have a whole lot to eat but we have a damn good grilled cheese"},
    {ID: "3", Prompt: "drink", Phrase: "we have some local mead left over from the harvest"},
    }},
  {Name: "Drunk Man", Responses: []response{
    {ID: "0", Prompt: "", Phrase: "fuck off unless youre bringing another round"},
    }},
}

func postNPC(c *gin.Context) {
  var newNPC npc

  if err := c.BindJSON(&newAlbum); err != nil {
    return
  }

  npcs = append(npcs, newNPC)
  c.IndentedJSON(http.StatusCreated, newNPC)
}

func getNPCs(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, npcs)
}

func getNPCByName(c *gin.Context) {
  name := c.Param("name")

  for _, a := range npcs {
    if a.Name == name {
      c.IndentedJSON(http.StatusOK, a)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "npc not found"})
}

func main() {
  router := gin.Default()
  router.GET("/npcs", getNPCs)
  router.GET("/npcs/:name", getNPCByName)
  router.POST("/npcs", postNPC)

  router.Run("localhost:8080")
}
