package vote_service

import (
	"time"
	"vote-ui/internal/pb/vote_service"
)

type Vote struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	FinishesAt string `json:"finishes-at"`
}

type Variant struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type VoteVariants struct {
	Vote
	Variants []*Variant
}

func pbToVote(pb *vote_service.ShortVoteInfo) *Vote {
	var (
		finishes         = time.Unix(pb.GetFinishes(), 0)
		year, month, day = time.Now().Local().Date()
		finishesFormat   string
	)
	if finishes.Local().Day() == day && finishes.Local().Month() == month && finishes.Local().Year() == year {
		finishesFormat = time.TimeOnly
	} else {
		finishesFormat = "02.01.2006"
	}

	return &Vote{
		ID:         pb.GetId(),
		Name:       pb.GetName(),
		FinishesAt: finishes.Format(finishesFormat),
	}
}

func pbToVoteVariants(pb *vote_service.ExtendedVoteInfo) *VoteVariants {
	res := &VoteVariants{
		Vote:     *pbToVote(pb.GetShort()),
		Variants: make([]*Variant, 0, len(pb.Variants)),
	}
	for _, v := range pb.Variants {
		res.Variants = append(res.Variants, &Variant{
			ID:   v.GetId(),
			Name: v.GetName(),
		})
	}
	return res
}
