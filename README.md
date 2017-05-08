# Bachelorthesis
A GUI Testing Framework to automatically detect layout inconsistencies in iOS apps

The number of mobile apps has increased drastically in the past few years. For every task a user wants to accomplish there are multiple app available. To stand out from the other apps not only the range of functions but also a good UI is a key factor. A good UI provides a consistent appearance and simplifies the usage of the app. In this thesis we introduce an OS independent approach to automatically detect layout inconsistencies in an app’s UI. After a developer submits a change into the source code version system, a comparison process is initialized. This process automatically navigates through the app and stores information about the encountered UI elements. The data sets are then compared and the user is presented with the results. Our comparison detects e.g., dislocated UI components, re-sized components, or truncated label texts. We implemented our approach in a Swift-based library, called UICompareKit, that can be integrated into existing iOS apps. To evaluate UICompareKit we integrated it into nine open source iOS apps. UICompareKit was able to extract information about UI components, compare the collected data, and detect inconsistencies in the data sets.

UICompareKit consists of a client library written in Swift and a server component written in Go.

The library hooks into existing iOS apps, using method swizzling, to extract information about UI components from each encountered view controller.
The extracted information is saved in a mongoDB database.
The server component handles the interaction between client library and database.
Furthermore, it provides the comparison algorithm, which compares two data sets with each other and generates a comparison report.

The following image highlights the five steps of my approach:

[[https://github.com/dbaule/Bachelorthesis/blob/master/ApproachLifecycle.png|alt=octocat]]
