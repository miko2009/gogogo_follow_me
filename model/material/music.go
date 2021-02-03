package material

import (
	"fmt"
	"github.com/miko2009/gogogo_follow_me/lib/mysql"
	"gorm.io/gorm"
)

type MusicModel struct {
	Id        int    `form:"id" json:"id" gorm:"column:id;AUTO_INCREMENT"`
	MaterialId int64 `form:"material_id" json:"material_id" gorm:"material_id"`
	MusicUrl  string `form:"music_url" json:"music_url" gorm:"music_url"`
	MType     int    `form:"m_type" json:"type" gorm:"m_type"`
	Status    int    `form:"status" json:"status" gorm:"status"`
	CreatedAt int    `form:"created_at" json:"created_at" gorm:"created_at"`
	UpdatedAt int    `form:"updated_at" json:"updated_at" gorm:"updated_at"`
}


func (m *MusicModel) TableName() string {
	return "musics"
}

func (m *MusicModel) FindOne(materialId int) *gorm.DB {
	db := mysql.Master.Table(m.TableName())
	db.Where(fmt.Sprintf("material_id in (%d)", materialId))
	return db
}


func (m *MusicModel) First(music MusicModel) *gorm.DB {
	db := mysql.Slave.Table(m.TableName())
	result := db.First(&music)
	return result
}

func update(id int64, attributes MusicModel)  {

}