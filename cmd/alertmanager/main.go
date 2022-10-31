package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		d, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		p := make(map[string]interface{})
		err = json.Unmarshal(d, &p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		b, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{})
			return
		}
		fmt.Println(string(b))
		c.JSON(http.StatusOK, gin.H{})
	})
	r.Run(":8080")
}
