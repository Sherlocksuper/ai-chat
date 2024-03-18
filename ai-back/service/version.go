package service

import (
	"awesomeProject3/api"
	"errors"
)

type VersionService interface {
	GetAllVersions(versions *[]api.Version) error
	AddVersion(version *api.Version) error
	judgeVersionIsEnable(version string) bool
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
	err := api.Db.Create(&version)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (v versionService) judgeVersionIsEnable(version string) bool {
	var ver api.Version
	api.Db.Where("VersionService = ?", version).First(&ver)
	return ver.Enable
}
