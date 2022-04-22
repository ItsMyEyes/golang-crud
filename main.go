package main

import (
	_ "crud_v2/app/database"
	_ "crud_v2/app/enviroment"
	_ "crud_v2/app/redis"
	"crud_v2/routes"
)

func main() {
	defer routes.RunApplication()
}
