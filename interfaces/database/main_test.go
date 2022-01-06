package database_test

// const (
// 	dbDriver = "postgres"
// 	dbSource = "user=root password=secret host=localhost dbname=ohlc sslmode=disable"
// )

// type TestRepository struct {
// 	repo *database.CurrencyRepository
// 	mock sqlmock.Sqlmock
// }

// var testRepo TestRepository

// func TestMain(m *testing.M) {
// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		log.Fatal("failed to init db mock:", err)
// 	}
// 	defer db.Close()

// 	repo := &database.CurrencyRepository{Db: db}
// 	testRepo = TestRepository{repo: repo, mock: mock}

// 	os.Exit(m.Run())
// }
