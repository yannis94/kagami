package helpers

import (
	"fmt"
	"os"
	"path"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var UploadImgPath string 

func init() {
    currentDir, err := os.Getwd()

    if err != nil {
        fmt.Printf("Error getting working dir: %s\n", err.Error())
    }

    UploadImgPath = path.Join(currentDir, "src/images/upload")
}

func FromStringToArray(s string, sep string) []string {
    return strings.Split(s, sep)
}

func FromArrayToString(a []string, sep string) string {
    var result string

    for i:=0; i<len(a); i++ {
        result += a[i]

        if i < len(a) - 1 {
            result += ","
        }
    }

    return result
}

func DeleteFile(filePath string) error {
    if filePath == "" {
        return nil
    }

    return os.Remove(filePath)
}

func HashPassword(pwd string) (string, error) {
    hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

    if err!= nil {
        return "", err
    }

    return string(hashedPwd), nil
}

func IsPasswordCorrect(hash, pwd string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
}

