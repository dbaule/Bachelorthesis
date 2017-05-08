//
//  UIComponentModel.swift
//  DummyApp
//
//  Created by Danny Baule on 15.11.16.
//  Copyright © 2016 Danny Baule. All rights reserved.
//

import Foundation

/// Represents a TransferableUIComponent that can be send to the server
struct UIComponentModel {
    
    // MARK: - UIComponent properties
    
    /// the type
    var ComponentType = ""
    
    /// the name that is assigned to it
    var Name = ""
    
    /// the relative x-Coordinate
    var XCoor : Double = 0.0
    
    /// the relative y-Coordinate
    var YCoor : Double = 0.0
    
    /// the height
    var Height : Double = 0.0
    
    /// the width
    var Width : Double = 0.0
    
    /// the subview components
    var SubviewComponents = [UIComponentModel]()
    
    // MARK: - UILabel properties
    
    /// the text of the label
    var UILabelText = ""
    
    /// the width of the text in the label
    var UILabelTextWidth : Double = 0.0
    
    /// the pointsize of the labels font
    var UILabelFontSize : Double = 0.0
    
    /// the font used for the label
    var UILabelFont = ""
    
    // MARK: - UIImageView properties
    
    /// the height of the image
    var UIImageViewImageHeight : Double = 0.0
    
    /// the width of the image
    var UIImageViewImageWidth : Double = 0.0
    
    /// the scale of the image
    var UIImageViewImageScale : Double = 0.0
    
    /// the orientation of the image
    var UIImageViewImageOrientation : Int = -1
    
    // MARK: - JSON representation
    
    /// the jsonRepresentation
    var jsonRepresentation : String {
        var result =  "{\"component_type\":\"\(ComponentType)\",\"name\":\"\(Name)\",\"x_coor\": \(XCoor),\"y_coor\": \(YCoor),\"height\": \(Height),\"width\": \(Width),\"uilabel_text\": \"\(UILabelText)\",\"uilabel_fontsize\": \(UILabelFontSize),\"uilabel_font\": \"\(UILabelFont)\",\"uilabel_textwidth\": \(UILabelTextWidth),\"uiimageview_imageheight\": \(UIImageViewImageHeight),\"uiimageview_imagewidth\": \(UIImageViewImageWidth),\"uiimageview_imagescale\": \(UIImageViewImageScale),\"uiimageview_imageorientation\": \(UIImageViewImageOrientation),\"subview_components\":["
        
        var x = 0
        
        for comp in SubviewComponents{
            if x == 0{
                result.append(comp.jsonRepresentation)
                x = 1
            }
            else{
                result.append("," + comp.jsonRepresentation)
            }
        }
        result.append("]}")
        return result
    }
}
