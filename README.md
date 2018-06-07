# gobrtempo

Go (golang) lib to handle time/date in Brazilian format, timezone and standards

[![Build Status](https://travis-ci.org/novatrixtech/gobrtempo.svg?branch=master)](https://travis-ci.org/novatrixtech/gobrtempo) [![Go Report Card](https://goreportcard.com/badge/github.com/novatrixtech/gobrtempo)](https://goreportcard.com/report/github.com/novatrixtech/gobrtempo)

## Examples

```go
dataPrimeiroPagto, _ := gobrtempo.DataFormatoBrasileiroParaTime("29/04/2018")
dataDeHoje := gobrtempo.AgoraDataEHoraEmTime()
```
