package create_users

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/cmd/batch/config"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

type User struct {
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	Year     int    `yaml:"year"`
}

var Command = &cobra.Command{
	Use: "create-users",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("path to yaml must be set")
		}

		config, err := config.New()
		if err != nil {
			return err
		}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
		entClient, err := ent.Open("mysql", dsn)
		if err != nil {
			return err
		}

		f, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer f.Close()
		users := make([]*User, 0)
		if err := yaml.NewDecoder(f).Decode(&users); err != nil {
			return err
		}

		ctx := context.Background()
		for _, user := range users {
			b, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			group, err := entClient.Group.Create().
				SetName(user.Name).
				SetEncryptedPassword(string(b)).
				SetRole(group.RoleContestant).
				SetCreatedAt(timejst.Now()).
				Save(ctx)
			if err != nil {
				return err
			}

			log.Println(group.ID, group.Name)
		}
		return nil
	},
}
