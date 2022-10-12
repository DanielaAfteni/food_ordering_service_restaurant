# food_ordering_service_restaurant

This is the food_ordering_service_restaurant. It is related to the dining_hall_restaurant repository.

## Food ordering service app with Docker (used here docker compose)

It is required to introduce in Terminal:

```bash
$ docker compose up --build
```
## Run the app in the Terminal

Firstly switch the comment URL to: `"http://localhost"`

Then to run in the Terminal:

```bash
$ go run .
```
## Try it by yourself

Pay attention at the order of running, because everytime the kitchen_restaurant is running first, then food_ordering_service_restaurant, then dining_hall_restaurant, and the last client_restaurant.