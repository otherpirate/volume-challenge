package internal_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/otherpirate/volume-challenge/internal"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			path, err := internal.FindPath(tt.flights)
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
