package server

import (
	"context"
	"crypto/md5"
	"fmt"
	"foxcard/internal/biz/card"
	"foxcard/internal/conf"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

// NewWhiteListMatcher Path 白名单.
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.card.v1.Card/CreateCard"] = struct{}{}
	whiteList["/api.card.v1.Card/CardLogin"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}

// MiddlewareAuth Jwt Auth.
func MiddlewareAuth(ac *conf.Auth) middleware.Middleware {
	return selector.Server(
		jwt.Server(
			func(token *jwtv4.Token) (interface{}, error) {
				return []byte(ac.JwtKey), nil
			},
			jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
			jwt.WithClaims(func() jwtv4.Claims {
				return &jwtv4.MapClaims{}
			}),
		),
	).Match(NewWhiteListMatcher()).Build()
}

func MiddlewareJwt() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			return handler(ctx, req)
		}
	}
}

func MiddlewareSign() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			fmt.Println(req)
			sign := ParamsSign(req, "adfgadf", "sign", MD5ParamsSign)
			structValue := reflect.ValueOf(req).Elem()
			if structValue.FieldByName("Sign").IsValid() {
				if sign != structValue.FieldByName("Sign").String() {
					return nil, card.ErrSignInvalid
				}
			}
			return handler(ctx, req)
		}
	}
}

func ParamsSign(params interface{}, salt string, filter string, signFunc func(string) string) string {
	// 将请求参数的key提取出来，并排好序
	newKeys := make([]string, 0)

	structValue := reflect.ValueOf(params).Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		fieldType := structType.Field(i)
		fieldName := fieldType.Name

		if fieldName == filter || unicode.IsLower(rune(fieldName[0])) || FirstLower(fieldName) == filter {
			continue
		}
		newKeys = append(newKeys, fmt.Sprintf("%v", fieldName))
	}

	sort.Strings(newKeys)
	var originStr string
	// 将输入进行标准化的处理
	for _, v := range newKeys {
		originStr += fmt.Sprintf("%v=%v&", FirstLower(v), structValue.FieldByName(v))
	}
	originStr += fmt.Sprintf("salt=%v", salt)
	// 使用md5算法进行处理
	sign := signFunc(originStr)
	return sign
}

func MD5ParamsSign(str string) string {
	fmt.Println(str)
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func FirstLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
