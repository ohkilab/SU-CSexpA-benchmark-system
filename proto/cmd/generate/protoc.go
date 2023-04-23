package main

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

type protocOptFunc func(lang string, mp map[string]string)

func optOut(path string) protocOptFunc {
	return func(lang string, mp map[string]string) {
		key := "--" + lang + "_out"
		mp[key] = path
	}
}

func optOpt(opt string) protocOptFunc {
	return func(lang string, mp map[string]string) {
		key := "--" + lang + "_opt"
		mp[key] = opt
	}
}

func optGrpcOut(path string) protocOptFunc {
	return func(lang string, mp map[string]string) {
		key := "--" + lang + "-grpc_out"
		mp[key] = path
	}
}

func optGrpcOpt(opt string) protocOptFunc {
	return func(lang string, mp map[string]string) {
		key := "--" + lang + "-grpc_opt"
		mp[key] = opt
	}
}

func optPlugin(name, path string) protocOptFunc {
	return func(lang string, mp map[string]string) {
		mp["--plugin"] = name + "=" + path
	}
}

func optRawOption(key, value string) protocOptFunc {
	return func(lang string, mp map[string]string) {
		mp[key] = value
	}
}

func runProtoc(lang, protoPath string, opts ...protocOptFunc) error {
	mp := make(map[string]string)
	for _, optFunc := range opts {
		optFunc(lang, mp)
	}
	args := []string{"protoc"}
	for key, val := range mp {
		args = append(args, key+"="+val)
	}
	cmd := exec.Command("npx", append(args, protoPath, "-I./vendor/protobuf/src", "-I./services")...)
	log.Println(cmd.String())
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}
	return nil
}
