package utils

import (
	"encoding/json"
	log "github.com/cihub/seelog"
)

func GetMarshal(i interface{}) []byte {
	result, err := json.Marshal(i)
	if err != nil {
		log.Errorf("marshal err:%v", err)
		return []byte{}
	}
	return result
}

func GetUnmarshal(message string, dest interface{}) error {
	err := json.Unmarshal([]byte(message), &dest)
	if err != nil {
		log.Errorf("unmarshal err: %v", err)
	}
	return err
}
