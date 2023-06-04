package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/group"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/repository/ent/migrate"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// config from env
	config, err := newConfig()
	if err != nil {
		panic(err)
	}

	// mysql client
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	entClient, err := ent.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	if err := entClient.Schema.Create(ctx, migrate.WithDropColumn(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("ohkilab"), bcrypt.DefaultCost)
	_, err = entClient.Group.Create().SetName("ohkilab").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(99999).SetCreatedAt(timejst.Now()).Save(ctx)
	if err != nil {
		log.Println(err)
	}
	_, err = entClient.Group.Create().SetName("a01").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(555).SetCreatedAt(timejst.Now()).Save(ctx)
	if err != nil {
		log.Println(err)
	}
	_, err = entClient.Group.Create().SetName("a02").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(444).SetCreatedAt(timejst.Now()).Save(ctx)
	if err != nil {
		log.Println(err)
	}
	_, err = entClient.Group.Create().SetName("a03").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(666).SetCreatedAt(timejst.Now()).Save(ctx)
	if err != nil {
		log.Println(err)
	}

	_, err = entClient.Contest.Create().
		SetTitle("test contest").
		SetStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetSubmitLimit(9999).
		SetYear(2023).
		SetTagSelectionLogic(contest.TagSelectionLogicAuto).
		SetCreatedAt(timejst.Now()).
		Save(ctx)
	if err != nil {
		log.Println(err)
	}
}
