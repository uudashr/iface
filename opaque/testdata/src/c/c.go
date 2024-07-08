package donothing

func run(pass string) (interface{}, error) {
	inspect(func(n string) bool {
		return true
	})

	return nil, nil
}

func inspect(fn func(n string) bool) {

}
