package ec_store

import (
	"context"
	"github.com/golang/protobuf/proto"
	"icesfs/command/vars"
	"icesfs/ec/ec_store/ec_store_pb"
	"icesfs/errors"
	"icesfs/log"
	"icesfs/set"
)

func (shard *PlanShard) toPb() *ec_store_pb.PlanShard {
	if shard == nil {
		return nil
	}

	return &ec_store_pb.PlanShard{
		Host:      shard.Host,
		ShardSize: shard.ShardSize,
	}
}

func (plan *Plan) toPb() *ec_store_pb.Plan {
	if plan == nil {
		return nil
	}

	shardsPb := make([]*ec_store_pb.PlanShard, len(plan.Shards))
	for i, u := range plan.Shards {
		shardsPb[i] = u.toPb()
	}

	return &ec_store_pb.Plan{
		Set:        string(plan.Set),
		DataShards: plan.DataShards,
		Shards:     shardsPb,
	}
}

func planShardPbToInstance(pb *ec_store_pb.PlanShard) *PlanShard {
	if pb == nil {
		return nil
	}

	return &PlanShard{
		Host:      pb.Host,
		ShardSize: pb.ShardSize,
	}
}

func planPbToInstance(pb *ec_store_pb.Plan) *Plan {
	if pb == nil {
		return nil
	}

	shards := make([]PlanShard, len(pb.Shards))
	for i, u := range pb.Shards {
		if u == nil {
			continue
		}
		shards[i] = *planShardPbToInstance(u)
	}

	if pb.Shards == nil {
		shards = nil
	}

	return &Plan{
		Set:        set.Set(pb.Set),
		DataShards: pb.DataShards,
		Shards:     shards,
	}
}

func (shard *PlanShard) EncodeProto(ctx context.Context) ([]byte, error) {
	message := shard.toPb()
	b, err := proto.Marshal(message)
	if err != nil {
		log.Errorw("encode plan shard proto error", vars.UUIDKey, ctx.Value(vars.UUIDKey), vars.UserKey, ctx.Value(vars.UserKey), vars.ErrorKey, err.Error())
		err = errors.GetAPIErr(errors.ErrProto)
	}
	return b, err
}

func (plan *Plan) EncodeProto(ctx context.Context) ([]byte, error) {
	message := plan.toPb()
	b, err := proto.Marshal(message)
	if err != nil {
		log.Errorw("encode plan proto error", vars.UUIDKey, ctx.Value(vars.UUIDKey), vars.UserKey, ctx.Value(vars.UserKey), vars.ErrorKey, err.Error())
		err = errors.GetAPIErr(errors.ErrProto)
	}
	return b, err
}

func DecodePlanShardProto(ctx context.Context, b []byte) (*PlanShard, error) {
	message := &ec_store_pb.PlanShard{}
	if err := proto.Unmarshal(b, message); err != nil {
		log.Errorw("decode plan shard proto error", vars.UUIDKey, ctx.Value(vars.UUIDKey), vars.UserKey, ctx.Value(vars.UserKey), vars.ErrorKey, err.Error())
		return nil, errors.GetAPIErr(errors.ErrProto)
	}
	return planShardPbToInstance(message), nil
}

func DecodePlanProto(ctx context.Context, b []byte) (*Plan, error) {
	message := &ec_store_pb.Plan{}
	if err := proto.Unmarshal(b, message); err != nil {
		log.Errorw("decode plan shard proto error", vars.UUIDKey, ctx.Value(vars.UUIDKey), vars.UserKey, ctx.Value(vars.UserKey), vars.ErrorKey, err.Error())
		return nil, errors.GetAPIErr(errors.ErrProto)
	}
	return planPbToInstance(message), nil
}
