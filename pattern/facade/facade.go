package facade

type OrderFacade struct {
	waiter  Waiter
	kitchen Kitchen
}

func NewOrderFacade() *OrderFacade {
	return &OrderFacade{
		waiter:  Waiter{},
		kitchen: Kitchen{},
	}
}

func (o *OrderFacade) Order() int {
	o.waiter.TakeOrder()
	o.waiter.SendToKitchen()

	o.kitchen.Prepare()
	o.kitchen.CallWaiter()

	o.waiter.Serve()
	o.kitchen.WashDishes()

	println("Facade done")

	return 1
}

type Waiter struct {
}

func (w *Waiter) TakeOrder() {
	println("Waiter: Take order")
}

func (w *Waiter) Serve() {
	println("Waiter: Serve")
}

func (w *Waiter) SendToKitchen() {
	println("Waiter: Send to kitchen")
}

type Kitchen struct {
}

func (k *Kitchen) Prepare() {
	println("Kitchen: Prepare")
}

func (k *Kitchen) CallWaiter() {
	println("Kitchen: Call waiter")
}

func (k *Kitchen) WashDishes() {
	println("Kitchen: Wash dishes")
}
