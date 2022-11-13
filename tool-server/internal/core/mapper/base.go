package mapper

import (
	"github.com/sirupsen/logrus"
	"tool-server/internal/global"
)

type BaseMapper[T any] struct {
}

type Page[T any] struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Result   []T `json:"result"`
}

func (m BaseMapper[T]) Insert(entity T) (err error) {
	logrus.Debugf("insert %+v", entity)
	global.DBMaxThread <- struct{}{}
	err = global.DBClient.Create(&entity).Error
	<-global.DBMaxThread
	return err
}

func (m BaseMapper[T]) Delete(entity T) (err error) {
	logrus.Debugf("delete %+v", entity)
	global.DBMaxThread <- struct{}{}
	err = global.DBClient.Delete(*&entity).Error
	<-global.DBMaxThread
	return err
}

func (m BaseMapper[T]) Update(entity T) (err error) {
	logrus.Debugf("update %+v", entity)
	global.DBMaxThread <- struct{}{}
	err = global.DBClient.Updates(*&entity).Error
	<-global.DBMaxThread
	return err
}

func (m BaseMapper[T]) Detail(entity T) (res T, err error) {
	logrus.Debugf("detail %+v", entity)
	global.DBMaxThread <- struct{}{}
	err = global.DBClient.First(&res, *&entity).Error
	<-global.DBMaxThread
	return res, err
}

func (m BaseMapper[T]) List(entity T) (res []T, err error) {
	logrus.Debugf("list %+v", entity)
	global.DBMaxThread <- struct{}{}
	err = global.DBClient.Find(&res, *&entity).Error
	<-global.DBMaxThread
	return res, err
}

func (m BaseMapper[T]) Count(entity T) (res int64, err error) {
	logrus.Debugf("count %+v", entity)
	global.DBMaxThread <- struct{}{}
	err = global.DBClient.Model(&entity).Count(&res).Error
	<-global.DBMaxThread
	return res, err
}

func (m BaseMapper[T]) Page(entity T, page int, pageSize int) (res Page[T], err error) {
	logrus.Debugf("page %+v, page %d, pageSize %d", entity, page, pageSize)
	offset := (page - 1) * pageSize
	res.Page = page
	res.PageSize = pageSize
	r := make([]T, 0)
	global.DBMaxThread <- struct{}{}
	err = global.DBClient.Order("id DESC").Where(entity).Offset(offset).Limit(pageSize).Find(&r).Error
	<-global.DBMaxThread
	res.Result = r
	return res, err
}
