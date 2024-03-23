package handler

import (
	"awesomeProject3/api"
	"awesomeProject3/service"
	"github.com/gin-gonic/gin"
)

type VersionHandler struct {
	versionService service.VersionService
}

func NewVersionHandler(versionService service.VersionService) *VersionHandler {
	return &VersionHandler{
		versionService: versionService,
	}
}

// GetAllVersions 获取所有版本 GET
func (f *VersionHandler) GetAllVersions(c *gin.Context) {
	var versions []api.Version
	err := f.versionService.GetAllVersions(&versions)
	if err != nil {
		c.JSON(400, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "获取成功", versions))
}

func (f *VersionHandler) AddVersion(c *gin.Context) {
	var version *api.Version
	err := c.BindJSON(&version)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, "参数错误", nil))
		return
	}
	err = f.versionService.AddVersion(version)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "添加成功", nil))
}

// GetLatestVersion 获取最新版本
func (f *VersionHandler) GetLatestVersion(c *gin.Context) {
	var version api.Version
	err := f.versionService.GetLatestVersion(&version)
	if err != nil {
		c.JSON(200, api.M(api.FAIL, err.Error(), nil))
		return
	}
	c.JSON(200, api.M(api.SUCCESS, "获取成功", version))
}
