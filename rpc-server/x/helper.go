package x

import (
	"encoding/base64"
	"fmt"

	"github.com/panda8z/shorturl/model"
)

func CreateUrl(long string) *model.Url {
	url := &model.Url{}
	url.Origin = long
	model.CreateUrl(url)
	url.Short = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("surl:%d", url.ID)))
	err := model.UpdateUrlByID(url.ID, url)
	if err != nil {
		return nil
	}
	url, err = model.GetUrlByID(int(url.ID))
	if err != nil {
		return nil
	}
	return url
}
