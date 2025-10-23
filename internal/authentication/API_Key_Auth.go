package authentication

import (
	"crypto/sha256"
	"fmt"
	"slices"
	"strings"

	"github.com/Abhishek191100/Enterprise-API-Gateway/logger"
	"github.com/Abhishek191100/Enterprise-API-Gateway/utils"
)

const API_KEY_LENGTH = 44

func ValidateKey(key string) bool {

	length := len(key)
	if length != API_KEY_LENGTH {
		logger.Log("Input API key length does not match, invalid key", "ERROR")
		return false
	}

	ind := strings.Index(key,"_")
	if ind == -1{
		logger.Log("Invalid key", "ERROR")
		return false
	}
	
	prefix := key[0:ind]
	if prefix != "EAGW"{
		logger.Log("Invalid key","ERROR")
		return false
	}

	keyWithoutPrefix := key[ind+1:]
	keyHash := sha256.Sum256([]byte(keyWithoutPrefix))
	strHash := fmt.Sprintf("%x",keyHash)
	
	hashes,err := utils.GetAllAPIKeyHashes()
	if err!=nil {
		logger.Log(err.Error(),"ERROR")
		return false
	}
	
	if !slices.Contains(hashes,strHash){
		logger.Log("API key does not exist","ERROR")
		return false
	}
	logger.Log("API key matched","INFO")
	return true
}