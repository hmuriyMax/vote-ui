package vote_service

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"vote-ui/internal/pb/vote_service"
)

type VoteService struct {
	voteService vote_service.VoteServiceClient
}

func NewVoteService(grpcClient *grpc.ClientConn) *VoteService {
	return &VoteService{
		voteService: vote_service.NewVoteServiceClient(grpcClient),
	}
}

func (v *VoteService) GetVotesByUserId(ctx context.Context, userID int64, userToken string) ([]*Vote, error) {
	resp, err := v.voteService.GetVotesForUser(ctx, &vote_service.GetVotesRequest{
		Auth: &vote_service.AuthInfo{
			UserID: userID,
			Token:  userToken,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("rpc.GetVotesForUser: %w", err)
	}

	var votes []*Vote
	for _, vote := range resp.GetVotes() {
		votes = append(votes, pbToVote(vote))
	}
	return votes, nil
}

func (v *VoteService) GetVoteById(ctx context.Context, voteId int64, userID int64, userToken string) (*VoteVariants, error) {
	resp, err := v.voteService.GetVoteInfo(ctx, &vote_service.GetVoteInfoRequest{
		VoteId: voteId,
		Auth: &vote_service.AuthInfo{
			UserID: userID,
			Token:  userToken,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("rpc.GetVoteById: %w", err)
	}

	return pbToVoteVariants(resp.GetInfo()), nil
}
