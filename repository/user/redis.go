package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"keid/common/types"
	"keid/model"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type UserRepository struct {
	Client *redis.Client
}

type GetResult struct {
	Users  []model.User
	Cursor uint64
}

// Utility function to generate a key for a user ID.
func userIDKey(id uuid.UUID) string {
	return fmt.Sprintf("user:%d", id)
}

var ErrNotExist = errors.New("user does not exist")

func (r *UserRepository) Create(ctx context.Context, user model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to encode user: %w", err)
	}

	txn := r.Client.TxPipeline()

	key := userIDKey(user.ID)

	res := txn.SetNX(ctx, key, string(data), 0)
	if err := res.Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to set: %w", err)
	}

	if err := txn.SAdd(ctx, "users", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to add to users set: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (r *UserRepository) GetById(ctx context.Context, id uuid.UUID) (model.User, error) {
	key := userIDKey(id)

	value, err := r.Client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return model.User{}, ErrNotExist
	} else if err != nil {
		return model.User{}, fmt.Errorf("get user: %w", err)
	}

	var user model.User
	err = json.Unmarshal([]byte(value), &user)
	if err != nil {
		return model.User{}, fmt.Errorf("failed to decode user json: %w", err)
	}

	return user, nil
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	key := userIDKey(id)

	txn := r.Client.TxPipeline()

	err := txn.Del(ctx, key).Err()
	if errors.Is(err, redis.Nil) {
		txn.Discard()
		return ErrNotExist
	} else if err != nil {
		txn.Discard()
		return fmt.Errorf("get user: %w", err)
	}

	if err := txn.SRem(ctx, "users", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to remove from users set: %w", err)
	}

	if _, err := txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (r *UserRepository) Update(ctx context.Context, user model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to encode user: %w", err)
	}

	key := userIDKey(user.ID)

	err = r.Client.SetXX(ctx, key, string(data), 0).Err()
	if errors.Is(err, redis.Nil) {
		return ErrNotExist
	} else if err != nil {
		return fmt.Errorf("set user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetAll(ctx context.Context, page types.GetAllPage) (GetResult, error) {
	res := r.Client.SScan(ctx, "users", page.Offset, "*", int64(page.Size))

	keys, cursor, err := res.Result()
	if err != nil {
		return GetResult{}, fmt.Errorf("failed to get user ids: %w", err)
	}

	if len(keys) == 0 {
		return GetResult{
			Users: []model.User{},
		}, nil
	}

	xs, err := r.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return GetResult{}, fmt.Errorf("failed to get users: %w", err)
	}

	users := make([]model.User, len(xs))

	for i, x := range xs {
		x := x.(string)
		var user model.User

		err := json.Unmarshal([]byte(x), &user)
		if err != nil {
			return GetResult{}, fmt.Errorf("failed to decode user json: %w", err)
		}

		users[i] = user
	}

	return GetResult{
		Users:  users,
		Cursor: cursor,
	}, nil
}
