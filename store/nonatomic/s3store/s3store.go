package s3store

import (
    "context"
    "log"
    "github.com/luhhujbb/goinventory/ivtype"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client s3.Client

func Load(store ivtype.Store) (interface{}, error){
    ctx := context.TODO()
    cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
    if &s3Client == nil {
        s3Client = *s3.NewFromConfig(cfg)
    }
    input := &s3.GetObjectInput{
        Bucket: &store.Bucket,
        Key: &store.Key,
    }
    resp,err := s3Client.GetObject(ctx,input)
    return resp, err
}
