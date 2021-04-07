package constant

// CsvMaxRetryVarName is the name used for the environment variable with the value for the maximun number of retries to append data to CSV file
const CsvMaxRetryVarName = "CSV_W_MAX_RETRY"

// CsvTimeRetryVarName is the name used for the environment variable with the value for the time (seconds) to wait between retries
const CsvTimeRetryVarName = "CSV_TIME_RETRY"

// EnvironmentVarName is the name used for the environment variable with the value for the current environment
const EnvironmentVarName = "ENVIRONMENT"

// PokemonServiceVarName is the name used for the environment variable with the value for the pokemon service URL
const PokemonServiceVarName = "POKEMON_SERVICE_URL"

// PokemonSourceVarName is the name used for the environment variable with the value for the pokemon source data
const PokemonSourceVarName = "POKEMON_SOURCE"

// ServerPortVarName is the name used for the environment variable with the value for the server port
const ServerPortVarName = "SERVER_PORT"

// DefaultEnvironment contains the default value when the env var 'ENVIRONMENT' is not present
const DefaultEnvironment = "test"

// DefaultPokemonService contains the default value when the env var 'POKEMON_SERVICE_URL' is not present
const DefaultPokemonService = "https://pokeapi.co/api/v2/pokemon"

// DefaultServerPort contains the default value when the env var 'SERVER_PORT' is not present
const DefaultServerPort = "8080"

// DefaultMaxRetries contains the default value when the env var 'CSV_W_MAX_RETRY' is not present
const DefaultMaxRetries = "10"

// DefaultTimeRetries contains the default value when the env var 'CSV_TIME_RETRY' is not present
const DefaultTimeRetries = "2"
