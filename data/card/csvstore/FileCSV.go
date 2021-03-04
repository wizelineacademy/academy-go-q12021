package csvstore

import (
  "io/ioutil"
  "strconv"
  "strings"

  "github.com/spf13/viper"

  dataCard "github.com/gbrayhan/academy-go-q12021/data/card"
)

// FileCSV
type FileCSV struct {
  Error error
}

func (FileCSV) mapCSVFile() (data map[int]dataCard.Card, err error) {
  data = make(map[int]dataCard.Card)

  viper.SetConfigFile("config.json")
  _ = viper.ReadInConfig()
  fileName := viper.GetString("Databases.CSV.CardsYGOFile")
  contentBytes, err := ioutil.ReadFile(fileName)

  for _, line := range strings.Split(string(contentBytes), "\n") {
    var rows []string
    if line != "" {
      rows = strings.Split(line, ",")
    }

    if len(rows) != 0 {
      key, _ := strconv.Atoi(rows[0])
      level, _ := strconv.Atoi(rows[3])
      atk, _ := strconv.Atoi(rows[6])
      def, _ := strconv.Atoi(rows[7])

      if key != 0 {
        data[key] = dataCard.Card{
          ID:        key,
          Name:      rows[1],
          Type:      rows[2],
          Level:     level,
          Race:      rows[4],
          Attribute: rows[5],
          ATK:       atk,
          DEF:       def}
      }
    }
  }
  return
}

func (f *FileCSV) FindCardByID(id int) (card dataCard.Card, err error) {
  dataMap, err := f.mapCSVFile()
  card = dataMap[id]
  return
}

func (f *FileCSV) FindAllCards(cards *[]dataCard.Card, ) (err error) {
  dataMap, err := f.mapCSVFile()

  for _, card := range dataMap {
    *cards = append(*cards, card)
  }

  return
}
