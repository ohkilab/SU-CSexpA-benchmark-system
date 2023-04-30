package main

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/group"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// config from env
	config, err := newConfig()
	if err != nil {
		log.Fatal(err)
	}

	// mysql client
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", config.DBUser, config.DBPass, config.DBHost, config.DBPort, config.DBName)
	entClient, err := ent.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte("ohkilab"), bcrypt.DefaultCost)
	_, err = entClient.Group.Create().SetID("ohkilab").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(99999).Save(ctx)
	if err != nil {
		log.Println(err)
	}
	_, err = entClient.Group.Create().SetID("a01").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(555).Save(ctx)
	if err != nil {
		log.Println(err)
	}
	_, err = entClient.Group.Create().SetID("a02").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(444).Save(ctx)
	if err != nil {
		log.Println(err)
	}
	_, err = entClient.Group.Create().SetID("a03").SetEncryptedPassword(string(encryptedPassword)).SetYear(2023).SetRole(group.RoleContestant).SetScore(666).Save(ctx)
	if err != nil {
		log.Println(err)
	}

	_, err = entClient.Contest.Create().
		SetQualifierStartAt(time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetQualifierEndAt(time.Date(2023, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetQualifierSubmitLimit(9999).
		SetFinalStartAt(time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)).
		SetFinalEndAt(time.Date(2024, time.December, 31, 23, 59, 59, 0, time.UTC)).
		SetFinalSubmitLimit(9999).
		Save(ctx)
	if err != nil {
		log.Println(err)
	}
}