package main

import (
	"flag"
	"fmt"
	log "log/slog"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "config.yaml", "This tools will load its initial configuration from this file. Omit this false to use the default configuration.")
	flag.Parse()
}

func main() {
	cfg, err := GetConfig(configFile)
	if err != nil {
		log.Error("can't read configuration", "file", configFile, "error", err)
		os.Exit(1)
	}

	log.Info("Starting")

	awsCfg := &aws.Config{
		Region:      aws.String(cfg.S3Endpoint.Region),
		Endpoint:    aws.String(cfg.S3Endpoint.Endpoint),
		Credentials: credentials.NewStaticCredentials(cfg.S3Endpoint.AccessKey, cfg.S3Endpoint.SecretKey, ""),
		DisableSSL:  aws.Bool(true),
	}

	client, err := NewS3Client(awsCfg)
	if err != nil {
		log.Error("can't create AWS session", "error", err)
		os.Exit(1)
	}

	in := make(chan int)
	go func(num int) {
		for i := 0; i < num; i++ {
			in <- i
		}
		for i := 0; i < cfg.Concurrent; i++ {
			in <- -1
		}
	}(cfg.Total)

	result := processAndGather(in, func(i int) int {
		start := time.Now()
		if err := operateBucket(client, fmt.Sprintf("%s-%d-%d", cfg.BucketPrefile, cfg.Concurrent, i), cfg.Operation); err != nil {
			return -1
		} else {
			return int(time.Since(start) / (1000 * 1000)) // micro second
		}
	}, cfg.Concurrent)

	var total int64 = 0
	for i := 0; i < len(result); i++ {
		total += int64(result[i])
		log.Info(fmt.Sprintf("[Operation %d]", i), "time(ms)", result[i])
	}

	log.Info("Completed", "average time(ms)", total/int64(cfg.Total))
}
