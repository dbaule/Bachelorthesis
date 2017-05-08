package main

// ViewController struct that holds the UIComponents for the current ViewController
type ViewController struct {
	BuildVersion string        `json:"build_version" bson:"build_version"`
	Name         string        `json:"name" bson:"name"`
	UIComponents []UIComponent `json:"ui_components" bson:"ui_components"`
}

// ViewControllers a collection of ViewController's
type ViewControllers []ViewController
