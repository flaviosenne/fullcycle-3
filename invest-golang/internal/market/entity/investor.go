package entity

type Investor struct {
	ID            string
	Name          string
	AssetPosition []*InvestorAssetPosition
}

func NewInvestor(id string) *Investor {
	return &Investor{
		ID:            id,
		AssetPosition: []*InvestorAssetPosition{},
	}
}

func (i *Investor) AddAssetPosition(assetPosition *InvestorAssetPosition) {
	i.AssetPosition = append(i.AssetPosition, assetPosition)
}

func (i *Investor) UpdateAssetPostion(assetID string, qtdShares int) {
	assetPostion := i.GetAssetPosition(assetID)
	if assetPostion == nil {
		i.AssetPosition = append(i.AssetPosition, NewInvestoAssetPosition(assetID, qtdShares))
	} else {
		assetPostion.Shares += qtdShares
	}

}

func (i *Investor) GetAssetPosition(assetID string) *InvestorAssetPosition {
	for _, assetPostion := range i.AssetPosition {
		if assetPostion.ID == assetID {
			return assetPostion
		}
	}
	return nil

}

type InvestorAssetPosition struct {
	ID     string
	Shares int
}

func NewInvestoAssetPosition(assetID string, shares int) *InvestorAssetPosition {
	return &InvestorAssetPosition{
		ID:     assetID,
		Shares: shares,
	}
}
