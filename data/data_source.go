package data

/* TODO: Add more data sources here */
type Source interface {
    InitCsvSource() ([][]string, error)
}
