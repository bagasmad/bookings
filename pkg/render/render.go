package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/bagasmad/bookings/pkg/config"
	"github.com/bagasmad/bookings/pkg/models"
)

// inevitably where we want certain kinds of data to be available to every page in every side, take template data that is passed by handler
// add data to it that we want to be available on every page of our website
func AddDefaultData(data *models.TemplateData) *models.TemplateData {
	return data
}

// we're going to add the templates dynamically
func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	var tc map[string]*template.Template
	if appConfig.UseCache {
		tc = appConfig.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	templateData = AddDefaultData(templateData)

	//execute buffer
	buf := new(bytes.Buffer)
	_ = t.Execute(buf, templateData)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//populate that with every possible templates and roots and partials
	//get all of the files .page.tmpl from the ./templates
	//Glob will search for that files that ends with .page.tmpl

	//the downside of this code is that it will associate all the layout to that specific page, let's say about.page.tmpl only use base.layout.tmpl
	//this code will return template set where about.page.tmpl is associated with base.layout.tmpl and other *.layout.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//range through the pages
	for _, page := range pages {
		//get name of the path like something.page.tmpl
		name := filepath.Base(page)
		//parse something.page.tmpl and store template as name
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//look for any layouts that exist in that directory
		matchLayout, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matchLayout) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, err

}

var appConfig *config.AppConfig

func AccessConfig(app *config.AppConfig) {
	appConfig = app
}

//previous version

// // template cache a package level variable
// var tc = make(map[string]*template.Template)

// // instead of reading the disk every single time, we want a data structure that we can store a parsed template to that data structure
// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//i want to check to see if we already have the template in our cache
// 	_, inMap := tc[t] //looking in the map tc
// 	//if not inMap then
// 	if !inMap {
// 		//need to create the template
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			//print the error
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have the template in the cache
// 		log.Println("using cached template")
// 	}

// 	//will get the template using the t string as a key
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		//print the error
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	//creating the slice of templates directory
// 	templates := []string{
// 		//first template is to get the template from when we access or hit the endpoint
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl", //second template is the base layout template that all the other template use
// 	}
// 	//parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache (map)
// 	tc[t] = tmpl
// 	return nil
// }
