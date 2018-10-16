package operation

import (
	"db/model"
	"log"
	"time"
)

func AddVideo(authorId int, videoName string) (*model.Video, error) {

	videoId, err := uuid.NewV4()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	stmt, err := Conn.Prepare(`
INSERT INTO videos(id, author_id, name, display, display_time)
VALUES (?,?,?,?)`)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if _, err = stmt.Exec(videoId, authorId, videoName, ctime); err != nil {
		log.Println(err)
		return nil, err
	}

	res := &model.Video{
		VideoID:videoId,
	}

	defer stmt.Close()

	return res, nil
}
