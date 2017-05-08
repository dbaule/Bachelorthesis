//
//  UIComponentConverter.swift
//  DummyApp
//
//  Created by Danny Baule on 15.11.16.
//  Copyright © 2016 Danny Baule. All rights reserved.
//

import Foundation
import UIKit

/// Implements functionality to convert UIComponents into transferable UIComponents
class UIComponentConverter{

    /// Transfers a UIView object into a UIComponentModel
    ///
    /// - Parameter view: the View object to transfer to a UIComponentModel
    /// - Returns: the newly created UIComponentModel
    func asTransferableObject(view: UIView) -> UIComponentModel{
    
        var comp = UIComponentModel()
        
        // writes the UIComponent data to the TransferableUIComponent
        comp.Name = view.uniqueIdentifer() ?? ""
        comp.ComponentType = type(of: view).description()
        comp.XCoor = Double(view.frame.origin.x)
        comp.YCoor = Double(view.frame.origin.y)
        comp.Height = Double(view.frame.height)
        comp.Width = Double(view.frame.width)
        
        for component in view.subviews{
            let compAsTransferable = asTransferableObject(view: component)
            comp.SubviewComponents.append(compAsTransferable)
        }
        
        // writes the UILabel or UIImageView data to the component
        if comp.ComponentType.contains("Label"){
            if let text = (view as! UILabel).text{
            comp.UILabelText = text
            }
            //comp.UILabelText = (view as! UILabel).text!
            comp.UILabelTextWidth = evaluateStringWidth(textToEvaluate: comp.UILabelText)
            comp.UILabelFont = (view as! UILabel).font.fontName
            comp.UILabelFontSize = Double((view as! UILabel).font.pointSize)
            
        } else if comp.ComponentType.contains("UIImageView"){
            
            if let image = (view as! UIImageView).image {
                comp.UIImageViewImageHeight = Double(image.size.height)
                comp.UIImageViewImageWidth = Double(image.size.width)
                comp.UIImageViewImageScale = Double(image.scale)
                comp.UIImageViewImageOrientation = image.imageOrientation.rawValue
            }
        }
        
        return comp
    }
    
    /// Transfers a UIViewController into a ViewControllerModel
    ///
    /// - Parameter viewController: the ViewController to transfer to a ViewControllerModel
    /// - Returns: the newly created ViewControllerModel
    func asTransferableObject(viewController: UIViewController) -> ViewControllerModel{
        
        // extracts the Build-Version
        let buildVersion = Bundle.main.infoDictionary?["CFBundleVersion"] as? String
        
        // gets the ViewControllers name if it has one, otherwise just writes the viewController
        let vcName = viewController.title ?? uniqueVCName(viewcontroller: viewController)//"\(viewController)"
        
        var vc = ViewControllerModel()
        
        // writes the gathered data to the ViewControllerModel
        vc.BuildVersion = buildVersion!
        vc.Name = vcName
        
        return vc
    }
    
    func uniqueVCName(viewcontroller: UIViewController) -> String{
        var uniqueIdentifier = ""
        
        if let par = viewcontroller.parent{
            uniqueIdentifier = type(of: par).description()
            uniqueIdentifier.append((par.childViewControllers.index(of: viewcontroller)?.description)!)
        }else{
            uniqueIdentifier = "Root"
        }
        
        uniqueIdentifier.append(type(of: viewcontroller).description())
        
        return uniqueIdentifier
    }
    
    // MARK: - Helper methods
    
    /// Calculates the width of the given string
    ///
    /// - Parameter textToEvaluate: the string to calculate the width for
    /// - Returns: a Double value representing the width of the given string
    private func evaluateStringWidth (textToEvaluate: String) -> Double{
        let font = UIFont.systemFont(ofSize: UIFont.systemFontSize)
        let attributes = NSDictionary(object: font, forKey:NSFontAttributeName as NSCopying)
        let sizeOfText = textToEvaluate.size(attributes: (attributes as! [String : AnyObject]))
        
        return Double(sizeOfText.width)
    }
    
}
