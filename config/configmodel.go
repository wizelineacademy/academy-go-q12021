package configz

type Configuration struct {
	AppName string 	`mapstructure:"application_name"	validate:"required"`
	AppHost string 	`mapstructure:"application_host" 	validate:"require"`
	CsvName string 	`mapstructure:"csv_name" 			validate:"required,hostname_port"`
}
