//
//  UIApplicationExtension.swift
//  DummyApp
//
//  Created by Danny Baule on 25.11.16.
//  Copyright © 2016 Danny Baule. All rights reserved.
//

import Foundation
import UIKit

/// Defines further functionality to the UIApplication class
extension UIApplication{
   
    /// overrides the initialize method
    open override class func initialize() {
        
        if self !== UIApplication.self{
            return
        }

        let _: () = {
        // Method Swizzling with the sendAction method
        let originalSelector = #selector(UIApplication.sendAction(_:to:from:for:))
        let swizzledSelector = #selector(UIApplication.CK_sendAction(_:to:from:for:))
        
        let originalMethod = class_getInstanceMethod(self, originalSelector)
        let swizzledMethod = class_getInstanceMethod(self, swizzledSelector)
        
        let didAddMethod = class_addMethod(self, originalSelector, method_getImplementation(swizzledMethod), method_getTypeEncoding(swizzledMethod))
        
        if didAddMethod{
            class_replaceMethod(self, swizzledSelector, method_getImplementation(originalMethod), method_getTypeEncoding(originalMethod))
        } else{
            method_exchangeImplementations(originalMethod, swizzledMethod)
            }}()
    }
    
    // MARK: - Method Swizzling

    /// add additional functionality to the sendAction method in extracting the Action properties and sending it to the server as an event
    ///
    /// - Parameters:
    ///   - action: A selector identifying an action method. See the discussion for information on the permitted selector forms.
    ///   - target: The object to receive the action message. If target is nil, the app sends the message to the first responder, from whence it progresses up the responder chain until it is handled.
    ///   - sender: The object that is sending the action message. The default sender is the UIControl object that invokes this method.
    ///   - event: A UIEvent object that encapsulates information about the event originating the action message.
    func CK_sendAction(_ action: Selector, to target: AnyObject?, from sender: AnyObject?, for event: UIEvent?){
        self.CK_sendAction(action, to: target, from: sender, for: event)
        
        //TODO: transfer data to server, but which data will be enough
        print("action : \(action.description) \n")
        
        if let target = target {
            print("to target : \(target.description) \n")
        }
        
        if let sender = sender {
            print("from sender : \(sender.description) \n")
        }
        
        print("event : \(event.debugDescription) \n")
        
        let keyWindow = self.keyWindow
        let rootVC = keyWindow?.rootViewController
        print("\(rootVC?.description)")
        var vcNum = 0
        var vcs = [UIViewController]()
//        if let naviVC = rootVC as! UINavigationController?{
//        //let navVC = rootVC?.presentedViewController
//        //let bla = rootVC?.presentingViewController
//            vcs = naviVC.viewControllers
//            vcNum = vcs.count
//        }
//        else{
            vcs = (rootVC?.childViewControllers)!
            vcNum = vcs.count
        //}
        var prevVC = UIViewController()
        
        // like this it will get the previousVC, but the event doesnt raise on Back button
        // and will also get prevVC wrong when there is no viewController transition only a button click in e.g. View2VC
        if vcNum < 2{
            prevVC = vcs[vcNum - 1]
        } else {
            prevVC = vcs[vcNum - 2]
        }
        print("previous VC : \(prevVC.description)")
        
        TrackingManager.sharedInstance.trackingArray.append(action.description)
        
    }
}
