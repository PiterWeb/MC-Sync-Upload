package src

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
)

var configFileMut sync.Mutex

func SetPlayersConfig(folderName string, rawConfig []byte) error {

	_, err := toml.Decode(string(rawConfig), &users)

	if err != nil {
		return err
	}

	statsDirectory := path.Join(folderName, "stats")
	playerdataDirectory := path.Join(folderName, "playerdata")

	var wg sync.WaitGroup

	errChan := make(chan error)

	for _, user := range users.Users {

		wg.Add(1)

		go func(user User, errChan chan<- error) {
			defer wg.Done()

			err := setUserConfig(statsDirectory, playerdataDirectory, user)

			if err != nil {
				errChan <- err
			}
		}(user, errChan)

	}

	wg.Wait()

	select {
	case err := <-errChan:
		close(errChan)
		return err
	case _ = <-time.After(1 * time.Second):
		close(errChan)
		return nil
	}

}

func generateOfflineUuid(name string) (string, error) {

	md5Hash := md5.New()

	_, err := io.WriteString(md5Hash, "OfflinePlayer:"+name)

	if err != nil {
		return "", err
	}

	hashInBytes := md5Hash.Sum(nil)

	hashInBytes[6] = (hashInBytes[6] & 0x0f) | 0x30
	hashInBytes[8] = (hashInBytes[8] & 0x3f) | 0x80

	hashInString := hex.EncodeToString(hashInBytes)

	hashInString = hashInString[:8] + "-" + hashInString[8:12] + "-" + hashInString[12:16] + "-" + hashInString[16:20] + "-" + hashInString[20:]

	return hashInString, nil
}

func getUUID(name string) (string, error) {

	res, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	type MCResponse struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Error string `json:"errorMessage"`
		Path  string `json:"path"`
	}

	mcResponse := MCResponse{}

	err = json.Unmarshal(body, &mcResponse)

	if err != nil {
		return "", err
	}

	UUID := mcResponse.ID[0:8] + "-" + mcResponse.ID[8:12] + "-" + mcResponse.ID[12:16] + "-" + mcResponse.ID[16:20] + "-" + mcResponse.ID[20:]

	return UUID, nil
}

func setUserConfig(statsDirectory string, playerdataDirectory string, user User) error {

	// Generate the Uuids if are not present
	if user.Uuid == "" || user.OfflineUuid == "" {
		if uuid, err := getUUID(user.Name); err == nil {
			user.Uuid = uuid
		}

		if uuid, err := generateOfflineUuid(user.Name); err == nil {
			user.OfflineUuid = uuid
		}

		if user.Uuid != "" || user.OfflineUuid != "" {
			configFileMut.Lock()
			defer configFileMut.Unlock()

		}
	}

	userPremiumStatsFile := path.Join(statsDirectory, user.Uuid+".json")
	userPremiumPlayerdataFile := path.Join(playerdataDirectory, user.Uuid+".dat")

	userNoPremiumStatsFile := path.Join(statsDirectory, user.OfflineUuid+".json")
	userNoPremiumPlayerdataFile := path.Join(playerdataDirectory, user.OfflineUuid+".dat")

	premiumStatsFileInfo, err := os.Stat(userPremiumStatsFile)

	if err != nil {
		f, err := os.Create(userPremiumStatsFile)

		if err != nil {
			return err
		}

		err = f.Close()

		if err != nil {
			return err
		}
	}

	noPremiumStatsFileInfo, err := os.Stat(userNoPremiumStatsFile)

	if err != nil {
		f, err := os.Create(userNoPremiumStatsFile)

		if err != nil {
			return err
		}

		err = f.Close()

		if err != nil {
			return err
		}

	}

	if premiumStatsFileInfo != nil && noPremiumStatsFileInfo != nil {

		if premiumStatsFileInfo.ModTime().After(noPremiumStatsFileInfo.ModTime()) {

			data, err := os.ReadFile(userPremiumStatsFile)

			if err != nil {
				return err
			}

			err = os.WriteFile(userNoPremiumStatsFile, data, os.ModePerm)

			if err != nil {
				return err
			}

		} else {

			data, err := os.ReadFile(userNoPremiumStatsFile)

			if err != nil {
				return err
			}

			err = os.WriteFile(userPremiumStatsFile, data, os.ModePerm)

			if err != nil {
				return err
			}
		}

	} else {

		if premiumStatsFileInfo != nil {
			data, err := os.ReadFile(userPremiumStatsFile)

			if err != nil {
				return err
			}

			err = os.WriteFile(userNoPremiumStatsFile, data, os.ModePerm)

			if err != nil {
				return err
			}

		}

		if noPremiumStatsFileInfo != nil {

			data, err := os.ReadFile(userNoPremiumStatsFile)

			if err != nil {
				return err
			}

			err = os.WriteFile(userPremiumStatsFile, data, os.ModePerm)

			if err != nil {
				return err
			}

		}

	}

	premiumPlayerdataFileInfo, err := os.Stat(userPremiumPlayerdataFile)

	if err != nil {
		f, err := os.Create(userPremiumPlayerdataFile)
		if err != nil {
			return err
		}

		err = f.Close()

		if err != nil {
			return err
		}

	}

	noPremiumPlayerdataFileInfo, err := os.Stat(userNoPremiumPlayerdataFile)

	if err != nil {
		f, err := os.Create(userNoPremiumPlayerdataFile)
		if err != nil {
			return err
		}

		err = f.Close()

		if err != nil {
			return err
		}
	}

	if premiumPlayerdataFileInfo != nil && noPremiumPlayerdataFileInfo != nil {

		if premiumPlayerdataFileInfo.ModTime().After(noPremiumPlayerdataFileInfo.ModTime()) {

			data, err := os.ReadFile(userPremiumPlayerdataFile)

			if err != nil {
				return err
			}

			err = os.WriteFile(userNoPremiumPlayerdataFile, data, os.ModePerm)

			if err != nil {
				return err
			}

		} else {

			data, err := os.ReadFile(userNoPremiumPlayerdataFile)

			if err != nil {
				return err
			}

			err = os.WriteFile(userPremiumPlayerdataFile, data, os.ModePerm)

			if err != nil {
				return err
			}

		}
	} else {

		if premiumPlayerdataFileInfo != nil {
			data, err := os.ReadFile(userPremiumPlayerdataFile)

			if err != nil {
				return err
			}

			err = os.WriteFile(userNoPremiumPlayerdataFile, data, os.ModePerm)

			if err != nil {
				return err
			}

		}

		if noPremiumPlayerdataFileInfo != nil {

			data, err := os.ReadFile(userNoPremiumPlayerdataFile)

			if err != nil {
				return err
			}

			err = os.WriteFile(userPremiumPlayerdataFile, data, os.ModePerm)

			if err != nil {
				return err
			}

		}

	}

	return nil
}
