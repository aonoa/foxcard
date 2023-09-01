package card

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"reflect"
)

const (
	MaxDuration      int64 = 60 * 60 * 24 * 365 * 5 // 卡密最大有效时长
	HeartbeatTimeout int64 = 60 * 3                 // 卡密心跳过期时间 3分钟无心跳过期
)

type LoginStatusCard struct {
	// id md5(cardId+设备+app_key)
	Id string
	// 登陆的时间戳
	Timestamp int64
	// 过期时间
	ExpiresTs int64
	// 卡密ID
	CardId string
	// app_key
	AppKey string
	// Device
	Device string
}

type HeartbeatToken struct {
	// LoginStatusID
	LoginStatusID string
	// 时间戳
	Timestamp int64
	// 过期时间
	ExpiresTs int64
	// app_key
	AppKey string
	// Device
	Device string
}

// GenerateCardID 创建卡密
func GenerateCardID() string {
	uuid := uuid.New()
	cardId := uuid.String()
	h := sha1.New()
	h.Write([]byte(cardId))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
	return fmt.Sprintf("%x", bs)
}

// GenerateStatusID 生成登陆状态id
func GenerateStatusID(cardId, device, appKey string) string {
	data := []byte(cardId + device + appKey)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

// GenerateToken 生成登陆状态id
func GenerateToken(cardId, device, appKey string) string {
	// 本次登陆的状态id
	// AppKey
	// token过期时间
	// 设备id

	return fmt.Sprintf("%x", "") //将[]byte转成16进制
}

func convertMapToStruct(m map[string]interface{}, s interface{}) {
	structValue := reflect.ValueOf(s).Elem()
	structType := structValue.Type()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structType.Field(i)
		fieldName := fieldType.Name

		if val, ok := m[fieldName]; ok {
			if field.Kind() == reflect.Int64 {
				field.SetInt(int64(val.(float64)))
			} else {
				field.Set(reflect.ValueOf(val))
			}
		}
	}
}
