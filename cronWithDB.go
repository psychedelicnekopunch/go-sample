package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/psychedelicnekopunch/go-sample/infrastructure"

	"github.com/robfig/cron/v3"
)


type Crons struct {
	ID int `json:"id"`
	Title string `json:"title"`
	CreatedAt int64 `json:"createdAt"`
}


type CronBlocks struct {
	ID int `json:"id"`
	Category string `json:"category"`
	CreatedAt int64 `json:"createdAt"`
}


func main() {

	i := 0

	c := cron.New()
	c.AddFunc("@every 1s", func() {
		i++
		call(i)
	})
	c.Start()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf(err.Error())
	}
}


func call(index int) {
	d := infrastructure.NewDB()
	db := d.Connect()

	foundCronBlocks := []CronBlocks{}
	// SELECT * FROM users;
	db.Find(&foundCronBlocks)
	if len(foundCronBlocks) > 0 {
		fmt.Printf("block %d\n", index)
		return
	}

	cronBlock := CronBlocks{
		Category: "category",
		CreatedAt: time.Now().Unix(),
	}
	if !db.NewRecord(&cronBlock) {
		db.Delete(&cronBlock)
		panic("could not create new record")
	}
	if err := db.Create(&cronBlock).Error; err != nil {
		db.Delete(&cronBlock)
		panic(err.Error())
	}

	for i := 0; i < 1000; i++ {
		cron := Crons{
			Title: fmt.Sprintf("title %d", index),
			CreatedAt: time.Now().Unix(),
		}
		if !db.NewRecord(&cron) {
			db.Delete(&cronBlock)
			panic("could not create new record")
		}
		if err := db.Create(&cron).Error; err != nil {
			db.Delete(&cronBlock)
			panic(err.Error())
		}
	}

	db.Delete(&cronBlock)

	fmt.Printf("created %d\n", index)
}
