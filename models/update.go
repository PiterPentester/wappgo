package models

import "fmt"

type Update struct {
	key string
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
	_, err = pipe.Exec()
	if err != nil {
		return nil, err
	}
	return &Update{key}, nil
}

func (update *Update) GetBody() (string, error) {
	return client.HGet(update.key, "body").Result()
}

func (update *Update) GetUser() (*User, error) {
	userID, err := client.HGet(update.key, "user_id").Int64()
	if err != nil {
		return nil, err
	}
	return GetUserByID(userID)
}

func GetUpdates() ([]*Update, error) {
	updateIDs, err := client.LRange("updates", 0, 10).Result()
	if err != nil {
		return nil, err
	}
	updates := make([]*Update, len(updateIDs))
	for i, id := range updateIDs {
		key := "update:" + id
		updates[i] = &Update{key}
	}
	return updates, nil
}

func PostUpdate(userID int64, body string) error {
	_, err := NewUpdate(userID, body)
	return err
}
