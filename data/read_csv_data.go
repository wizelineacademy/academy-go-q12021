package data

import (
    "os"
    "fmt"
    "encoding/csv"

    "pokedex/model/errors"
)

type CsvSource string

/**
 * Reads CSV contents from CsvSrouce
 */
func (source CsvSource) InitCsvSource() ([][]string, error) {
    file, err := os.Open(string(source))
    defer file.Close()

    if err != nil {
        fmt.Println("Unable to open file, err: ", err)
        data_src_err := errors.DataSourceIOError{ErrMsg: err}
        return [][]string{}, data_src_err
    }

    rows, err := csv.NewReader(file).ReadAll()
    if err != nil {
        fmt.Println("Error happened: ", err)
        data_src_err := errors.DataSourceIOError{ErrMsg: err}
        return [][]string{}, data_src_err
    }

    return rows, nil
}
