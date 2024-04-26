package exece

import "awesomeProject3/config"

func exectest() error {

	mysql := *config.GetConfig().MySQL
	err := mysql.BackupDb()

	if err != nil {
		return err
	}

	return nil
}
