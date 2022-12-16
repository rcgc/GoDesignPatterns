package main

func main() {
	// Construct two DataListener observrs and
	// give each one a name

	listener1 := DataListener{
		Name: "Listener1",
	}
	listener2 := DataListener{
		Name: "Listener2",
	}

	// Create the DataSubject that the Listeners will observe
	subj := &DataSubject{}
	// TODO: Register each listener with the DataSubject
	subj.registerObserver(listener1)
	subj.registerObserver(listener2)

	// TODO: CHange the data in the DataSubject - this will cause the
	// onUpdate method of each listener to be called
	subj.changeItem("Monday!")
	subj.changeItem("Wednesday!")

	// TODO: try to unregister one of the observers
	subj.unregisterObserver(listener2)
	subj.changeItem("Friday!")
}