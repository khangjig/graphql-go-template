package demo

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"graphql-go-template/model"
)

func NewPG(getDB func(ctx context.Context) *gorm.DB) Repository {
	return pgDemoRepository{getDB}
}

type pgDemoRepository struct {
	getDB func(ctx context.Context) *gorm.DB
}

func (p pgDemoRepository) GetTableName(ctx context.Context) string {
	return p.getDB(ctx).NewScope(model.Demo{}).GetModelStruct().TableName(p.getDB(ctx))
}

func (p pgDemoRepository) GetAll(ctx context.Context) ([]model.Demo, error) {
	var obj []model.Demo
	err := p.getDB(ctx).Find(&obj).Error

	return obj, errors.Wrap(err, "GetAll fail")
}

func (p pgDemoRepository) Create(ctx context.Context, obj *model.Demo) error {
	tx := p.getDB(ctx).Begin()

	if err := tx.Create(obj).Error; err != nil {
		tx.Rollback()

		return errors.Wrap(err, "Create fail")
	}

	err := tx.Commit().Error

	return errors.Wrap(err, "Commit fail")
}

func (p pgDemoRepository) GetByID(ctx context.Context, id int64) (*model.Demo, error) {
	var obj model.Demo
	err := p.getDB(ctx).First(&obj, id).Error

	return &obj, errors.Wrap(err, "GetByID fail")
}

func (p pgDemoRepository) Update(ctx context.Context, obj *model.Demo) error {
	tx := p.getDB(ctx).Begin()

	if err := tx.Save(obj).Error; err != nil {
		tx.Rollback()

		return errors.Wrap(err, "Update fail")
	}

	err := tx.Commit().Error

	return errors.Wrap(err, "Commit fail")
}
