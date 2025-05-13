package repository_test

import (
	_ "github.com/lib/pq"
)

//func setupTestDB() (*sql.DB, *db.Queries) {
//	connStr := "postgresql://postgres:1234@localhost:5432/test_chat_db?sslmode=disable"
//	database, err := sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatal("cannot connect to db", err)
//	}
//	queries := db.New(database)
//	return database, queries
//}
//
//func setupTestTx(t *testing.T) (*sql.Tx, *db.Queries) {
//	database, _ := setupTestDB()
//	tx, err := database.Begin()
//	require.NoError(t, err)
//
//	txQueries := db.New(tx)
//	return tx, txQueries
//}
//
//func TestListUserRepository_WithTestDB(t *testing.T) {
//	database, queries := setupTestDB()
//	defer database.Close()
//
//	repo := user.NewRepository(queries)
//
//	ctx := context.Background()
//	users, err := repo.ListUserRepository(ctx)
//	require.NoError(t, err)
//	require.NotEmpty(t, users)
//}
//
//func TestCreateUserRepository_WithTestDB(t *testing.T) {
//	tx, txQueries := setupTestTx(t)
//	defer tx.Rollback()
//
//	input := db.CreateUserParams{
//		Name:     "thisIsTest",
//		Email:    "testtest@example.com",
//		Phone:    "09031303642",
//		Password: "hashed_password",
//	}
//
//	repo := user.NewRepository(txQueries)
//
//	ctx := context.Background()
//	createdUser, err := repo.CreateUserRepository(ctx, input)
//
//	require.NoError(t, err, "CreateUserRepository failed: %v", err)
//	require.Equal(t, input.Name, createdUser.Name)
//	require.Equal(t, input.Email, createdUser.Email)
//}
