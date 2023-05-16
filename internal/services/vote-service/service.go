package vote_service

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"vote-ui/internal/pb/vote_service"
)

type VoteService struct {
	voteService vote_service.VoteServiceClient
	publicKey   rsa.PublicKey
}

func NewVoteService(grpcClient *grpc.ClientConn) *VoteService {
	return &VoteService{
		voteService: vote_service.NewVoteServiceClient(grpcClient),
	}
}

func (v *VoteService) InitPublicKey(ctx context.Context) error {
	keyBytes, err := v.voteService.GetVotePublicKey(ctx, &vote_service.Empty{})
	if err != nil {
		return fmt.Errorf("GetVotePublicKey: %w", err)
	}

	err = json.Unmarshal(keyBytes.PublicKeyJson, &v.publicKey)
	if err != nil {
		return fmt.Errorf("unmarshal public key: %w", err)
	}
	return nil
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

func (v *VoteService) Vote(ctx context.Context, variantID int64, voteID int64, userID int64, userToken string) error {
	vote := vote_service.VoteVariantPair{
		VariantID: variantID,
		VoteID:    voteID,
	}

	byteMessage, err := proto.Marshal(&vote)
	if err != nil {
		return fmt.Errorf("proto.Marshal: %w", err)
	}
	hash := sha256.New()
	rnd := rand.Reader
	encryptedMessage, err := rsa.EncryptOAEP(hash, rnd, &v.publicKey, byteMessage, []byte("vote"))
	if err != nil {
		return fmt.Errorf("rsa.EncryptOAEP: %w", err)
	}

	bytes, err := json.Marshal(&v.publicKey)
	if err != nil {
		return fmt.Errorf("marshal public key: %w", err)
	}

	voteResp, err := v.voteService.Vote(ctx, &vote_service.VoteRequest{
		Auth: &vote_service.AuthInfo{
			UserID: userID,
			Token:  userToken,
		},
		CypherVote: encryptedMessage,
		PublicKey:  bytes,
	})
	if err != nil {
		return fmt.Errorf("rpc.Vote: %w", err)
	}

	if voteResp.Status == vote_service.VoteResponse_Cancelled {
		return fmt.Errorf("vote cancelled")
	}
	return nil
}
