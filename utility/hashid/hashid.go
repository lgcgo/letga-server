// Copyright Letga Author(https://letga.net). All Rights Reserved.
// Apache License 2.0(https://github.com/lgcgo/letga-server/blob/main/LICENSE)

package hashid

import (
	"fmt"
	"sync"

	"github.com/speps/go-hashids/v2"
)

type Config struct {
	Salt      string
	MinLength int
}

var (
	// 加密器
	hasher *hashids.HashID

	once sync.Once
)

// 实现单例模式
func hasherInstance(salt string, MinLength int) *hashids.HashID {
	var (
		hasher *hashids.HashID
		err    error
	)
	once.Do(func() {
		hd := hashids.NewData()
		hd.Salt = salt
		hd.MinLength = MinLength
		if hasher, err = hashids.NewWithData(hd); err != nil {
			panic("init hashid fail")
		}
	})
	return hasher
}

// 初始化
func Init(cfg *Config) {
	hasher = hasherInstance(cfg.Salt, cfg.MinLength)
}

// 加密表ID
func Encode(keyID uint, tableSalt int) (string, error) {
	var (
		key string
		err error
	)

	if key, err = hasher.Encode([]int{int(keyID), tableSalt}); err != nil {
		return "", err
	}

	return key, nil
}

// 批量加密表ID
func BatchEncode(keyIDs []uint, tableSalt int) ([]string, error) {
	var (
		keys []string
		err  error
		key  string
	)

	for _, v := range keyIDs {
		if key, err = Encode(v, tableSalt); err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}

	return keys, nil
}

// 解析Key值
func Decode(key string, tableSalt int) (uint, error) {
	var (
		nums []int
		err  error
	)

	if nums, err = hasher.DecodeWithError(key); err != nil {
		return 0, err
	}
	if len(nums) != 2 {
		return 0, fmt.Errorf("encrypted length greater than 1 detected")
	}
	if nums[1] != tableSalt {
		return 0, fmt.Errorf("decode salt match error")
	}

	return uint(nums[0]), nil
}

// 批量解析Key值
func BatchDecode(keys []string, tableSalt int) ([]uint, error) {
	var (
		keyIDs []uint
		err    error
		keyID  uint
	)

	for _, v := range keys {
		if keyID, err = Decode(v, tableSalt); err != nil {
			return nil, err
		}
		keyIDs = append(keyIDs, keyID)
	}

	return keyIDs, nil
}
