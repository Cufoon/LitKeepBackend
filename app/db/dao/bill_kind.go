package dao

import (
	"cufoon.litkeep.service/app/db/entity"
	"cufoon.litkeep.service/app/schema"
	"gorm.io/gorm"
)

func NewBillKindDAO(db *gorm.DB, billKindModel *entity.BillKind) *BillKindDAO {
	return &BillKindDAO{db: db, billKindModel: billKindModel}
}

type BillKindDAO struct {
	db            *gorm.DB
	billKindModel *entity.BillKind
}

func (bk *BillKindDAO) model() *gorm.DB {
	return bk.db.Model(bk.billKindModel)
}

func (bk *BillKindDAO) Create(billKind *schema.BillKind) error {
	return bk.model().Create(&entity.BillKind{
		KindID:        billKind.KindID,
		UserID:        billKind.UserID,
		Name:          billKind.Name,
		Description:   billKind.Description,
		Owner:         billKind.Owner,
		ReplaceSystem: billKind.ReplaceSystem,
	}).Error
}

func (bk *BillKindDAO) ExistByKindID(kindID string) bool {
	record := new(entity.BillKind)
	r := bk.model().Where("id_kind", kindID).Order("name asc").First(record)
	return r.Error == nil
}

func (bk *BillKindDAO) QueryByKindID(kindID string) (*entity.BillKind, error) {
	record := new(entity.BillKind)
	r := bk.model().Where("id_kind", kindID).Order("name asc").First(record)
	return record, r.Error
}

func (bk *BillKindDAO) QueryByUserID(userID string) ([]entity.BillKind, error) {
	var records []entity.BillKind
	var recordsSystem []entity.BillKind
	r := bk.model().Where("id_user", "Lit52317").Order("name asc").Find(&recordsSystem)
	if r.Error != nil {
		return nil, r.Error
	}
	r = bk.model().Where("id_user", userID).Order("name asc").Find(&records)
	return append(records, recordsSystem...), r.Error
}

func (bk *BillKindDAO) Modify(billKind *schema.BillKind) error {
	result := bk.model().Where("id_user", billKind.UserID).
		Where("id_kind", billKind.KindID).Updates(&entity.BillKind{
		Name:        billKind.Name,
		Description: billKind.Description,
	})
	return result.Error
}

func (bk *BillKindDAO) DeleteByUserID(userID string) error {
	result := bk.model().Where("id_user", userID).Delete(bk.billKindModel)
	return result.Error
}

func (bk *BillKindDAO) DeleteByUserIDAndKindID(userID, kindID string) error {
	result := bk.model().Where("id_user", userID).Where("id_kind", kindID).Delete(bk.billKindModel)
	return result.Error
}
