package model

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

type SnowflakeID int64

// 用于前端传递 string ID
func (s SnowflakeID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%d\"", s)), nil
}
func (s *SnowflakeID) UnmarshalJSON(data []byte) error {
	str := string(data)
	str = strings.Trim(str, "\"")

	if str == "" || str == "null" {
		*s = 0
		return nil
	}

	id, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	*s = SnowflakeID(id)
	return nil
}

// ✅ 让数据库能识别
func (s SnowflakeID) Value() (driver.Value, error) {
	return int64(s), nil
}
func (s *SnowflakeID) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		*s = SnowflakeID(v)
		return nil
	case []byte:
		// 数据库返回字节字符串
		i, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return err
		}
		*s = SnowflakeID(i)
		return nil
	case string:
		// 万一数据库驱动返回 string（通常是 text 类型）
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		*s = SnowflakeID(i)
		return nil
	default:
		return fmt.Errorf("cannot scan value %v into SnowflakeID", value)
	}
}
