package handler

import (
	"context"
	"fmt"
	recall "go-micro-etcd/service/general_recall_server"
)

type GeneralRecall struct{}

func (s *GeneralRecall) RecallResult(ctx context.Context, in *recall.RecallResultReq, out *recall.RecallResultRsp) error {
	uid := in.GetUid()

	fmt.Println("RecallResult() 15 line", in.Uid)
	var itemList []*recall.Item

	item := &recall.Item{}
	item.ItemId = "12345"
	item.ItemSceneId = 1
	item.ItemTypeId = 1
	item.ItemHotRank = 0.5
	item.ItemName = "test"
	item.ItemInfo = "test"
	item.ItemPrice = 99
	item.ItemTags = "test"
	item.ItemModifTime = 1610096132
	item.ItemCreateTimestamp = 1610096132
	item.ItemCity = "深圳"
	item.ItemCreateUserid = "12345"
	item.ReserveProperty_1 = "ReserveProperty_1"
	item.ReserveProperty_2 = "ReserveProperty_2"
	item.ReserveProperty_3 = "ReserveProperty_3"
	item.ReserveProperty_4 = "ReserveProperty_4"

	itemList = append(itemList, item)

	//recallResultRsp := general_recall.RecallResultRsp{}
	out.Code = 200
	out.Uid = uid
	out.ItemList = itemList
	return nil
}

func (s *GeneralRecall) RecallReloadRStrategy(ctx context.Context, in *recall.GeneralReloadRStrategyReq, out *recall.GeneralReloadRStrategyRsp) error {
	return nil
}

func (s *GeneralRecall) RecallReloadRSSubStrategy(ctx context.Context, in *recall.GeneralReloadRSSubtrategyReq, out *recall.GeneralReloadRSSubStrategyRsp) error {
	return nil
}
