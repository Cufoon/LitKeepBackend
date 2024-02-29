package service

import (
	"fmt"
	"sort"

	"cufoon.litkeep.service/app/constant"
	"cufoon.litkeep.service/app/db/dao"
	"cufoon.litkeep.service/app/db/entity"
	"cufoon.litkeep.service/app/schema"
	"cufoon.litkeep.service/app/util"
)

func NewBillKindService(billKindDAO *dao.BillKindDAO) *BillKindService {
	return &BillKindService{billKindDAO: billKindDAO}
}

type BillKindService struct {
	billKindDAO *dao.BillKindDAO
}

func (brs *BillKindService) Create(record *schema.BillKind) error {
	return brs.billKindDAO.Create(record)
}

func generateMap(rmap map[string]*schema.BillKindTree, kinds []entity.BillKind) {
	tmpMap := make(map[string]*entity.BillKind, len(kinds))
	tmpList := make([]*entity.BillKind, 0, len(kinds))
	tmp2List := make([]*entity.BillKind, 0, len(kinds))
	for _, item := range kinds {
		item2 := item
		if item.ReplaceSystem != "" {
			tmpMap[item.ReplaceSystem] = &item2
		} else {
			tmpList = append(tmpList, &item2)
		}
	}
	if len(tmpMap) > 0 {
		for _, item := range tmpList {
			if value, ok := tmpMap[item.KindID]; ok {
				tmp2List = append(tmp2List, value)
			} else {
				item2 := item
				tmp2List = append(tmp2List, item2)
			}
		}
	} else {
		tmp2List = tmpList
	}
	for _, item := range tmp2List {
		if item == nil {
			break
		}
		if item.Owner == "KindLitWorldRoot" {
			value, ok := rmap[item.KindID]
			if ok {
				value.UserID = item.UserID
				value.KindID = item.KindID
				value.Name = item.Name
				value.Description = item.Description
			} else {
				tmp := new(schema.BillKindTree)
				tmp.UserID = item.UserID
				tmp.KindID = item.KindID
				tmp.Name = item.Name
				tmp.Description = item.Description
				rmap[item.KindID] = tmp
			}
		} else {
			value, ok := rmap[item.Owner]
			if ok {
				value.Children = append(value.Children, schema.BillKind{
					UserID:      item.UserID,
					KindID:      item.KindID,
					Name:        item.Name,
					Description: item.Description,
				})
			} else {
				tmp := new(schema.BillKindTree)
				tmp.Children = append(tmp.Children, schema.BillKind{
					UserID:      item.UserID,
					KindID:      item.KindID,
					Name:        item.Name,
					Description: item.Description,
				})
				rmap[item.Owner] = tmp
			}
		}
	}
}

func getKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (brs *BillKindService) QueryLinear(userID string) ([]*schema.BillKind, error) {
	userKinds, err := brs.billKindDAO.QueryByUserID(userID)
	if err != nil {
		return nil, err
	}
	r := make([]*schema.BillKind, len(userKinds))
	for idx, item := range userKinds {
		r[idx] = &schema.BillKind{
			UserID:      item.UserID,
			KindID:      item.KindID,
			Name:        item.Name,
			Description: item.Description,
		}
	}
	return r, nil
}

func (brs *BillKindService) Query(userID string) ([]*schema.BillKindTree, error) {
	userKinds, err := brs.billKindDAO.QueryByUserID(userID)
	if err != nil {
		return nil, err
	}
	rmap := make(map[string]*schema.BillKindTree)
	generateMap(rmap, userKinds)
	// for key, value := range rmap {
	// 	fmt.Printf("%s: %+v\n", key, value)
	// }
	keylst := getKeys(rmap)
	sort.Strings(keylst)
	r := make([]*schema.BillKindTree, len(keylst))
	for index, value := range keylst {
		r[index] = rmap[value]
	}
	return r, nil
}

func (brs *BillKindService) Modify(userID string, billKind *schema.BillKind) error {
	if billKind.UserID == "" || billKind.KindID == "" {
		return constant.ErrBillRecordModifyParamsWrong
	}
	fmt.Println(billKind)
	recordExist, err := brs.billKindDAO.QueryByKindID(billKind.KindID)
	fmt.Println(recordExist)
	if err != nil {
		return err
	}
	fmt.Println(recordExist)
	if recordExist.UserID == "Lit52317" {
		newKindID := util.GenerateBytes(16)
		for i := 0; i < 10; i++ {
			if !brs.billKindDAO.ExistByKindID(newKindID) {
				break
			}
			newKindID = util.GenerateBytes(16)
		}
		return brs.billKindDAO.Create(&schema.BillKind{
			UserID:        userID,
			KindID:        newKindID,
			Name:          billKind.Name,
			Owner:         recordExist.Owner,
			Description:   billKind.Description,
			ReplaceSystem: recordExist.KindID,
		})
	}
	return brs.billKindDAO.Modify(billKind)
}

func (brs *BillKindService) Delete(userID, kindID string) error {
	if userID == "" {
		return constant.ErrBillRecordDeleteNoUserID
	}
	return brs.billKindDAO.DeleteByUserIDAndKindID(userID, kindID)
}
