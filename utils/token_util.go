package utils

import (
	"github.com/google/uuid"
	"strings"
)

/**
 * @Description 生成32位的token
 * @Param
 * @return
 **/
func CreateToken() string {
	get32UUID := uuid.New()
	uuidStr := get32UUID.String()
	idd := strings.Split(uuidStr, "-")
	return idd[0] + idd[1] + idd[2] + idd[3] + idd[4]
}
