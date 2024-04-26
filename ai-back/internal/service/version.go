package service

import (
	"awesomeProject3/api"
	"errors"
	"github.com/rs/zerolog/log"
)

type VersionService interface {
	GetAllVersions(versions *[]api.Version) error
	AddVersion(version *api.Version) error
	judgeVersionIsEnable(version string) bool
	GetLatestVersion(version *api.Version) error
	UpdateVersion(version *api.Version) error
}

type versionService struct {
}

func NewVersionService() VersionService {
	return &versionService{}
}

func (v versionService) GetAllVersions(versions *[]api.Version) error {
	err := api.Db.Find(&versions)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (v versionService) AddVersion(version *api.Version) error {
	if version.Version == "" || version.Introduction == "" {
		return errors.New("版本名和介绍不可为空")
	}
	var ver api.Version
	api.Db.Where("version = ?", version.Version).First(&ver)
	if ver.ID != 0 {
		log.Error().Msg("添加版本失败，版本号重复")
		return errors.New("版本号重复")
	}
	err := api.Db.Create(&version)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (v versionService) GetLatestVersion(version *api.Version) error {
	err := api.Db.Last(&version)
	if err.Error != nil {
		return errors.New("获取最新版本失败")
	}
	return nil
}

func (v versionService) UpdateVersion(version *api.Version) error {
	if version.Version == "" {
		return errors.New("版本名不可为空")
	}

	if version.Enable {
		log.Info().Msg("更新版本：" + version.Version + "为启用状态")
	} else {
		log.Info().Msg("更新版本：" + version.Version + "为禁用状态")
	}

	err := api.Db.Model(&version).Where("id = ?", version.ID).Updates(map[string]interface{}{"enable": version.Enable,
		"introduction": version.Introduction,
		"download_url": version.DownloadUrl,
		"version":      version.Version,
	})

	if err.Error != nil {
		log.Error().Msg("更新版本失败：" + err.Error.Error())
		return err.Error
	}

	return nil
}

func (v versionService) judgeVersionIsEnable(version string) bool {
	var ver api.Version
	api.Db.Where("VersionService = ?", version).First(&ver)
	return ver.Enable
}
