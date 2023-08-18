package models

// Create New type to hold data that we will display inside our app, this is far more efficient than
// just sending data with variables in func
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	//when we're not sure about the type of data, we will call it interface
	Data map[string]interface{}
	//Cross site request forgery token
	CSRFToken string
	//Flash message
	Flash string
	//warning message
	Warning string
	//Error message
	Error string
}
