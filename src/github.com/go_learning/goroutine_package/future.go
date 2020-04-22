package goroutine_package

//所谓 Futures 就是指：有时候在你使用某一个值之前需要先对其进行计算。这种情况下，你就可以在另一个处理器上进行该值的计算，到使用时，该值就已经计算完毕了。

/*func InverseFuture(a Matrix) {
	future := make(chan Matrix)
	go func() {
		future <- Inverse(a)
	}()
	return future
}

func InverseProduct(a Matrix, b Matrix) {
	a_inv_future := InverseFuture(a)   // start as a goroutine
	b_inv_future := InverseFuture(b)   // start as a goroutine
	a_inv := <-a_inv_future
	b_inv := <-b_inv_future
	return Product(a_inv, b_inv)
}

*/
