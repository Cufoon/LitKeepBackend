package service

import (
	"fmt"
	"math"

	"cufoon.litkeep.service/app/constant"
	"cufoon.litkeep.service/app/db/dao"
	"cufoon.litkeep.service/app/db/dto"
	"cufoon.litkeep.service/app/db/entity"
	"cufoon.litkeep.service/app/schema"
	"cufoon.litkeep.service/app/vo"
)

func NewBillRecordService(billRecordDAO *dao.BillRecordDAO) *BillRecordService {
	return &BillRecordService{billRecordDAO: billRecordDAO}
}

type BillRecordService struct {
	billRecordDAO *dao.BillRecordDAO
}

func (brs *BillRecordService) Create(record *schema.BillRecord) error {
	return brs.billRecordDAO.Create(record)
}

func (brs *BillRecordService) QueryRecordCount(userID string) (int64, error) {
	if userID == "" {
		return 0, constant.ErrBillRecordQueryNoUserID
	}
	count := brs.billRecordDAO.QueryByUserIDForCount(userID)
	return count, nil
}

func (brs *BillRecordService) Query(query *schema.BillRecordQuery) ([]entity.BillRecord, error) {
	if query.UserID == "" {
		return nil, constant.ErrBillRecordQueryNoUserID
	}
	which := 0
	if query.KindID != "" {
		which++
	}
	if query.StartTime != constant.ZeroTime &&
		query.EndTime != constant.ZeroTime &&
		query.StartTime.Before(query.EndTime) {
		which += 2
	}
	var result []entity.BillRecord
	switch which {
	case 0:
		result = brs.billRecordDAO.QueryByUserID(query.UserID)
	case 1:
		result = brs.billRecordDAO.QueryByUserIDAndKind(query.UserID, query.KindID)
	case 2:
		result = brs.billRecordDAO.QueryByUserIDPeriod(query.UserID, query.StartTime, query.EndTime)
	case 3:
		result = brs.billRecordDAO.QueryByUserIDAndKindPeriod(query.UserID, query.KindID, query.StartTime, query.EndTime)
	}
	return result, nil
}

func (brs *BillRecordService) QueryPage(query *schema.BillRecordPageQuery) ([]entity.BillRecord, error) {
	if query.UserID == "" {
		return nil, constant.ErrBillRecordQueryNoUserID
	}
	result := brs.billRecordDAO.QueryByUserIDAndPage(query.UserID, query.Page)
	return result, nil
}

func (brs *BillRecordService) QueryStatisticsDay(userID string, query *vo.BillRecordStatisticsDayQueryReq) ([]dto.QueryStatisticsDayData, error) {
	if userID == "" {
		return nil, constant.ErrBillRecordQueryNoUserID
	}
	oResult := brs.billRecordDAO.QueryStatisticsDay(userID, query.StartTime, query.EndTime)
	diff := query.EndTime.Sub(query.StartTime)
	diffDays := int64(math.Floor(diff.Abs().Hours() / 24))
	dateList := make([]string, 0, diffDays)
	for d := query.StartTime.AddDate(0, 0, 1); !d.After(query.EndTime); d = d.AddDate(0, 0, 1) {
		dateList = append(dateList, d.Format("2006-01-02"))
	}
	result := make([]dto.QueryStatisticsDayData, 0, diffDays)
	idx1 := 0
	idx2 := 0
	fmt.Println("dateList", len(dateList))
	for idx1 < len(dateList) && idx2 < len(oResult) {
		if dateList[idx1] == oResult[idx2].Day {
			result = append(result, dto.QueryStatisticsDayData{Day: oResult[idx2].Day, Money: oResult[idx2].Money})
			idx1++
			idx2++
		} else {
			result = append(result, dto.QueryStatisticsDayData{Day: dateList[idx1], Money: 0})
			idx1++
		}
	}
	return result, nil
}

func (brs *BillRecordService) Modify(billRecord *schema.BillRecord) error {
	if billRecord.ID == 0 || billRecord.UserID == "" || billRecord.KindID == "" {
		return constant.ErrBillRecordModifyParamsWrong
	}
	return brs.billRecordDAO.Modify(billRecord)
}

func (brs *BillRecordService) Delete(ids []uint, userID string) (notDeleted []uint, err error) {
	if userID == "" {
		return nil, constant.ErrBillRecordDeleteNoUserID
	}
	notDeleted, err = brs.billRecordDAO.Delete(ids, userID)
	return
}
