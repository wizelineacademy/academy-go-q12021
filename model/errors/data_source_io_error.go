package errors

import (
    "fmt"
)

type DataSourceIOError struct {
    ErrMsg error
}

func (e DataSourceIOError) Error() string {
    return fmt.Sprintf("Error while manipulating data source")
}

func (e DataSourceIOError) Unwrap() error {
    return e.ErrMsg
}
