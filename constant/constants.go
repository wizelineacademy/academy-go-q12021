package constant

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
