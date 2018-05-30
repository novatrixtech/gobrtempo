package gobrtempo

import "time"

//LocalizacaoBrasil representa a localizacao para data da capital do Brasil (Brasilia = Sao_Paulo pelo banco de dados IANA Time Zone)
//Detalhes: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
var LocalizacaoBrasil *time.Location

const (
	//FormatoBrasileiroData representa data no formato mais popular no Brasil
	FormatoBrasileiroData = "02/01/2006"
	//FormatoBrasileiroHora representa a hora no formato mais popular no Brasil
	FormatoBrasileiroHora = "15:04:05"
	//FormatoBrasileiroDataHora representa data e hora no formato mais popular no Brasil
	FormatoBrasileiroDataHora = "02/01/2006 15:04:05"
	//FormatoMySQLData representa data e hora no formato padrão do MySQL
	FormatoMySQLData = "2006-01-02 15:04:05"
	//FormatoMySQLSomenteData representa data no formato padrão do MySQL
	FormatoMySQLSomenteData = "2006-01-02"
)

func init() {
	LocalizacaoBrasil, _ = time.LoadLocation("America/Sao_Paulo")
}

//DataFormatoBrasileiroParaTime - de uma data em formato brasileiro (dd/mm/aaaa) para time do Go
func DataFormatoBrasileiroParaTime(dataEmFormatoBrasileiro string) (tempo time.Time, err error) {
	tempo, err = time.ParseInLocation(FormatoBrasileiroData, dataEmFormatoBrasileiro, LocalizacaoBrasil)
	return
}

//DataFormatoMySQLParaTime - de uma data em formato de data padrão do MySQL brasileiro (YYYY-MM-DD HH:MM:SS) para time do Go
func DataFormatoMySQLParaTime(dataEmFormatoMySQL string) (tempo time.Time, err error) {
	tempo, err = time.ParseInLocation(FormatoMySQLData, dataEmFormatoMySQL, LocalizacaoBrasil)
	if err != nil {
		return
	}
	return
}

//TimeParaDataFormatoBrasileiro - de um time em Go para o formato brasileiro (dd/mm/aaaa)
func TimeParaDataFormatoBrasileiro(tempo time.Time) (dataEmFormatoBrasileiro string) {
	dataEmFormatoBrasileiro = tempo.Format(FormatoBrasileiroData)
	return
}

//TimeParaHoraFormatoBrasileiro - de um time em Go para o formato brasileiro (dd/mm/aaaa)
func TimeParaHoraFormatoBrasileiro(tempo time.Time) (horaEmFormatoBrasileiro string) {
	horaEmFormatoBrasileiro = tempo.Format(FormatoBrasileiroHora)
	return
}

//TimeParaDataHoraFormatoBrasileiro - de um time em Go para o formato brasileiro (dd/mm/aaaa hh:mm:ss)
func TimeParaDataHoraFormatoBrasileiro(tempo time.Time) (dataEmFormatoBrasileiro string) {
	dataEmFormatoBrasileiro = tempo.Format(FormatoBrasileiroDataHora)
	return
}

//TimeParaDataHoraFormatoMySQL - de um time em Go para o formato MySQL (aaaa-mm-dd hh:mm:ss)
func TimeParaDataHoraFormatoMySQL(tempo time.Time) (dataEmFormatoMySQL string) {
	dataEmFormatoMySQL = tempo.Format(FormatoMySQLData)
	return
}

//TimeParaDataFormatoMySQL - de um time em Go para o formato MySQL (aaaa-mm-dd)
func TimeParaDataFormatoMySQL(tempo time.Time) (dataEmFormatoMySQL string) {
	dataEmFormatoMySQL = tempo.Format(FormatoMySQLSomenteData)
	return
}

//Agora - obtem a hora atual do sistema no time-zone de SaoPaulo/Brasilia e retorna a hora
func Agora() (tempo string) {
	agora := time.Now().In(LocalizacaoBrasil)
	tempo = TimeParaHoraFormatoBrasileiro(agora)
	return
}

//Hoje - obtem a data atual do sistema no time-zone de SaoPaulo/Brasilia e retorna a data
func Hoje() (tempo string) {
	agora := time.Now().In(LocalizacaoBrasil)
	tempo = TimeParaDataFormatoBrasileiro(agora)
	return
}

//AgoraDataEHora - obtem a data e hora atual do sistema no time-zone de SaoPaulo/Brasilia e retorna data e hora
func AgoraDataEHora() (tempo string) {
	agora := time.Now().In(LocalizacaoBrasil)
	tempo = TimeParaDataHoraFormatoBrasileiro(agora)
	return
}

//AgoraDataEHoraEmTime - obtem a data e hora atual do sistema no time-zone de SaoPaulo/Brasilia e retorna o time
func AgoraDataEHoraEmTime() (tempo time.Time) {
	tempo = time.Now().In(LocalizacaoBrasil)
	return
}

//AgoraDataEHoraParaFormatoMySQL - obtem a data e hora atual do sistema no time-zone de SaoPaulo/Brasilia e retorna data e hora no formato MySQL
func AgoraDataEHoraParaFormatoMySQL() (tempo string, err error) {
	agora := time.Now().In(LocalizacaoBrasil)
	tempo = TimeParaDataHoraFormatoMySQL(agora)
	return
}

//ConverteStringDataEmFormatoBrasileiroParaMySQL converte uma string de data 25/12/2010 para formato MySQL 2010-12-25
func ConverteStringDataEmFormatoBrasileiroParaMySQL(data string) (dataSQL string, err error) {
	err = nil
	tmpData, err := DataFormatoBrasileiroParaTime(data)
	if err != nil {
		return
	}
	dataSQL = TimeParaDataHoraFormatoMySQL(tmpData)
	return
}
