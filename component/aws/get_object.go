package aws

import (
	"context"
	s32 "github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
)

func (s *s3) GetObjectByKey(ctx context.Context, key string) ([]byte, error) {
	res, err := s.service.GetObject(&s32.GetObjectInput{
		Bucket: &s.cfg.s3Bucket,
		Key:    &key,
	})

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
