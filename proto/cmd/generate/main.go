package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

const servicesRoot = "services"

type ServiceYaml struct {
	GrpcServer map[string]LangYaml `yaml:"grpc_server"`
	Client     map[string]LangYaml `yaml:"client"`
}

type LangYaml struct {
	Dist   string   `yaml:"dist"`
	Opts   []string `yaml:"opts"`
	Plugin *struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	} `yaml:"plugin"`
	RawOptions map[string]string `yaml:"raw_options"`
}

func (l *LangYaml) GrpcServerOptions() []protocOptFunc {
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
	opts = append(opts, optOut(l.Dist))
	if l.Plugin != nil {
		opts = append(opts, optPlugin(l.Plugin.Name, l.Plugin.Path))
	}
	for _, opt := range l.Opts {
		opts = append(opts, optOpt(opt))
	}
	for key, value := range l.RawOptions {
		opts = append(opts, optRawOption(key, value))
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
			for langName, lang := range serviceYaml.Client {
				options := lang.ClientOptions()
				if err := runProtoc(langName, protoFile, options...); err != nil {
					log.Fatal(err)
				}
			}
			for langName, lang := range serviceYaml.GrpcServer {
				options := lang.GrpcServerOptions()
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
	var serviceYaml *ServiceYaml
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
	return serviceYaml, protoFileList, nil
}
