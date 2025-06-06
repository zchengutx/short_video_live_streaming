package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/video"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(video.VideoUser{}, video.VideoWorks{}, video.VideoTopic{})
	if err != nil {
		return err
	}
	return nil
}
