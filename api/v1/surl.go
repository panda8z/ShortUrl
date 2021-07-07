package v1

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/panda8z/shorturl/errors"
	"github.com/panda8z/shorturl/model"
)

type LongUrl struct {
}

func QuaryOriginUrl(c *gin.Context) {
	short := c.Params.ByName("short")
	// 1. quary to check.
	url, errCode := model.CheckUrlShort(short)
	if errCode == errors.Exist {
		c.JSON(http.StatusOK, gin.H{"error": "", "url": url.Origin})
	} else if errCode == errors.UnExist {
		c.JSON(http.StatusOK, gin.H{"error": errors.CodeMap[errCode], "url": ""})
	}
}

func GenShorturl(c *gin.Context) {
	var long struct {
		Url string `json:"url" binding:"required"`
	}
	if c.Bind(&long) == nil {
		// 1. quary is exist
		// 2. create short url
		// 3. exist return false
		// 4. unexist save Url Model return Url.Short
		log.Println(long.Url)
		url, errCode := model.CheckUrlOrigin(long.Url)

		if errCode == errors.SUCCESS {
			url := createUrl(long.Url)
			c.JSON(http.StatusOK, gin.H{
				"error": "",
				"surl":  url.Short,
			})
		}

		if errCode == errors.Exist && url != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "this url already exist",
				"surl":  url.Short,
			})
		}

	}
}

func createUrl(long string) *model.Url {
	url := &model.Url{Origin: long}
	model.CreateUrl(url)
	url.Short = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("surl:%d", url.ID)))
	err := model.UpdateUrlByID(url.ID, url)
	if err != nil {
		return nil
	}
	return url
}
