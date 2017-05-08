//
//  UIViewControllerExtension.swift
//  DummyApp
//
//  Created by Danny Baule on 28.09.16.
//  Copyright © 2016 Danny Baule. All rights reserved.
//

import Foundation
import UIKit

/// Defines further functionality to the UIViewController class
extension UIViewController {
    
    /// overrides the initialize method
    open override class func initialize(){
        
        if self !== UIViewController.self{
            return
        }
        
        // Method Swizzling with the viewDidAppear method
        let _: () = {
            let originalSelector = #selector(UIViewController.viewDidAppear(_:))
            let swizzledSelector = #selector(UIViewController.CK_viewDidAppear(animated:))
            
            let originalMethod = class_getInstanceMethod(self, originalSelector)
            let swizzledMethod = class_getInstanceMethod(self, swizzledSelector)
            
            let didAddMethod = class_addMethod(self, originalSelector, method_getImplementation(swizzledMethod), method_getTypeEncoding(swizzledMethod))
            
            if didAddMethod {
                class_replaceMethod(self, swizzledSelector, method_getImplementation(originalMethod), method_getTypeEncoding(originalMethod))
            } else {
                method_exchangeImplementations(originalMethod, swizzledMethod)
            }
            
            // Swizzling the viewWillDisappear method, to track down the back button action
            let originalSelector2 = #selector(UIViewController.viewWillDisappear(_:))
            let swizzledSelector2 = #selector(UIViewController.CK_viewWillDisappear(animated:))
            
            let originalMethod2 = class_getInstanceMethod(self, originalSelector2)
            let swizzledMethod2 = class_getInstanceMethod(self, swizzledSelector2)
            
            let didAddMethod2 = class_addMethod(self, originalSelector2, method_getImplementation(swizzledMethod2), method_getTypeEncoding(swizzledMethod2))
            
            if didAddMethod2 {
                class_replaceMethod(self, swizzledSelector2, method_getImplementation(originalMethod2), method_getTypeEncoding(originalMethod2))
            } else {
                method_exchangeImplementations(originalMethod2, swizzledMethod2)
            }
        }()
    }
    
    // MARK: - Method Swizzling

    /// adds additional functionality to the viewWillAppear method
    ///
    /// - Parameter animated: WillAppear animated or not
    func CK_viewDidAppear(animated: Bool) {
        self.CK_viewDidAppear(animated: animated)
   
        // converts the current ViewController to a Transferable ViewController
        var vc = UIComponentConverter().asTransferableObject(viewController: self)
        
        // prints the properties
        print("viewDidAppear: \(vc.Name) on Build-Version: \(vc.BuildVersion) \n")
        
        // loop over all subviews of the VC
        for x in self.view.subviews{
            
            //TODO: schöner machen
            if !x.isHidden {
                // converts the UIComponent to a UIComponentModel
                let comp = UIComponentConverter().asTransferableObject(view: x)
            
                let componentType = "Component : \(comp.ComponentType) \n"
                let height = "Height : \(comp.Height) \n"
                let width = "Width : \(comp.Width) \n"
                let xcoor = "X : \(comp.XCoor) \n"
                let ycoor = "Y : \(comp.YCoor) \n"
            
                // prints the properties
                print(
                    componentType +
                    height +
                    width +
                    xcoor +
                    ycoor 
                )
                
                // appends the UIComponent to the ViewControllerModel
                vc.UIComponents.append(comp)
            }
        }
        
        // saves an image of the current ViewController and appends its directory path to the VCModel
        let path = createImageFolder(buildVersion: vc.BuildVersion)
        if let image = screenShotMethod(vc: vc){
            saveImageToFileSystem(image: image, vc: vc, path: path)
        }
        
        // posts the VC to the server
        ServerInteractionHandler().postVCtoServer(vc: vc)
        // writes to the tracking array
        TrackingManager.sharedInstance.trackingArray.append(vc)
    }
    
    /// adds additional functionality to the viewWillDisappear method in providing a way to track down the back button event, by listening to the isMovingFromParentViewController property
    ///
    /// - Parameter animated: Disappear animated or not
    func CK_viewWillDisappear(animated: Bool){
        // call to the original method
        self.CK_viewWillDisappear(animated: animated)
        
        // if isMovingFromParentViewController, then back button was pressed
        if self.isMovingFromParentViewController {
            TrackingManager.sharedInstance.trackingArray.append("backButton:")
        }
    }
    
    
    func screenShotMethod(vc: ViewControllerModel) -> UIImage? {
        //create screenshot if possible
        if let layer = UIApplication.shared.keyWindow?.layer {
            let scale = UIScreen.main.scale
            
            // fixed text properties
            let textColor = UIColor.black
            let textFont = UIFont(name: "Helvetica Bold", size: 8)
            
            UIGraphicsBeginImageContextWithOptions(layer.frame.size, false, scale)
            if let context = UIGraphicsGetCurrentContext() {
                layer.render(in: context)
                
                //create Rectangles with options
                for comp in vc.UIComponents{
                    // attributes for rectangles
                    context.setLineWidth(5.0)
                    context.setStrokeColor(UIColor.orange.cgColor)
                    context.setFillColor(UIColor.clear.cgColor)
                    
                    // attributes for comp name
                    let fontAttributes = [
                        NSFontAttributeName: textFont ?? UIFont.systemFont(ofSize: 8),
                        NSForegroundColorAttributeName: textColor,
                        ] as [String : Any]
                    
                    // rectangle creation and addition
                    let rectangle = CGRect(x:comp.XCoor, y:comp.YCoor, width:comp.Width, height:comp.Height)
                    context.addRect(rectangle)
                    context.drawPath(using: .fillStroke)
                    
                    // text addition
                    (comp.Name as NSString).draw(in: rectangle, withAttributes: fontAttributes)
                }
                let screenshot = UIGraphicsGetImageFromCurrentImageContext()
                UIGraphicsEndImageContext()
                return screenshot
            }
        }
        return nil
    }
    
    func saveImageToFileSystem(image:UIImage, vc:ViewControllerModel, path: String){
        //create imagePath
        var imagePath = path
        let imageName = "/\(vc.Name).png"
        imagePath = imagePath.appending(imageName)
        
        //create png image
        let data = UIImagePNGRepresentation(image)
        
        //save image
        var objcBool:ObjCBool = true
        let isExist = FileManager.default.fileExists(atPath: imagePath, isDirectory: &objcBool)
        
        // If the folder with the given path doesn't exist already, create it
        if isExist == false{
            FileManager.default.createFile(atPath: imagePath, contents: data, attributes: nil)
        }
    }
    
    func createImageFolder(buildVersion: String) -> String {
        var imagesDirectoryPath:String!
        
        // Get the Document directory path
        let documentDirectorPath:String = "/Users/Danny/Documents/Studium/Bachelorarbeit"
        let appName = Bundle.main.bundleIdentifier ?? ""
        //let appName = Bundle.main.infoDictionary![kCFBundleNameKey as String] as! String
        // Create a new path for the new images folder
        imagesDirectoryPath = documentDirectorPath.appending("/UICompareKit_Screenshots")
        imagesDirectoryPath = imagesDirectoryPath.appending("/\(appName)/\(buildVersion)")
        var objcBool:ObjCBool = true
        let isExist = FileManager.default.fileExists(atPath: imagesDirectoryPath, isDirectory: &objcBool)
        
        // If the folder with the given path doesn't exist already, create it
        if isExist == false{
            do{
                try FileManager.default.createDirectory(atPath: imagesDirectoryPath, withIntermediateDirectories: true, attributes: nil)
            }catch{
                print("Something went wrong while creating a new folder")
            }
        }
        return imagesDirectoryPath
    }
    
}
