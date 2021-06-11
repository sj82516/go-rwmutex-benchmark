package mutex_benchmark

func FakeRead(){
	for i:=0; i< 1e6; i++ {}
}

func FakeWrite(){
	for j:=0; j <5e6; j++{}
}
