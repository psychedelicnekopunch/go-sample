package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/psychedelicnekopunch/go-sample/internal/infrastructure"

	"github.com/robfig/cron/v3"
)


/*
--
-- テーブルの構造 `crons`
--
CREATE TABLE `crons` (
  `id` int UNSIGNED NOT NULL,
  `title` text NOT NULL,
  `created_at` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- テーブルのインデックス `crons`
--
ALTER TABLE `crons`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `crons`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;
*/

type Crons struct {
	ID int `json:"id"`
	Title string `json:"title"`
	CreatedAt int64 `json:"createdAt"`
}

/*
--
-- テーブルの構造 `cron_blocks`
--
CREATE TABLE `cron_blocks` (
  `id` int UNSIGNED NOT NULL,
  `category` text NOT NULL,
  `created_at` int UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- テーブルのインデックス `cron_blocks`
--
ALTER TABLE `cron_blocks`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `cron_blocks`
  MODIFY `id` int UNSIGNED NOT NULL AUTO_INCREMENT;
*/

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
