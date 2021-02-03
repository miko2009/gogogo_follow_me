package material

import (
	"fmt"
	"github.com/miko2009/gogogo_follow_me/model/material"
)

func Show(materialId int) (musics []material.MusicModel) {
	musicModel := material.MusicModel{}
	query := musicModel.FindOne(materialId)
	rows, err := query.Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		err = query.ScanRows(rows, &musicModel)
		if err == nil {
			musics = append(musics, musicModel)
		}
	}
	return
}

func First() (musics []material.MusicModel) {
	musicModel := material.MusicModel{}
	result := musicModel.First(musicModel)
	fmt.Printf("%d", result.RowsAffected)
	rows, err := result.Rows()
	for rows.Next() {
		err = result.ScanRows(rows, &musicModel)
		if err == nil {
			musics = append(musics, musicModel)
		}
	}
	return
}

func User() {
	panic("this is error")
}
