package goroutine_package

type Data struct {
}

//串行化处理数据
func SerialProcessData(in <-chan *Data, out <-chan *Data) {

	//PreprocessData，ProcessStepA，ProcessStepB，PostProcessData
	//分别代表了一个流程的四个步骤
	/*for data := range in {

		tmpA := PreprocessData(data)

		tmpB := ProcessStepA(tmpA)

		tmpC := ProcessStepB(tmpB)

		out <- PostProcessData(tmpC)

	}*/

}

//并行处理数据
func ParallelProcessData(in, out chan *Data) {
	/*// make channels:

	preOut := make(chan *Data, 100)

	stepAOut := make(chan *Data, 100)

	stepBOut := make(chan *Data, 100)

	stepCOut := make(chan *Data, 100)

	// start parallel computations:

	go PreprocessData(in, preOut)

	go ProcessStepA(preOut, stepAOut)

	go ProcessStepB(stepAOut, stepBOut)

	go ProcessStepC(stepBOut, stepCOut)

	go PostProcessData(stepCOut, out)*/
}
