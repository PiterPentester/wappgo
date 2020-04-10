package models

import (
	"fmt"
	"strconv"
)

type Update struct {
	id int64
}

func NewUpdate(userID int64, body string) (*Update, error) {
	id, err := client.Incr("update:next-id").Result()
	if err != nil {
		return nil, err
	}
	key := fmt.Sprintf("update:%d", id)
	pipe := client.Pipeline()
	pipe.HSet(key, "id", id)
	pipe.HSet(key, "user_id", userID)
	pipe.HSet(key, "body", body)
	pipe.LPush("updates", id)
	pipe.LPush(fmt.Sprintf("user:%d:updates", userID), id)
	_, err = pipe.Exec()
	if err != nil {
		return nil, err
	}
	return &Update{id}, nil
}

func (update *Update) GetBody() (string, error) {
	key := fmt.Sprintf("update:%d", update.id)
	return client.HGet(key, "body").Result()
}

func (update *Update) GetUser() (*User, error) {
	key := fmt.Sprintf("update:%d", update.id)
	userID, err := client.HGet(key, "user_id").Int64()
	if err != nil {
		return nil, err
	}
	return GetUserByID(userID)
}

func queryUpdates(key string) ([]*Update, error) {
	updateIDs, err := client.LRange(key, 0, 10).Result()
	if err != nil {
		return nil, err
	}
	updates := make([]*Update, len(updateIDs))
	for i, strID := range updateIDs {
		id, err := strconv.Atoi(strID)
		if err != nil {
			return nil, err
		}
		updates[i] = &Update{int64(id)}
	}
	return updates, nil
}

func GetAllUpdates() ([]*Update, error) {
	return queryUpdates("updates")
}

func GetUpdates(userID int64) ([]*Update, error) {
	key := fmt.Sprintf("user:%d:updates", userID)
	return queryUpdates(key)
}

func PostUpdate(userID int64, body string) error {
	_, err := NewUpdate(userID, body)
	return err
}
