// =================================================================
//
// Copyright (C) 2019 Spatial Current, Inc. - All Rights Reserved
// Released as open source under the MIT License.  See LICENSE file.
//
// =================================================================

package grwjs

import (
	"fmt"
)

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	awssession "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gopherjs/gopherjs/js"
	"github.com/pkg/errors"
)

import (
	"github.com/spatialcurrent/go-try-get/pkg/gtg"
)

import (
	"github.com/spatialcurrent/go-reader-writer/pkg/grw"
	"github.com/spatialcurrent/go-reader-writer/pkg/splitter"
)

func newS3Client(options map[string]interface{}) (*s3.S3, error) {

	region := gtg.TryGetString(options, "AWS_REGION", "")
	if len(region) == 0 {
		if defaultRegion := gtg.TryGetString(options, "AWS_DEFAULT_REGION", ""); len(defaultRegion) > 0 {
			region = defaultRegion
		}
	}

	config := aws.Config{
		MaxRetries: aws.Int(3),
		Region:     aws.String(region),
	}

	if accessKeyID := gtg.TryGetString(options, "AWS_ACCESS_KEY_ID", ""); len(accessKeyID) > 0 {
		if secretAccessKey := gtg.TryGetString(options, "AWS_SECRET_ACCESS_KEY", ""); len(secretAccessKey) > 0 {
			sessionToken := gtg.TryGetString(options, "AWS_SESSION_TOKEN", "")
			config.Credentials = credentials.NewStaticCredentials(
				accessKeyID,
				secretAccessKey,
				sessionToken)
		}
	}

	session, err := awssession.NewSessionWithOptions(awssession.Options{
		Config: config,
	})
	if err != nil {
		return nil, err
	}
	return s3.New(session), nil
}

var Exports = map[string]interface{}{
	"algorithms": grw.Algorithms,
	"schemes":    grw.Schemes,
	"read": func(uri string, alg string, options map[string]interface{}, callback func(...interface{}) *js.Object) {
		go func() {
			scheme, _ := splitter.SplitUri(uri)
			switch scheme {
			case grw.SchemeS3:
				s3Client, err := newS3Client(options)
				if err != nil {
					callback(nil, errors.Wrapf(err, "error creating s3 client").Error())
					return
				}
				reader, _, err := grw.ReadFromResource(uri, alg, 4096, s3Client)
				if err != nil {
					callback(nil, errors.Wrapf(err, "error opening reader for uri %q with compression algorithm %q", uri, alg).Error())
					return
				}
				callback(js.MakeWrapper(&Reader{Reader: reader}), nil)
			case grw.SchemeFile, grw.SchemeHTTP, grw.SchemeHTTPS, "":
				reader, _, err := grw.ReadFromResource(uri, alg, 4096, nil)
				if err != nil {
					callback(nil, errors.Wrapf(err, "error opening reader for uri %q with compression algorithm %q", uri, alg).Error())
					return
				}
				callback(js.MakeWrapper(&Reader{Reader: reader}), nil)
			default:
				callback(nil, fmt.Sprintf("error opening reader for uri %q: scheme %q is not supported", uri, scheme))
			}
		}()
	},
	/*"writer": func(uri string, alg string, callback func(...interface{}) *js.Object, options map[string]interface{}) {
		go func() {
			writer, err := grw.WriteToResource(uri, alg, false, nil)
			if err != nil {
				callback(nil, errors.Wrapf(err, "error opening writer for uri %q with compression algorithm %q", uri, alg).Error())
			}
			callback(js.MakeWrapper(writer), nil)
		}()
	},*/
}
