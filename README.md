# LEHACODE

LEHACODE is an encoding made by me. Only two people on this planet know the purpose of this project and why is it implemented like
that.

# About LEHACODE

This mehod of encoding uses 6 bits for a character, which means it only supports 64 different symbols: all english alphabet (uppercase and lowercase), all the numbers, underscore sign and the dollar sign (cus money). All other character leha_bytes (don't ask) will be encoded as 'NOT_SUPPORTED'.

# Usage

You can copy the LEHACODE.go file content to use the functions in your project, or you can use a shell that I wrote. You can run the shell in the directory with this repo with the command below.

`go run LEHACODE.go shell.go`

# Shell manual:

Nothing too serious, just two commands: encode and decode.
1. You can encode a string to LEHACODE with the encode command. You can also use the --full-string flag with this command to remove
the leha_byte separators.
2. Use decode to decode LEHACODE to a regular string. No flags are provided with that command.