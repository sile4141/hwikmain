package model

import hwikconf "github.com/hwikpass/hwik-go/core/conf/hwik"

type Settings struct {
	AuthMemberURIHwikgo string
	S3Region            string
	S3AccessID          string
	S3SecretPW          string
	S3Bucket            string
}

var Config = &Settings{}

// InitConfig : 환경변수 초기화.
func InitConfig() {
	cfg := &Settings{}
	cfg.AuthMemberURIHwikgo = hwikconf.GetValueOnly("auth.member.uri")
	cfg.S3Region = hwikconf.GetEncValueOnly("s3Region")
	cfg.S3AccessID = hwikconf.GetEncValueOnly("s3AccessID")
	cfg.S3SecretPW = hwikconf.GetEncValueOnly("s3SecretPW")
	cfg.S3Bucket = hwikconf.GetEncValueOnly("s3Bucket")
	Config = cfg
}
