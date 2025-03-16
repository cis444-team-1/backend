package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cis444-team-1/backend/config"
	"github.com/labstack/echo/v4"
)

// GeneratePresignedURL handles the upload of an image or audio file to AWS S3 and returns a publicly accessible URL.
//
// This function relies on the application's configuration to be correctly set up with the necessary AWS credentials and settings.
// It assumes that the CloudFront distribution is configured to serve the images from the specified S3 bucket.
//
// Parameters:
//   - c: echo.Context - The Echo HTTP context, providing access to the request and response.
//
// Returns:
//   - error: An error if any step in the process fails, otherwise nil. The error is returned as a JSON response to the client.
func (*Handler) GeneratePresignedFileURL(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid file request. Send file as FormFile."})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not read file."})
	}
	defer src.Close()

	cfg := config.LoadConfig()

	sess, err := session.NewSession(&aws.Config{
		Region:      &cfg.AWSRegion,
		Credentials: credentials.NewStaticCredentials(cfg.AWSAccessKey, cfg.AWSSecretAccessKey, ""),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not connect to the file storage."})
	}

	s3Svc := s3.New(sess)
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	fileName := fmt.Sprintf("%d%s", file.Size, fileExt)

	// Validate MIME type for images and audio
	contentType := file.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") && !strings.HasPrefix(contentType, "audio/") {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Unsupported file type."})
	}

	_, err = s3Svc.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(cfg.S3BucketName),
		Key:         aws.String(fileName),
		Body:        src,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not upload file.", "message": err.Error()})
	}

	fileURL := fmt.Sprintf("%s%s", cfg.AWSCloudfrontDomain, fileName)
	return c.JSON(http.StatusOK, map[string]string{"url": fileURL})
}
