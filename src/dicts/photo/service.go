package photo

import (
	"2019_2_Shtoby_shto/src/config"
	"2019_2_Shtoby_shto/src/customType"
	"2019_2_Shtoby_shto/src/database"
	"2019_2_Shtoby_shto/src/dicts/models"
	transport "2019_2_Shtoby_shto/src/handle"
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"time"
)

type HandlerPhotoService interface {
	DownloadPhoto(photo *bufio.Reader) (customType.StringUUID, error)
	GetPhotoByUser(photoID customType.StringUUID) ([]byte, error)
}

type service struct {
	transport.HandlerImpl
	db database.IDataManager
}

func CreateInstance(db database.IDataManager) HandlerPhotoService {
	return &service{
		db: db,
	}
}

func (s service) DownloadPhoto(photo *bufio.Reader) (customType.StringUUID, error) {
	photoPath := config.GetInstance().ImagePath
	if err := os.MkdirAll(photoPath, os.ModePerm); err != nil {
		return "", err
	}
	newPhoto := models.Photo{
		TimeLoad: time.Now(),
		Path:     photoPath,
	}
	if err := s.db.CreateRecord(&newPhoto); err != nil {
		return "", err
	}
	file, err := os.Create(path.Join(photoPath, newPhoto.ID.String()+".jpg"))
	if err != nil {
		return "", err
	}
	buf := bytes.Buffer{}
	if _, err := buf.ReadFrom(photo); err != nil {
		return "", err
	}
	if _, err := bufio.NewWriter(file).Write(buf.Bytes()); err != nil {
		return "", err
	}
	return newPhoto.ID, nil
}

func (s service) GetPhotoByUser(photoID customType.StringUUID) ([]byte, error) {
	photoPath := config.GetInstance().ImagePath
	return ioutil.ReadFile(path.Join(photoPath, photoID.String()+".jpg"))
}
