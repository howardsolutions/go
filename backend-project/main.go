package backendproject

import "howardsolutions/go/backend-project/internal/app"

func main() {
	app, err := app.NewApplication()

	if err != nil {
		panic(err)
	}

	app.Logger.Println("We are running our app!")
}
