package datastore

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alexis-aguirre/golang-bootcamp-2020/infraestructure/services"
)

type Logger struct { //Implements interface Logger and Service
	FilePath string
	File     *os.File
	status   int
}

func InitializeLogger(filePath string) *Logger {
	logger := &Logger{FilePath: filePath}
	logger.Start()
	return logger
}

func (lo *Logger) Start() error {
	file, err := os.OpenFile(lo.FilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Cannot open log file")
		return err
	}
	lo.File = file
	lo.status = services.STATUSOK
	return nil
}
func (lo *Logger) Stop() error {
	err := lo.File.Close()
	if err != nil {
		log.Println("Cannot close log file")
	}
	return nil
}
func (lo *Logger) Status() int {
	return lo.status
}

func (lo *Logger) Append(record string) error {
	writter := bufio.NewWriter(lo.File)
	_, err := writter.WriteString(time.Now().String() + "," + record + "\n")
	if err != nil {
		log.Println("Cannot add the record to the log", err)
		return err
	}
	err = writter.Flush()
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

func (lo *Logger) Get() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(lo.File)
	for scanner.Scan() {
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}

		text := scanner.Text()
		lines = append(lines, text)
	}
	return lines, nil
}

func UserLogToJson(records []string) string {
	for i := 0; i < len(records); i++ {
		records[i] = fmt.Sprintf("{%s},", records[i])
	}
	jointRecords := strings.Join(records, "\n")
	return fmt.Sprintf(`{
		"records": [
%s
		]
		}`, jointRecords)

}
