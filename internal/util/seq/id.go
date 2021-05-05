package seq

import (
	"context"
	"github.com/sony/sonyflake"
	"google.golang.org/grpc/metadata"
	"strconv"
	"time"
)

var (
	sf        *sonyflake.Sonyflake
	startTime = time.Date(2020, 3, 31, 0, 0, 0, 0, time.UTC)
)

func init() {
	var st sonyflake.Settings
	st.StartTime = startTime
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		panic("sonyflake not created")
	}
}

func GenNumID() uint64 {
	id, _ := sf.NextID()
	return id
}

func GenID() string {
	id, _ := sf.NextID()
	return strconv.FormatUint(id, 10)
}

const IdPrefixKey = "p-id"

func GenIdByCtx(ctx context.Context) string {
	nextId, err := sf.NextID()
	if err != nil {
		return err.Error()
	}
	// uit64 转成 str
	id := strconv.FormatUint(nextId, 10)

	if ctx == nil {
		ctx = context.Background()
	}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ps := md.Get(IdPrefixKey); len(ps) != 0 {
			return ps[0] + id
		}
	}

	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		if ps := md.Get(IdPrefixKey); len(ps) != 0 {
			return ps[0] + id
		}
	}
	return id
}
