//
//  ViewControllerModel.swift
//  DummyApp
//
//  Created by Danny Baule on 15.11.16.
//  Copyright © 2016 Danny Baule. All rights reserved.
//

import Foundation

/// Represents a transferable ViewController that can be send to the server
struct ViewControllerModel {
    
    /// the build version
    var BuildVersion = ""
    
    /// the assigned name
    var Name = ""
    
    /// the ui components that it holds
    var UIComponents = [UIComponentModel]()

    /// the JsonRepresentation
    var jsonRepresentation : String {
        var result = "{\"build_version\":\"\(BuildVersion)\",\"name\":\"\(Name)\",\"ui_components\":["
        
        var x = 0
        
        for comp in UIComponents{
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
