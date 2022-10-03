package redis

import (
	"context"
	"github.com/swiggy-private/grill"
	"github.com/swiggy-private/grill/pkg/grillredis"
	"testing"
)

func Test_GrillRedis(t *testing.T) {
	helper := &grillredis.Redis{}
	if err := helper.Start(context.TODO()); err != nil {
		t.Errorf("error starting redis grill, error=%v", err)
		return
	}

	tests := []grill.TestCase{
		{
			Name: "Test_GetSet",
			Stubs: []grill.Stub{
				helper.SelectDB(1),
				helper.Set("one", "1"),
			},
			Action: func() interface{} {
				return nil
			},
			Assertions: []grill.Assertion{
				helper.AssertValue("one", "1"),
			},
			Cleaners: []grill.Cleaner{
				helper.FlushDB(),
			},
		},
	}

	grill.Run(t, tests)
}
