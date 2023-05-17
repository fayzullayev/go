package database

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

var err error

type Article struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Rate        string `json:"rate"`
}

func getEnvVariable(key string) string {
	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err.Error(), key)
	}

	return os.Getenv(key)
}

func NewPostgresSQLClient() {
	var (
		host     = getEnvVariable("DB_HOST")
		port     = getEnvVariable("DB_PORT")
		user     = getEnvVariable("DB_USER")
		dbname   = getEnvVariable("DB_NAME")
		password = getEnvVariable("DB_PASSWORD")
	)

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(Article{})
}

func CreateArticle(a *Article) (*Article, error) {

	//res := DB.Create(a)

	if res := DB.Create(a); res.Error != nil {
		return &Article{}, errors.New("article not created")
	}

	return a, nil
}

func ReadArticle(id string) (*Article, error) {
	var article Article

	if res := DB.First(&article, id).Error; res != nil {
		return nil, errors.New("article not found")
	}
	return &article, nil
}

func ReadArticles() ([]*Article, error) {

	var articles []*Article
	fmt.Println("ReadArticles : ")

	if res := DB.Find(&articles).Error; res != nil {
		return articles, errors.New("author not found")
	}

	return articles, nil

}

func UpdateArticle(article *Article) (*Article, error) {
	var updateArticle Article

	result := DB.Model(&updateArticle).Where(article.ID).Updates(article)

	if result.RowsAffected == 0 {
		return &Article{}, errors.New("artcile not updated")
	}

	return &updateArticle, nil
}

func DeleteArticle(id string) error {
	var deleteArticle Article

	result := DB.Where(id).Delete(&deleteArticle)

	if result.RowsAffected == 0 {
		return errors.New("article data not deleted")
	}

	return nil
}
