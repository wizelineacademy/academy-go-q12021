package constants

// Odd number string
const Odd = "odd"

// Even number string
const Even = "even"

// Maximum Size of an unsigned integer
const UintSize = 32 << (^uint(0) >> 32 & 1)

// Maximum Size of an integer
const MaxInt = 1<<(UintSize-1) - 1
