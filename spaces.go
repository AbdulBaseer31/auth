package auth

import (
	"context"
	"fmt"
)

func Create_space(name string, authority int) error {
	if !init_check {
		return fmt.Errorf("run auth.Init() first as a function outside API calls")
	}

	_, err := conn.Exec(context.Background(),
		"INSERT INTO spaces(space_name, authority) VALUES ($1, $2)",
		name, authority,
	)

	if err != nil {
		return err
	}

	return nil
}

func Delete_space(name string) error {
	if !init_check {
		return fmt.Errorf("run auth.Init() first as a function outside API calls")
	}

	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM spaces WHERE space_name = $1",
		name,
	)
	
	if err != nil {
		return err
	}

	return nil
}
