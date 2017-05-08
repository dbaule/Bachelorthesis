//
//  EventModel.swift
//  DummyApp
//
//  Created by Danny Baule on 25.11.16.
//  Copyright © 2016 Danny Baule. All rights reserved.
//

import Foundation

//TODO: like this or do i need less information.. get other information from viewcontrollers db on server maybe, as there is no possibility to get the new vc in sendAction
/// Represents a transferable Event that can be send to the server
struct EventModel {
    
    /// The ViewController from where the event was send
    var Source = ""
    
    /// The name of the event that was fired
    var Name = ""
    
    /// The target ViewController that is present after the event was handled
    var Destination = ""
    
    /// The JSON Representation of the event
    var jsonRepresentation : String{
        return "{\"source\":\"\(Source)\",\"name\":\"\(Name)\",\"destination\":\"\(Destination)\"}"
    }
}
