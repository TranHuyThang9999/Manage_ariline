func connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=ariline port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect fail")
	} else {
		fmt.Print("Connect successfully")
	}
	return db
}
