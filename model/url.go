package model

import (
	"fmt"
	"log"

	"github.com/panda8z/shorturl/errors"
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Origin string `json:"origin" gorm:"type:varchar(255);not null"`
	Short  string `json:"short" gorm:"type:varchar(64);not null"`
}

func (u *Url) String() string {
	return fmt.Sprintf("\n{\n origin:%s,\n short:%s\n\n}", u.Origin, u.Short)
}

// CheckUrlOrigin  check origin is existed
func CheckUrlOrigin(origin string) (*Url, int) {
	var url Url
	log.Println("long url is :", origin)
	db.Where("origin = ?", origin).First(&url)
	log.Println(url.String())
	if url.ID > 0 {
		return &url, errors.Exist // 2000
	} else {
		return &url, errors.SUCCESS
	}
}

// CheckUrlShort check short is existed
func CheckUrlShort(short string) (*Url, int) {
	var url Url
	db.Where("short = ?", short).First(&url)
	if url.ID > 0 { // TODO 判断逻辑需稍微改动
		return &url, errors.Exist // 2000
	}
	return nil, errors.UnExist
}

// GetUrlByID search url with specified id
func GetUrlByID(id int) (*Url, error) {
	var url Url
	err := db.Where("id = ?", id).First(&url).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &url, nil
}

// GetUrlByShort search url with specified id
func GetUrlByShort(short string) (*Url, error) {
	var url Url
	err := db.Where("short = ?", short).First(&url).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &url, nil
}

// GetUrlByOrigin search url with specified id
func GetUrlByOrigin(origin string) (*Url, error) {
	var url Url
	err := db.Where("origin = ?", origin).First(&url).Error
	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	return &url, nil
}

// CreateUrl add a new url to database
func CreateUrl(url *Url) (*Url, error) {
	err := db.Create(&url).Error
	if err != nil {
		return nil, err
	}
	return url, nil
}

// UrlList get url list in pageable
func UrlList(pageSize int, pageNum int) ([]Url, int) {
	var cates []Url
	var total int

	db.Model(&cates).Count(&total)
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return cates, total
}

// SoftDeletUrl delete url softy
func SoftDeletUrl(id int) int {
	err := db.Where("id = ?", id).Delete(&Url{}).Error
	if err != nil {
		return errors.ERROR
	}
	return errors.SUCCESS
}

func UpdateUrlByID(id uint, url *Url) error {
	uMap := map[string]interface{}{
		"Short": url.Short,
	}
	err := db.Model(&url).Where("id = ?", id).Updates(uMap).Error
	if err != nil {
		return err
	}
	return nil
}
