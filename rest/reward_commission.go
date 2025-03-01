package rest

import (
	"encoding/json"
	"strings"

	"go.uber.org/zap"
)

type rewardsAndCommisson struct {
	Height string `json:"height"`
	Result struct {
		Operator_Address string `json:"operator_address"`
		Selfbond_Rewards []Coin `json:"self_bond_rewards"`
		Commission       `json:"val_commission"`
	}
}

type Commission struct {
	Commission []Coin `json:"commission"`
}

func (rd *RESTData) getRewardsCommission() {
	var rc rewardsAndCommisson

	res, err := RESTQuery("/distribution/validators/" + OperAddr)
	if err != nil {
		zap.L().Fatal("", zap.Bool("Success", false), zap.String("err", err.Error()))
	}
	json.Unmarshal(res, &rc)
	if strings.Contains(string(res), "not found") {
		zap.L().Fatal("", zap.Bool("Success", false), zap.String("err", string(res)))
	} else if strings.Contains(string(res), "error:") || strings.Contains(string(res), "error\\\":") {
		zap.L().Fatal("", zap.Bool("Success", false), zap.String("err", string(res)))
	}

	rd.Rewards = rc.Result.Selfbond_Rewards
	rd.Commission = rc.Result.Commission.Commission
}
