package errs

import (
	"fmt"
)

type EmptyDataError string

func (e EmptyDataError) Error() string {
	return fmt.Sprintf("There are not any %vs", string(e))
}
