package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/niltonSugawara/app/models"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "password"
	dbname = "lojinha_g"
	PostgresDriver = "postgres"
)

func main(){
	fmt.Println("comecou")
	postgresqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host,port,user,password,dbname)

	db,err := gorm.Open(PostgresDriver,postgresqlconn)
	CheckError(err)

	var item models.Item
	if result := db.Table("public.item").First(&item); result.Error != nil{panic(result.Error)}
	inserir := db.Begin()
	itemInserir := models.Item_inserir{Nome: "desktop",Marca: "dell",Valor: 3000.00,Descricao: "pc topzera"}
	inserir.Table("public.item").Create(&itemInserir)
	inserir.Commit()

	fmt.Println("Item: ", item)

	defer db.Close()

	err = db.Error

	CheckError(err)

	fmt.Println("Conectado")

}

func CheckError(err error)  {
	if err != nil {
		panic(err)
	}
	
}