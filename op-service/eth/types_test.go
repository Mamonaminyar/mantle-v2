package eth

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInputError(t *testing.T) {
	err := InputError{
		Inner: errors.New("test error"),
		Code:  InvalidForkchoiceState,
	}
	var x InputError
	if !errors.As(err, &x) {
		t.Fatalf("need InputError to be detected as such")
	}
	require.ErrorIs(t, err, InputError{}, "need to detect input error with errors.Is")
}

type scalarTest struct {
	name              string
	val               Bytes32
	fail              bool
	blobBaseFeeScalar uint32
	baseFeeScalar     uint32
}

func TestEcotoneScalars(t *testing.T) {
	testCases := []scalarTest{
		{"dirty padding v0 scalar", Bytes32{0: 0, 27: 1, 31: 2}, false, 0, math.MaxUint32},
		{"dirty padding v0 scalar v2", Bytes32{0: 0, 1: 1, 31: 2}, false, 0, math.MaxUint32},
		{"valid v0 scalar", Bytes32{0: 0, 27: 0, 31: 2}, false, 0, 2},
		{"invalid v1 scalar", Bytes32{0: 1, 7: 1, 31: 2}, true, 0, 0},
		{"valid v1 scalar with 0 blob scalar", Bytes32{0: 1, 27: 0, 31: 2}, false, 0, 2},
		{"valid v1 scalar with non-0 blob scalar", Bytes32{0: 1, 27: 123, 31: 2}, false, 123, 2},
		{"valid v1 scalar with non-0 blob scalar and 0 scalar", Bytes32{0: 1, 27: 123, 31: 0}, false, 123, 0},
		{"zero v0 scalar", Bytes32{0: 0}, false, 0, 0},
		{"zero v1 scalar", Bytes32{0: 1}, false, 0, 0},
		{"unknown version", Bytes32{0: 2}, true, 0, 0},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			sysConfig := SystemConfig{Scalar: tc.val}
			blobScalar, regScalar, err := sysConfig.EcotoneScalars()
			if tc.fail {
				require.NotNil(t, err)
			} else {
				require.Equal(t, tc.blobBaseFeeScalar, blobScalar)
				require.Equal(t, tc.baseFeeScalar, regScalar)
			}
		})
	}
}

func Test_toArray(t *testing.T) {
	type args struct {
		slice []byte
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "t1",
			args: args{
				slice: []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			},
			want: [16]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toArray(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toArray() = %v, want %v", got, tt.want)
			} else {
				fmt.Println(got)
			}
		})
	}
}
