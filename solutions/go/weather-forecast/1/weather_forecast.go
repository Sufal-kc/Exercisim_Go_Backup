/*Package weather is the package responsible for providing the details 
about the weather of the Goblinocus country. */
package weather

var (
//CurrentCondition stores the currect weather conidtion in string form.
	CurrentCondition string
// CurrentLocation stores which location we are referring to.
	CurrentLocation  string
)

/*Forecast function is used to forecast the weather condition of the country
based on CurrentLocation and CurrentCondition of the country. */
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
