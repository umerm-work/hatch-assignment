package postgres

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/umerm-work/hatch-assignment/data"
)

func (d *DB) List(sortProperty, orderProperty string) ([]data.Todo, error) {
	var out []data.Todo
	err := d.dbInstance.Order(fmt.Sprintf("%s %s", sortProperty, orderProperty)).Find(&out).Error
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return out, nil
}

func (d *DB) Delete(id int64) error {
	err := d.dbInstance.Delete(&data.Todo{}, data.Todo{ID: id}).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
func (d *DB) Update(in data.Todo) error {
	err := d.dbInstance.Updates(in).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
func (d *DB) Create(in data.Todo) error {
	err := d.dbInstance.Create(&in).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
