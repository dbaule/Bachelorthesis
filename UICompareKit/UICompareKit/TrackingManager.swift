//
//  TrackingManager.swift
//  DummyApp
//
//  Created by Danny Baule on 29.11.16.
//  Copyright © 2016 Danny Baule. All rights reserved.
//

import Foundation

/// a class that helps to track events and ViewController
class TrackingManager {
    
    //MARK: Shared Instance
    
    static let sharedInstance : TrackingManager = {
        let instance = TrackingManager(array: [])
        return instance
    }()
    
    //MARK: Local Variable
    
    /// an array that holds all the information about events and ViewController that were execute and presented
    var trackingArray : [Any]{
        didSet{
            transitionHappened()
        }
    }
    
    //TODO: Methods to write event / track viewController / based on types in didSet
    
    //MARK: Init
    
    /// initializes the singleton with a new array
    init( array : [Any]) {
        trackingArray = array
    }
    
    //MARK: Functions
    
    
    /// sends a newly detected event to the server
    private func transitionHappened() {
        
        // the number of elements in the trackingArray
        let trackingArraySize = trackingArray.count
        
        // if more than 2 elements are in the array it will check for a possible transition
        if trackingArraySize > 2 {
            
            // when the lastEvent that happened is a ViewController
            if let lastEvent = trackingArray[trackingArraySize - 1] as? ViewControllerModel{
                
                // and the previous to last event is an action
                if var previousToLastEvent = trackingArray[trackingArraySize - 2] as? String{
                    
                    if previousToLastEvent == "perform:" {
                        previousToLastEvent = trackingArray[trackingArraySize - 3] as! String
                    }
                    
                    // removes the lastEvent and the previousToLastEvent
                    let prunedArray = trackingArray.dropLast(2)
                    
                    // and look for the last ViewController in the prunedArray
                    if let previousVC = lastViewControllerInArray(array: prunedArray){
                       
                        // if it found the last ViewController, it will create a new EventModel
                        var event = EventModel()
                        event.Source = previousVC.Name
                        event.Name = previousToLastEvent
                        event.Destination = lastEvent.Name
                        
                        // and publish it to the server
                        ServerInteractionHandler().postEventToServer(event: event)
                    }
                    
                }
            }
        }
    }
    
    /// iterates the array backwards and looks for a ViewController
    ///
    /// - Parameter array: the array to search for the lastViewController
    /// - Returns: returns the first found ViewController or nil if there is no VC
    private func lastViewControllerInArray(array : ArraySlice<Any>) -> ViewControllerModel?{
        
        // backwards iteration
        for index in stride(from: (array.count - 1), to: 0, by: -1){
            // if element at index is a ViewController, returns it
            if let viewController = array[index] as? ViewControllerModel{
                return viewController
            }
        }
        // if no ViewController is found, returns nil
        return nil
    }
}
