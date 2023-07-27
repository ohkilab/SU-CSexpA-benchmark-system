package seed

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/cmd/batch/config"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/migrate"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/services/backend"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/bcrypt"
)

var Command = &cobra.Command{
	Use: "seed",
	RunE: func(cmd *cobra.Command, args []string) error {
		// config from env
		config, err := config.New()
		if err != nil {
			return err
		}

		// mysql client
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
		entClient, err := ent.Open("mysql", dsn)
		if err != nil {
			return err
		}

		ctx := context.Background()
		if err := entClient.Schema.Create(ctx, migrate.WithDropColumn(true)); err != nil {
			return err
		}

		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("ohkilab"), bcrypt.DefaultCost)
		_, err = entClient.Group.Create().SetName("a01").SetEncryptedPassword(string(encryptedPassword)).SetRole(group.RoleContestant).SetCreatedAt(timejst.Now()).Save(ctx)
		if err != nil {
			return err
		}
		_, err = entClient.Group.Create().SetName("a02").SetEncryptedPassword(string(encryptedPassword)).SetRole(group.RoleContestant).SetCreatedAt(timejst.Now()).Save(ctx)
		if err != nil {
			return err
		}
		_, err = entClient.Group.Create().SetName("a03").SetEncryptedPassword(string(encryptedPassword)).SetRole(group.RoleContestant).SetCreatedAt(timejst.Now()).Save(ctx)
		if err != nil {
			return err
		}

		_, err = entClient.Contest.Create().
			SetTitle("test contest(予選)").
			SetSlug("test-contest").
			SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
			SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
			SetSubmitLimit(9999).
			SetTagSelectionLogic(contest.TagSelectionLogicAuto).
			SetCreatedAt(timejst.Now()).
			SetValidator(backend.Validator_V2023.String()).
			SetTimeLimitPerTask(int64(30 * time.Second)).
			Save(ctx)
		if err != nil {
			return err
		}
		_, err = entClient.Contest.Create().
			SetTitle("test contest(本戦)").
			SetSlug("test-contest-ho").
			SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
			SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
			SetSubmitLimit(10).
			SetTagSelectionLogic(contest.TagSelectionLogicManual).
			SetCreatedAt(timejst.Now()).
			SetValidator(backend.Validator_V2023.String()).
			SetTimeLimitPerTask(int64(30 * time.Second)).
			Save(ctx)
		if err != nil {
			return err
		}
		return nil
	},
}
