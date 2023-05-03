package submit

import (
	"context"

	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/ent/contest"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/api/grpc/interceptor"
	"github.com/ohkilab/SU-CSexpA-benchmark-system/backend/server/core/timejst"
	pb "github.com/ohkilab/SU-CSexpA-benchmark-system/proto-gen/go/backend"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor struct {
	entClient *ent.Client
}

func NewInteractor(entClient *ent.Client) *Interactor {
	return &Interactor{entClient}
}

func (i *Interactor) PostSubmit(ctx context.Context, req *pb.PostSubmitRequest) (*pb.PostSubmitResponse, error) {
	claims := interceptor.GetClaimsFromContext(ctx)
	exists, err := i.entClient.Contest.Query().Where(contest.ID(int(req.ContestId))).Exist(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if !exists {
		return nil, status.Error(codes.InvalidArgument, "no such contest")
	}
	submit, err := i.entClient.Submit.Create().
		SetIPAddr(req.IpAddr).
		SetYear(claims.Year).
		SetSubmitedAt(timejst.Now()).
		Save(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.PostSubmitResponse{
		Id:         int32(submit.ID),
		IpAddr:     submit.IPAddr,
		ContestId:  req.ContestId,
		SubmitedAt: timestamppb.New(submit.SubmitedAt),
	}, nil
}
