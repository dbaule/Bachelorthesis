package main

// UIComponent struct that holds the UIComponent attributes
type UIComponent struct {
	//ID            bson.ObjectId `json:"id" bson:"id"`
	ComponentType               string        `json:"component_type" bson:"component_type"`
	Name                        string        `json:"name" bson:"name"`
	XCoor                       float64       `json:"x_coor" bson:"x_coor"`
	YCoor                       float64       `json:"y_coor" bson:"y_coor"`
	Height                      float64       `json:"height" bson:"height"`
	Width                       float64       `json:"width" bson:"width"`
	UILabelText                 string        `json:"uilabel_text" bson:"uilabel_text"`
	UILabelTextWidth            float64       `json:"uilabel_textwidth" bson:"uilabel_textwidth"`
	UILabelFontSize             float64       `json:"uilabel_fontsize" bson:"uilabel_fontsize"`
	UILabelFont                 string        `json:"uilabel_font" bson:"uilabel_font"`
	UIImageViewImageHeight      float64       `json:"uiimageview_imageheight" bson:"uiimageview_imageheight"`
	UIImageViewImageWidth       float64       `json:"uiimageview_imagewidth" bson:"uiimageview_imagewidth"`
	UIImageViewImageScale       float64       `json:"uiimageview_imagescale" bson:"uiimageview_imagescale"`
	UIImageViewImageOrientation int           `json:"uiimageview_imageorientation" bson:"uiimageview_imageorientation"`
	SubviewComponents           []UIComponent `json:"subview_components" bson:"subview_components"`
}

// UIComponents a collection of UIComponent's
type UIComponents []UIComponent
