package crud

import (
	"context"
	"ent/ent"
	"ent/ent/user"
	"fmt"
	"log"
)

// CreateUser 新建用户
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.Create().
		SetAge(18).
		SetName("a8m").
		SetPhone("12345678").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil

}

// QueryUser 查询用户
func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// `Only` 如果没有发现用户则报错,
	// 否则正常返回。
	u, err := client.User.Query().Where(user.Name("a8m")).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func DeleteUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, _ := QueryUser(ctx, client)
	if err := client.User.DeleteOne(u).Exec(ctx); err != nil {
		return nil, fmt.Errorf("failed delete user: %w", err)

	}
	return nil, nil
}

func UpdateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, _ := QueryUser(ctx, client)
	u, err := client.User.UpdateOne(u).SetAge(28).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed update user: %w", err)
	}
	return u, nil

}
