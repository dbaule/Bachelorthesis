//
//  ServerInteractionHandler.swift
//  DummyApp
//
//  Created by Danny Baule on 15.11.16.
//  Copyright © 2016 Danny Baule. All rights reserved.
//

import Foundation

/// handles the interaction with the server
class ServerInteractionHandler{
    
    /// the server adress
    let serverAddress = "http://localhost:8080"
    
    /// the name of a post method
    let postMethodName = "POST"
    
    /// posts a ViewController to the server
    ///
    /// - Parameter vc: the ViewController to publish
    func postVCtoServer(vc: ViewControllerModel) {
        
        // the request setup
        var request = URLRequest(url: URL(string: serverAddress + "/viewcontrollers")!)
        request.httpMethod = postMethodName
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        
        // gets the postString (jsonRepresentation)
        let postString = vc.jsonRepresentation
        
        // adds the body to the request
        request.httpBody = postString.data(using: .utf8)
        
        // starts the task, terminates when an error occurs
        let task = URLSession.shared.dataTask(with: request) {
            data, response, error in
            guard let data = data, error == nil else {
                print("error=\(error)")
                return
            }
            
            // prints the response, when the StatusCode is not 200
            // can be adjusted as the server currently returns 201 for created
            if let httpStatus = response as? HTTPURLResponse, httpStatus.statusCode != 200 {
                print("statusCode should be 200, but is \(httpStatus.statusCode)")
                print("response = \(response)")
            }
            
            // prints the responsestring
            let responseString = String(data: data, encoding: .utf8)
            print("responseString = \(responseString)")
        }
        
        task.resume()
        
    }
    
    /// posts an Event to the server
    ///
    /// - Parameter event: the event to publish
    func postEventToServer(event: EventModel) {
        
        // the request setup
        var request = URLRequest(url: URL(string: serverAddress + "/events")!)
        request.httpMethod = postMethodName
        request.addValue("application/json", forHTTPHeaderField: "Content-Type")
        
        // gets the postString (jsonRepresentation)
        let postString = event.jsonRepresentation
        
        // adds the body to the request
        request.httpBody = postString.data(using: .utf8)
        
        // starts the task
        let task = URLSession.shared.dataTask(with: request) {
            data, response, error in
            guard let data = data, error == nil else {
                print("error=\(error)")
                return
            }
            
            // prints the response, when the StatusCode is not 200
            // can be adjusted as the server currently returns 201 for created
            if let httpStatus = response as? HTTPURLResponse, httpStatus.statusCode != 200 {
                print("statusCode should be 200, but is \(httpStatus.statusCode)")
                print("response = \(response)")
            }
            
            // prints the responsestring
            let responseString = String(data: data, encoding: .utf8)
            print("responseString = \(responseString)")
        }
        
        task.resume()
        
    }
}
