package gobrtempo

import (
	"reflect"
	"testing"
	"time"
)

func TestDataFormatoBrasileiroParaTime(t *testing.T) {

	utc, _ := time.LoadLocation("America/Sao_Paulo")

	type args struct {
		dataEmFormatoBrasileiro string
	}
	tests := []struct {
		name      string
		args      args
		wantTempo time.Time
		wantErr   bool
	}{
		{
			name: "Tudo Ok",
			args: args{
				dataEmFormatoBrasileiro: "15/04/1996",
			},
			wantTempo: time.Date(1996, time.April, 15, 0, 0, 0, 0, utc),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTempo, err := DataFormatoBrasileiroParaTime(tt.args.dataEmFormatoBrasileiro)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataFormatoBrasileiroParaTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTempo, tt.wantTempo) {
				t.Errorf("DataFormatoBrasileiroParaTime() = %v, want %v", gotTempo, tt.wantTempo)
			}
		})
	}
}
