package main

import (
	"context"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
	_ "time/tzdata"

	"github.com/tencentyun/cos-go-sdk-v5"
)

const (
	SecretID      = "secret_id"
	SecretKey     = "secret_key"
	BucketURL     = "bucket_url"
	SourcePath    = "source_path"
	TargetPath    = "target_path"
	StripPrefix   = "strip_prefix"
	PathSeparator = "/"
)

func ErrExit(err error) {
	log.Println(err.Error())
	os.Exit(1)
}

func GetConfig(key string) string {
	key = strings.ToUpper(key)
	if val := os.Getenv("INPUT_" + key); val != "" {
		return val
	}
	if val := os.Getenv("PLUGIN_" + key); val != "" {
		return val
	}
	return os.Getenv(key)
}

func StrIsEmpty(a ...string) bool {
	for _, v := range a {
		if v == "" {
			return true
		}
	}
	return false
}

func init() {
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "Asia/Shanghai"
	}
	time.Local, _ = time.LoadLocation(tz)
}

func main() {
	var (
		err         error
		secretID    = GetConfig(SecretID)
		secretKey   = GetConfig(SecretKey)
		bucketURL   = GetConfig(BucketURL)
		source      = GetConfig(SourcePath)
		target      = GetConfig(TargetPath)
		stripPrefix = GetConfig(StripPrefix)
	)
	if StrIsEmpty(secretID, secretKey, bucketURL, source, target) {
		ErrExit(errors.New("missing input parameter"))
	}
	sourceFiles := make([]string, 0, 0)
	if err = filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			sourceFiles = append(sourceFiles, path)
		}
		return nil
	}); err != nil {
		ErrExit(err)
	}
	sourceLen := len(sourceFiles)
	if sourceLen == 0 {
		return
	}
	if !strings.HasSuffix(target, PathSeparator) {
		target += PathSeparator
	}
	u, err := url.Parse(bucketURL)
	if err != nil {
		ErrExit(err)
	}
	cosClient := cos.NewClient(
		&cos.BaseURL{
			BucketURL: u,
		},
		&http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  secretID,
				SecretKey: secretKey,
			},
		})
	for i := 0; i < sourceLen; i++ {
		local := sourceFiles[i]
		remote := strings.TrimPrefix(local, stripPrefix)
		if strings.HasPrefix(remote, PathSeparator) {
			remote = strings.TrimPrefix(remote, PathSeparator)
		}
		remote = target + remote
		if _, err = cosClient.Object.PutFromFile(context.Background(), remote, local, &cos.ObjectPutOptions{
			ACLHeaderOptions:       &cos.ACLHeaderOptions{},
			ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{},
		}); err != nil {
			ErrExit(err)
		}
		log.Printf("source:%s target:%s\r\n", local, remote)
	}
}
