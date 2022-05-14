package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/examples/dal/model"
	"gorm.io/gen/examples/dal/query"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "C:\\Users\\lhqbalabala\\Desktop\\gorm\\gen\\examples\\dal/query",
	})
	db, _ := gorm.Open(mysql.Open("root:12345@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True"))
	g.UseDB(db)
	// generate all table from database
	g.ApplyInterface(func(method model.Method) {}, g.GenerateModelAs("users", "People"))

	//Create 100 users randomly
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		str := strconv.Itoa(rand.Intn(10))
		db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "uuid"}},
			DoUpdates: clause.Assignments(map[string]interface{}{"version": gorm.Expr("version + 1")}),
		}).Create(&model.People{str, "用户" + str, 0, 0})
	}
	//The function of grouping data by version and taking out the total number of users in the group with the highest version
	ctx, cancel := context.WithCancel(context.Background())
	u := query.Use(db).People
	sz, err := u.WithContext(ctx).FindMaxVersionCount()
	if err != nil {
		cancel()
	}
	fmt.Println(sz)
	g.Execute()
}
