package flights_test

import (
	"testing"

	"github.com/otherpirate/volume-challenge/internal/flights"
	"github.com/stretchr/testify/assert"
)

func TestInternal_GraphService(t *testing.T) {
	type args struct {
		flights [][]string
		start   string
		end     string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "single flight",
			args: args{
				flights: [][]string{{"SFO", "EWR"}},
				start:   "SFO",
				end:     "EWR",
			},
			want: []string{"SFO->EWR"},
		},
		{
			name: "flight with one jump",
			args: args{
				flights: [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}},
				start:   "SFO",
				end:     "EWR",
			},
			want: []string{"SFO->ATL->EWR"},
		},
		{
			name: "flight with many jump",
			args: args{
				flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
				start:   "SFO",
				end:     "EWR",
			},
			want: []string{"SFO->ATL->GSO->IND->EWR"},
		},
		{
			name: "flight with 2 paths",
			args: args{
				flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}, {"ATL", "EWR"}},
				start:   "SFO",
				end:     "EWR",
			},
			want: []string{
				"SFO->ATL->GSO->IND->EWR",
				"SFO->ATL->EWR",
			},
		},
		{
			name: "flight with 3 paths",
			args: args{
				flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"GSO", "EWR"}, {"ATL", "GSO"}, {"ATL", "EWR"}},
				start:   "SFO",
				end:     "EWR",
			},
			want: []string{
				"SFO->ATL->GSO->IND->EWR",
				"SFO->ATL->GSO->EWR",
				"SFO->ATL->EWR",
			},
		},
		{
			name: "graph with cycle",
			args: args{
				flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"GSO", "EWR"}, {"ATL", "GSO"}, {"ATL", "EWR"}, {"ATL", "SFO"}},
				start:   "SFO",
				end:     "EWR",
			},
			want: []string{
				"SFO->ATL->EWR",
				"SFO->ATL->GSO->EWR",
				"SFO->ATL->GSO->IND->EWR",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := flights.NewGraphService(tt.args.flights)
			result := s.FindPaths(tt.args.start, tt.args.end)
			if !assert.ElementsMatch(t, result, tt.want) {
				t.Fatalf("result is not equal\nExpected: %s\n   Gotta: %s", tt.want, result)
			}
		})
	}
}
