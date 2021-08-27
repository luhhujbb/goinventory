package s3store

import (
    "context"
    "log"
    "github.com/luhhujbb/goinventory/ivtype"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "bytes"
)

var s3Client *s3.Client

func Load(store ivtype.Store) (string, error){
    ctx := context.TODO()
    cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
    if s3Client == nil {
        s3Client = s3.NewFromConfig(cfg)
    }

    input := &s3.GetObjectInput{
        Bucket: &store.Bucket,
        Key: &store.Key,
    }

    resp,err := s3Client.GetObject(ctx,input)
    if err != nil {
        log.Print(err)
        return "", err
    } else {
        //put body into a string
        buf := new(bytes.Buffer)
        buf.ReadFrom(resp.Body)
        newStr := buf.String()
        //return string from s3 and eventual error
        return newStr, err
    }
}
