package utils

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"testing"
)

type StringPair struct {
	M   string
	Url string
}

func BenckmarkGetModelId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		url := "https://my.matterport.com/show/?play=1&m=nothing"
		GetModelId(url)
	}
}

func TestGetModelId(t *testing.T) {
	url := "https://my.matterport.com/show/?play=1&m=abcd"
	result, _ := GetModelId(url)
	assert.Equal(t, result, "abcd")
}

func TestGetModelIdSchemaError(t *testing.T) {
	url := "://my.matterport.com/show/?play=1"
	_, err := GetModelId(url)
	assert.NotEqual(t, err, nil)
}

func TestGetModelIdProperties(t *testing.T) {
	properties := gopter.NewProperties(nil)

	properties.Property(
		"nil on nothing inside \"m\" parameter",
		prop.ForAll(
			func(url string) bool {
				result, _ := GetModelId(url)
				return result == nil
			},
			gen.AlphaString().Map(func(s string) string {
				return fmt.Sprintf("https://%v.com/show/?play=", s)
			}),
		),
	)
	//prop.
	properties.Property(
		"extracting the first \"m\" parameter",
		prop.ForAll(
			func(st StringPair) bool {
				result, _ := GetModelId(st.Url)
				return st.M == "" && result == nil || result != nil && *result == st.M
			},
			gen.AlphaString().Map(
				func(s string) StringPair {
					st := StringPair{
						M:   s,
						Url: fmt.Sprintf("https://my.matterport.com/show/?m=%v", s),
					}
					return st
				},
			),
		),
	)

	properties.TestingRun(t)
}

func TestCreatePayload(t *testing.T) {
	modelId := "abcd"
	result := CreatePayload(modelId)
	assert.Equal(t, len(result) > 0, true)
	var fields [3]string = [3]string{
		"operationName",
		"query",
		"variables",
	}
	for _, field := range fields {
		_, ok := result[field]
		assert.Equal(t, ok, true)
	}
	v, ok := result["variables"]
	assert.Equal(t, ok, true)
	extractedModelId, ok := (v.(map[string]string))["modelId"]
	assert.Equal(t, ok, true)
	assert.Equal(t, modelId, extractedModelId)
}

func TestCreateHeaders(t *testing.T) {
	userAgent := "test1"
	referer := "test2"
	result := CreateHeaders(userAgent, referer)
	assert.Equal(t, len(result) > 0, true)
	extractedUserAgent, ok := result["user-agent"]
	assert.Equal(t, ok, true)
	assert.Equal(t, userAgent, extractedUserAgent)
	extractedReferer, ok := result["referer"]
	assert.Equal(t, ok, true)
	assert.Equal(t, referer, extractedReferer)
}
