package util

import "../global"

func CheckErr(err error) {
	if err != nil {
		global.Logger.Error("An Error occured : ", err)
	}

	RollingLog()
}
