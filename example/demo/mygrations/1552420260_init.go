package mygrations

func Mygration_1552420260_init_Up(dep interface{}) error {
	_, err := dep.(Deps).MasterDB.Begin()
	if err != nil {
		return err
	}
	// ...

	return nil
}

func Mygration_1552420260_init_Down(dep interface{}) error {
	_, err := dep.(Deps).MasterDB.Begin()
	if err != nil {
		return err
	}
	// ...

	return nil
}
