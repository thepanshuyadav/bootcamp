package main

import (
	"ecommerce/handler"
	"ecommerce/infrastructure/persistence"
	"ecommerce/routes"
	"errors"
)

func main() {
	repos, err := persistence.NewRepositories("mysql", "root", "password@123", "localhost", "ecommerce", 3306)
	if err != nil {
		panic(err)
	}
	defer repos.Close()

	router := routes.SetUpRouter(
		handler.NewCustomerService(repos.CustomerRepo),
		handler.NewRetailerService(repos.RetailerRepo),
		handler.NewProductService(repos.ProductRepo, repos.RetailerRepo),
		handler.NewOrderService(repos.OrderRepo, repos.CustomerRepo, repos.ProductRepo),
	)
	if err = router.Run(); err != nil {
		panic(errors.New("Unable to configure router"))
	}

}
