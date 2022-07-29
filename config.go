package main

import (
	"github.com/BurntSushi/toml"
	"os"
	"path"
)

func setPlayersConfig(folderName string) error {

	_, err := toml.Decode(string(accountData), &users)

	if err != nil {
		return err
	}

	statsDirectory := path.Join(folderName, "stats")
	playerdataDirectory := path.Join(folderName, "playerdata")

	for _, user := range users.Users {

		if user.Uuid != "" && user.OfflineUuid != "" {

			userPremiumStatsFile := path.Join(statsDirectory, user.Uuid+".json")
			userPremiumPlayerdataFile := path.Join(playerdataDirectory, user.Uuid+".dat")

			userNoPremiumStatsFile := path.Join(statsDirectory, user.OfflineUuid+".json")
			userNoPremiumPlayerdataFile := path.Join(playerdataDirectory, user.OfflineUuid+".dat")

			premiumStatsFileInfo, err := os.Stat(userPremiumStatsFile)

			if err != nil {
				if userPremiumStatsFile != "" {
					_, err = os.Create(userPremiumStatsFile)

					if err != nil {
						return err
					}

				}
			}

			noPremiumStatsFileInfo, err := os.Stat(userNoPremiumStatsFile)

			if err != nil {
				if userNoPremiumStatsFile != "" {
					_, err = os.Create(userNoPremiumStatsFile)

					if err != nil {
						return err
					}

				}
			}

			if premiumStatsFileInfo != nil && noPremiumStatsFileInfo != nil {

				if premiumStatsFileInfo.ModTime().After(noPremiumStatsFileInfo.ModTime()) {

					data, err := os.ReadFile(userPremiumStatsFile)

					if err != nil {
						return err
					}

					os.WriteFile(userNoPremiumStatsFile, data, os.ModePerm)

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
				if userPremiumPlayerdataFile != "" {
					_, err = os.Create(userPremiumPlayerdataFile)
					if err != nil {
						return err
					}
				}
			}

			noPremiumPlayerdataFileInfo, err := os.Stat(userNoPremiumPlayerdataFile)

			if err != nil {
				if userNoPremiumPlayerdataFile != "" {
					_, err = os.Create(userNoPremiumPlayerdataFile)
					if err != nil {
						return err
					}
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
		}
	}

	return nil

}
