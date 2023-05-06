package main

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

const (
	defaulShelfSize = 100
)

func NewDB(dsn string) (*gorm.DB, error) {
	//dsn := "root:123456@tcp(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		panic("failed to connect database")
	}
	//defer db.Close()

	//创建表  自动迁移(把结构体和数据表对应)
	db.AutoMigrate(&Shelf{}, &Book{})
	return db, nil
}

// 定义Model

// Shelf 书架
type Shelf struct {
	ID       int64 `gorm:"primaryKey"`
	Theme    string
	Size     int64
	CreateAt time.Time
	UpdateAt time.Time
}

// Book 图书
type Book struct {
	ID       int64 `gorm:"primaryKey"`
	Author   string
	Title    string
	ShelfID  int64
	CreateAt time.Time
	UpdateAt time.Time
}

// 数据库操作
type bookstore struct {
	db *gorm.DB
}

// CreateShelf 创建书架
func (b *bookstore) CreateShelf(ctx context.Context, data Shelf) (*Shelf, error) {
	if len(data.Theme) <= 0 {
		return nil, errors.New("invalid theme")
	}
	size := data.Size
	if size <= 0 {
		size = defaulShelfSize
	}
	v := Shelf{Theme: data.Theme, Size: size, CreateAt: time.Now(), UpdateAt: time.Now()}
	err := b.db.Create(&v).Error
	return &v, err
}

// GetShelf 获取书架
func (b *bookstore) GetShelf(ctx context.Context, id int64) (*Shelf, error) {
	v := Shelf{}
	err := b.db.First(&v, id).Error
	return &v, err
}

// ListShelves 书架列表
func (b *bookstore) ListShelves(ctx context.Context) ([]*Shelf, error) {
	var vl []*Shelf
	err := b.db.Find(&vl).Error
	return vl, err
}

// DeleteShelf 删除书架
func (b *bookstore) DeleteShelf(ctx context.Context, id int64) error {
	return b.db.Delete(&Shelf{}, id).Error
}
