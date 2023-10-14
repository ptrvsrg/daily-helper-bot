package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"

	"daily-helper-bot/internal/config"
	"daily-helper-bot/internal/log"
)

var db *sql.DB

func OpenDB() {
	dataSourceName := getDataSourceName()

	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Logger.Fatalf("Database error: %v", err)
	}
}

func getDataSourceName() string {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Config.DB.Username,
		config.Config.DB.Password,
		config.Config.DB.Host,
		config.Config.DB.Port,
		config.Config.DB.Name,
	)
	return dataSourceName
}

func CloseDB() {
	db.Close()
}

func FindByChatID(chatID int64) (*UserEntity, error) {
	statement := "SELECT users.chat_id, COALESCE(users.access_token, ''), " +
		"COALESCE(users.refresh_token, ''), scenarios.name FROM users " +
		"JOIN scenarios on users.scenario_id = scenarios.id WHERE users.chat_id = $1"
	row := db.QueryRow(statement, chatID)
	user := &UserEntity{}
	switch err := row.Scan(&user.ChatID, &user.AccessToken, &user.RefreshToken, &user.Scenario); {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil
	case err == nil:
		return user, nil
	default:
		log.Logger.Errorf("Database error: %v", err)
		return nil, err
	}
}

func DeleteByChatID(chatID int64) error {
	statement := "DELETE FROM users WHERE chat_id = $1"
	_, err := db.Exec(statement, chatID)
	if err != nil {
		log.Logger.Errorf("Database error: %v", err)
		return err
	}
	return nil
}

func SaveUser(chatID int64) error {
	statement := "INSERT INTO users(chat_id, scenario_id) " +
		"VALUES ($1, (SELECT id FROM scenarios WHERE name='start_not_authorized'))"
	_, err := db.Exec(statement, chatID)
	if err != nil {
		log.Logger.Errorf("Database error: %v", err)
		return err
	}
	return nil
}

func UpdateScenarioByChatID(chatID int64, scenario string) error {
	statement := "UPDATE users SET scenario_id=(SELECT id FROM scenarios WHERE name=$2) WHERE chat_id=$1"
	_, err := db.Exec(statement, chatID, scenario)
	if err != nil {
		log.Logger.Errorf("Database error: %v", err)
		return err
	}
	return nil
}

func UpdateTokensByChatID(chatID int64, accessToken, refreshToken string) error {
	statement := "UPDATE users SET access_token=$2, refresh_token=$3 WHERE chat_id=$1"
	_, err := db.Exec(statement, chatID, accessToken, refreshToken)
	if err != nil {
		log.Logger.Errorf("Database error: %v", err)
		return err
	}
	return nil
}
