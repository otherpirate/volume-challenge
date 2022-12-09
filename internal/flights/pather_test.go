package flights_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/otherpirate/volume-challenge/internal/flights"
)

func TestInternal_Pather(t *testing.T) {
	tests := []struct {
		name    string
		flights [][]string
		want    string
		wantErr bool
	}{
		{
			name:    "single flight",
			flights: [][]string{{"SFO", "EWR"}},
			want:    "SFO->EWR",
		},
		{
			name:    "flight with one jump",
			flights: [][]string{{"ATL", "EWR"}, {"SFO", "ATL"}},
			want:    "SFO->EWR",
		},
		{
			name:    "flight with many jump",
			flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}},
			want:    "SFO->EWR",
		},
		{
			name:    "flight many times same airport jump",
			flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"EWR", "GSO"}, {"ATL", "EWR"}},
			want:    "SFO->EWR",
		},
		{
			name:    "invalid flight paths",
			flights: [][]string{{"IND", "EWR"}, {"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"}, {"ATL", "EWR"}},
			wantErr: true,
		},
		{
			name:    "missing flight",
			flights: [][]string{{"airport1", "airport2"}, {"airport3", "airport4"}},
			wantErr: true,
		},
		{
			name:    "invaid flight list - cycle", //I can define the order
			flights: [][]string{{"airport1", "airport3"}, {"airport3", "airport1"}},
			wantErr: true, // I can't define the order
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := flights.FindPath(tt.flights)
			fmt.Println(tt.name, path)
			if (err != nil) != tt.wantErr {
				t.Errorf("internal.FindPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(path, tt.want); diff != "" {
				t.Errorf("mismatch (-got +want):\n%s", diff)
				return
			}
		})
	}
}
