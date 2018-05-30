package filters

import (
	"testing"
	"time"

	"github.com/husainaloos/event-bus/messages"
)

func TestAlwaysAllowFilter_Allow(t *testing.T) {
	type args struct {
		m messages.Message
	}
	tests := []struct {
		name string
		f    AlwaysAllowFilter
		args args
		want bool
	}{
		{
			name: "should allow any generic message",
			f:    AlwaysAllowFilter{},
			args: args{
				m: messages.Message{
					CreatedAt: time.Now(),
					ID:        "123",
					Payload:   "with payload",
					Tags:      nil,
				},
			},
			want: true,
		},

		{
			name: "should allow empty message",
			f:    AlwaysAllowFilter{},
			args: args{
				m: messages.Message{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Allow(tt.args.m); got != tt.want {
				t.Errorf("AlwaysAllowFilter.Allow() = %v, want %v", got, tt.want)
			}
		})
	}
}
