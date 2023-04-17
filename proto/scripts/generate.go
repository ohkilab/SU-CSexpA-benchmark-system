package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const servicesRoot = "services"

type ServiceYaml struct {
	GrpcServer map[string]LangYaml `yaml:"grpc-server"`
	Client     map[string]LangYaml `yaml:"client"`
}

type LangYaml struct {
	Dist   string   `yaml:"dist"`
	Opts   []string `yaml:"opts"`
	Plugin *struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	} `yaml:"plugin"`
}

func (l *LangYaml) GrpcOptions() []protocOptFunc {
	opts := make([]protocOptFunc, 0)
	opts = append(opts, optGrpcOut(l.Dist))
	if l.Plugin != nil {
		opts = append(opts, optPlugin(l.Plugin.Name, l.Plugin.Path))
	}
	for _, opt := range l.Opts {
		opts = append(opts, optGrpcOpt(opt))
	}
	return opts
}

func (l *LangYaml) ClientOptions() []protocOptFunc {
	opts := make([]protocOptFunc, 0)
	opts = append(opts, optClientOut(l.Dist))
	if l.Plugin != nil {
		opts = append(opts, optPlugin(l.Plugin.Name, l.Plugin.Path))
	}
	for _, opt := range l.Opts {
		opts = append(opts, optClientOpt(opt))
	}
	return opts
}

func main() {
	serviceEntries, err := os.ReadDir(servicesRoot)
	if err != nil {
		log.Fatal(err)
	}
	for _, serviceEntry := range serviceEntries {
		serviceYaml, protoFiles, err := loadService(serviceEntry.Name())
		if err != nil {
			log.Fatal(err)
		}

		for _, protoFile := range protoFiles {
			for langName, lang := range serviceYaml.GrpcServer {
				options := lang.GrpcOptions()
				if err := runProtoc(langName, protoFile, options...); err != nil {
					log.Fatal(err)
				}
			}
			for langName, lang := range serviceYaml.Client {
				options := lang.ClientOptions()
				if err := runProtoc(langName, protoFile, options...); err != nil {
					log.Fatal(err)
				}
			}
		}

	}
}

func loadService(serviceName string) (*ServiceYaml, []string, error) {
	serviceDirPath := filepath.Join("services", serviceName)
	f, err := os.Open(filepath.Join(serviceDirPath, "service.yaml"))
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()
	var serviceYaml ServiceYaml
	if err := yaml.NewDecoder(f).Decode(&serviceYaml); err != nil {
		return nil, nil, err
	}

	protoFiles, err := os.ReadDir(serviceDirPath)
	if err != nil {
		return nil, nil, err
	}
	protoFileList := make([]string, 0)
	for _, protoFile := range protoFiles {
		if !strings.HasSuffix(protoFile.Name(), ".proto") {
			continue
		}
		protoFileList = append(protoFileList, filepath.Join(serviceDirPath, protoFile.Name()))
	}
	return &serviceYaml, protoFileList, nil
}

type protocOptFunc func(lang string, mp map[string]string)

func optClientOut(path string) protocOptFunc {
	return func(lang string, mp map[string]string) {
		key := "--" + lang + "_out"
		mp[key] = path
	}
}

func optClientOpt(opt string) protocOptFunc {
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

func runProtoc(lang, protoPath string, opts ...protocOptFunc) error {
	mp := make(map[string]string)
	for _, optFunc := range opts {
		optFunc(lang, mp)
	}
	args := []string{}
	for key, val := range mp {
		args = append(args, key+"="+val)
	}
	cmd := exec.Command("protoc", append(args, protoPath, "-I./vendor/protobuf/src", "-I./services")...)
	log.Println(cmd.String())
	stderr := &bytes.Buffer{}
	cmd.Stderr = stderr
	if err := cmd.Run(); err != nil {
		return errors.New(stderr.String())
	}
	return nil
}
