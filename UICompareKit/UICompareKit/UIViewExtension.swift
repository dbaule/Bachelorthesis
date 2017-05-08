//
//  UIViewExtension.swift
//  UICompareKit
//
//  Created by Danny Baule on 11.01.17.
//  Copyright © 2017 Danny Baule. All rights reserved.
//

import Foundation
import UIKit

// MARK: - Defines further functionality for the UIView class
extension UIView {
    
    /// Creates a unique identifier for a UIView object based on its superviews
    ///
    /// - Returns: the unique identifier as a string object
    func uniqueIdentifer() -> String! {
        var identifier: String = "\(NSStringFromClass(type(of: self)))"
        /*if let responder = self.next {
            if responder is UIViewController {
                identifier += "\(NSStringFromClass(responder.classForCoder).componentsSeparatedByString(".").last!)"
            } }*/
        
        if let index = self.superview?.subviews.index(of: self) {
            identifier += "\(index)"
        }
        
        var nextView: UIView? = self
        while nextView != nil {
            
            if identifier.contains("UIViewController"){
                print(identifier)
                return identifier
            }
            
            nextView = nextView?.superview
            if let superview = nextView {
                identifier += "\(NSStringFromClass(type(of: superview)))"
            }
            if let index = nextView?.superview?.subviews.index(of: nextView!) {
                identifier += "\(index)"
            }
        }
        
        print(identifier)
        return identifier
    }

}
