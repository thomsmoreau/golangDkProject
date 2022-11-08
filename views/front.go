package views

import (
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents-heroicons/solid"
	. "github.com/maragudk/gomponents/html"
)

// It's a form that makes an "HTTP POST" request to "/newsletter/signup"
var signupForm = FormEl(
	Action("/newsletter/signup"), Method("post"), Class("flex items-center max-w-md"),
	Label(For("email"), Class("sr-only"), g.Text("Email")),
	Div(Class("relative rounded-md shadow-sm flex-grow"),
		Div(Class("absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"),
			solid.Mail(Class("h-5 w-5 text-gray-400")),
		),
		// There's a single required input field of both name and type email
		Input(
			Type("email"), Name("email"), ID("email"), AutoComplete("email"), Required(), Placeholder("me@example.com"), TabIndex("1"),
			Class("focus:ring-gray-500 focus:border-gray-500 block w-full pl-10 text-sm border-gray-300 rounded-md"),
		),
	),
	Button(
		Type("submit"), g.Text("Sign up"),
		Class("ml-3 inline-flex items-center px-4 py-2 border border-gray-300 shadow-sm text-sm font-medium rounded-md text-gray-700 bg-white hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 flex-none"),
	),
)

func FrontPage() g.Node {
	return Page(
		"Home",
		"/",
		// Our body for front page
		H1(g.Text(`PathFinder - Home page`)),
		P(g.Text(`Do you have problems finding your way ? We also have sometimes`)),
		P(g.Raw(`Then we created the <em>PathFinder</em> app, and now we don't !`)),
		H2(g.Text(`Do tou want to know more about future updates ?`)),
		P(g.Text(`Sign up to our newsletter now !`)),

		// Form to register to newsletter
		signupForm,
	)
}
