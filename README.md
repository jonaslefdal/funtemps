# Funtemps Converter and Funfacts

| Command flag | Description                                                                        |
| ------------ | ---------------------------------------------------------------------------------- |
| `-C float` | Temperature in degrees Celsius |
| `-F float`   | Temperature in degrees Fahrenheit  |  
| `-K float` | Temperature in Kelvin  |
| `-out string` | Calculate temperature in C - celsius, F - farhenheit, K- Kelvin (default "C") |
| `-funfacts string`  | "Funfacts" om sun - Solen, Luna - Månen og terra - Jorden (default "sun" |
| `-T string` | Temperature unit for funfacts (default "C") |

<br>

# How to use 

## Example convertion command: 

<br>

### <p style="color:#cfcecc">Input: 'go run main.go -F 32 -out C'</p>
### <p style="color:#ADEFD1"> Output: '32°F is equal to 0°C'</p><br>

### <p style="color:#cfcecc">Input: 'go run main.go -C 423 -out K'</p>
### <p style="color:#ADEFD1"> Output: '423°C is equal to 696.15°K'</p>

<br>

## Example funfact command: 

<br>

### <p style="color:#cfcecc">Input: 'go run main.go -funfacts sun -T C'</p>
### <p style="color:#ADEFD1"> Output: 'Temperatur i Solens kjerne: 15000000.00°C'</p>
### <p style="color:#ADEFD1"> Output: 'Temperatur på ytre lag av Solen: 5504.85°C'</p>
