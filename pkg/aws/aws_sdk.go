package aws

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	hwikerror "github.com/hwikpass/hwik-go/core/error"
	"github.com/hwikpass/hwik-go/core/logger"
	"github.com/hwikpass/hwik-sample/model"
)

type SDKInterface interface {
	InitAWSSession() *session.Session
	UploadBase64ToS3(base64File string, bucketFolderPath string, destBucket string) error
	UploadBytesToS3(file []byte, bucketFolderPath string, destBucket string) error
	DeleteFileFromS3(obj string, destBucket string) error
	DownloadFile(bucketFolderPath, destBucket string) ([]byte, error)
}

type SDKImpl struct {
}

// InitAWSSession : initialize AWS Session.
func (a *SDKImpl) InitAWSSession() *session.Session { // Create a single AWS session (we can re use this if we're uploading many files).
	s, err := session.NewSession(&aws.Config{
		Region:      aws.String(model.Config.S3Region),
		Credentials: credentials.NewStaticCredentials(model.Config.S3AccessID, model.Config.S3SecretPW, ""),
	})
	if err != nil {
		logger.Error(err)
	}

	// The session the S3 Uploader will use.
	sess := session.Must(s, err)

	return sess
}

// UploadBase64ToS3 : base64 파일 s3 업로드.
func (a *SDKImpl) UploadBase64ToS3(base64File string, bucketFolderPath string, destBucket string) error {
	decode, err := base64.StdEncoding.DecodeString(base64File)
	if err != nil {
		return hwikerror.InternalErrorWrap(err)
	}

	err = a.UploadBytesToS3(decode, bucketFolderPath, destBucket)
	if err != nil {
		return err
	}

	return err
}

// UploadBytesToS3 : bytes 파일 s3 업로드.
func (a *SDKImpl) UploadBytesToS3(file []byte, bucketFolderPath string, destBucket string) error {
	sess := a.InitAWSSession()
	acl := s3.ObjectCannedACLPublicRead
	bucket := destBucket

	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(bucketFolderPath),
		Body:   bytes.NewReader(file),
		ACL:    aws.String(acl),
	})
	if err != nil {
		return hwikerror.InternalErrorWrap(err, "failed to upload file")
	}
	logger.Debug("file uploaded to, %s\n", result.Location)

	return err
}

// DownloadFile : s3에서 파일 다운로드.
func (a *SDKImpl) DownloadFile(bucketFolderPath, destBucket string) ([]byte, error) {
	sess := a.InitAWSSession()
	bucket := destBucket

	// Create a downloader with the session and default options.
	downloader := s3manager.NewDownloader(sess)

	f, err := ioutil.TempFile(os.TempDir(), "pre-")
	if err != nil {
		return nil, hwikerror.InternalErrorWrap(err)
	}
	defer os.RemoveAll(f.Name())

	// Write the contents of S3 Object to the file.
	if _, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(bucketFolderPath),
	}); err != nil {
		return nil, hwikerror.InternalErrorWrap(err)
	}

	fileBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, hwikerror.InternalErrorWrap(err)
	}

	return fileBytes, nil
}

// DeleteFileFromS3 : S3 파일 삭제.
func (a *SDKImpl) DeleteFileFromS3(obj string, destBucket string) error {
	sess := a.InitAWSSession()
	bucket := destBucket

	svc := s3.New(sess)
	if _, err := svc.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucket), Key: aws.String(obj)}); err != nil {
		return hwikerror.InternalErrorWrap(err, "Unable to delete object %q from bucket %q", obj, bucket)
	}
	logger.Debug("Object %q successfully deleted\n", obj)

	return nil
}
