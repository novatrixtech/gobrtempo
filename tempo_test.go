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

func TestDataFormatoMySQLParaTime(t *testing.T) {
	utc, _ := time.LoadLocation("America/Sao_Paulo")

	type args struct {
		dataEmFormatoMySQL string
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
				dataEmFormatoMySQL: "1996-04-15 00:00:00",
			},
			wantTempo: time.Date(1996, time.April, 15, 0, 0, 0, 0, utc),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTempo, err := DataFormatoMySQLParaTime(tt.args.dataEmFormatoMySQL)
			if (err != nil) != tt.wantErr {
				t.Errorf("DataFormatoMySQLParaTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTempo, tt.wantTempo) {
				t.Errorf("DataFormatoMySQLParaTime() = %v, want %v", gotTempo, tt.wantTempo)
			}
		})
	}
}

func TestTimeParaDataFormatoBrasileiro(t *testing.T) {
	utc, _ := time.LoadLocation("America/Sao_Paulo")

	type args struct {
		tempo time.Time
	}
	tests := []struct {
		name                        string
		args                        args
		wantDataEmFormatoBrasileiro string
		wantErr                     bool
	}{
		{
			name: "Tudo Ok",
			args: args{
				tempo: time.Date(1996, time.April, 15, 0, 0, 0, 0, utc),
			},
			wantDataEmFormatoBrasileiro: "15/04/1996",
			wantErr:                     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDataEmFormatoBrasileiro, err := TimeParaDataFormatoBrasileiro(tt.args.tempo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeParaDataFormatoBrasileiro() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDataEmFormatoBrasileiro != tt.wantDataEmFormatoBrasileiro {
				t.Errorf("TimeParaDataFormatoBrasileiro() = %v, want %v", gotDataEmFormatoBrasileiro, tt.wantDataEmFormatoBrasileiro)
			}
		})
	}
}

func TestTimeParaHoraFormatoBrasileiro(t *testing.T) {
	utc, _ := time.LoadLocation("America/Sao_Paulo")

	type args struct {
		tempo time.Time
	}
	tests := []struct {
		name                        string
		args                        args
		wantHoraEmFormatoBrasileiro string
		wantErr                     bool
	}{
		{
			name: "Tudo Ok",
			args: args{
				tempo: time.Date(1996, time.April, 15, 15, 04, 28, 0, utc),
			},
			wantHoraEmFormatoBrasileiro: "15:04:28",
			wantErr:                     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHoraEmFormatoBrasileiro, err := TimeParaHoraFormatoBrasileiro(tt.args.tempo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeParaHoraFormatoBrasileiro() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHoraEmFormatoBrasileiro != tt.wantHoraEmFormatoBrasileiro {
				t.Errorf("TimeParaHoraFormatoBrasileiro() = %v, want %v", gotHoraEmFormatoBrasileiro, tt.wantHoraEmFormatoBrasileiro)
			}
		})
	}
}

func TestTimeParaDataHoraFormatoBrasileiro(t *testing.T) {
	utc, _ := time.LoadLocation("America/Sao_Paulo")

	type args struct {
		tempo time.Time
	}
	tests := []struct {
		name                        string
		args                        args
		wantDataEmFormatoBrasileiro string
		wantErr                     bool
	}{
		{
			name: "Tudo Ok",
			args: args{
				tempo: time.Date(1996, time.April, 15, 15, 04, 28, 0, utc),
			},
			wantDataEmFormatoBrasileiro: "15/04/1996 15:04:28",
			wantErr:                     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDataEmFormatoBrasileiro, err := TimeParaDataHoraFormatoBrasileiro(tt.args.tempo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeParaDataHoraFormatoBrasileiro() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDataEmFormatoBrasileiro != tt.wantDataEmFormatoBrasileiro {
				t.Errorf("TimeParaDataHoraFormatoBrasileiro() = %v, want %v", gotDataEmFormatoBrasileiro, tt.wantDataEmFormatoBrasileiro)
			}
		})
	}
}

func TestTimeParaDataHoraFormatoMySQL(t *testing.T) {
	utc, _ := time.LoadLocation("America/Sao_Paulo")

	type args struct {
		tempo time.Time
	}
	tests := []struct {
		name                   string
		args                   args
		wantDataEmFormatoMySQL string
		wantErr                bool
	}{
		{
			name: "Tudo Ok",
			args: args{
				tempo: time.Date(1996, time.April, 15, 15, 04, 28, 0, utc),
			},
			wantDataEmFormatoMySQL: "1996-04-15 15:04:28",
			wantErr:                false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDataEmFormatoMySQL, err := TimeParaDataHoraFormatoMySQL(tt.args.tempo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeParaDataHoraFormatoMySQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDataEmFormatoMySQL != tt.wantDataEmFormatoMySQL {
				t.Errorf("TimeParaDataHoraFormatoMySQL() = %v, want %v", gotDataEmFormatoMySQL, tt.wantDataEmFormatoMySQL)
			}
		})
	}
}

// TODO: verificar se realmete a função deve ter esse comportamento pois esta igual a da TimeParaDataHoraFormatoMySQL
func TestTimeParaDataFormatoMySQL(t *testing.T) {
	utc, _ := time.LoadLocation("America/Sao_Paulo")

	type args struct {
		tempo time.Time
	}
	tests := []struct {
		name                   string
		args                   args
		wantDataEmFormatoMySQL string
		wantErr                bool
	}{
		{
			name: "Tudo Ok",
			args: args{
				tempo: time.Date(1996, time.April, 15, 0, 0, 0, 0, utc),
			},
			wantDataEmFormatoMySQL: "1996-04-15 00:00:00",
			wantErr:                false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDataEmFormatoMySQL, err := TimeParaDataFormatoMySQL(tt.args.tempo)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeParaDataFormatoMySQL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDataEmFormatoMySQL != tt.wantDataEmFormatoMySQL {
				t.Errorf("TimeParaDataFormatoMySQL() = %v, want %v", gotDataEmFormatoMySQL, tt.wantDataEmFormatoMySQL)
			}
		})
	}
}
