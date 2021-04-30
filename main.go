package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

)

type EncryptRequest struct {
	Value string `form:"value" json:"value" xml:"value"  binding:"required"`
	Key string `form:"key" json:"key" xml:"key" `
}

type DecryptRequest struct {
	Value string `form:"value" json:"value" xml:"value"  binding:"required"`
	Key string `form:"key" json:"key" xml:"key" `
}

func main() {

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())


  v1 := router.Group("/v1")
  {
   v1.POST("/decrypt", CBCDecryptHandler)
	 v1.POST("/encrypt", CBCEncryptHandler)

  }
  router.Run()
}

func CBCDecryptHandler(context *gin.Context) {
	var data DecryptRequest
	if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	value := data.Value
	key := data.Key

	decrypt, err := Decrypt(value, key)

	if err != nil {
		context.JSON(500, gin.H{"error": err})
		return
	}
	

	context.JSON(200, gin.H{"value": string(decrypt)})
}

func CBCEncryptHandler(context *gin.Context) {
	var data EncryptRequest
	if err := context.ShouldBindJSON(&data); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	value := data.Value
	key := data.Key


	encrypt, err := Encrypt([]byte(value), key)

	if err != nil {
		context.JSON(500, gin.H{"error": err})
		return
	}

	context.JSON(200, gin.H{"value": encrypt})
}