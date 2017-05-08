package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// compareAllViewController takes two ViewController arrays and calls compareViewController for
// each pair of ViewController with the same name
// the newVCS array contains all ViewController with the current highest build_version
// the oldVCS array contains all ViewController with the current second highest build_version
func compareAllViewController(newVCS []ViewController, oldVCS []ViewController, w http.ResponseWriter, r *http.Request) int {
	// traces the number of differences found
	numOfDifferences := 0

	// iterates over all ViewController from newVCS
	for _, newVC := range newVCS {
		// iterates over all ViewController from oldVCS and looks for a ViewController with
		// the same name as newVC
		// then calls compareViewController on them and tracks the numOfDifferences
		for _, oldVC := range oldVCS {
			if oldVC.Name == newVC.Name {
				numOfDifferences += compareViewController(newVC, oldVC)
			}
		}
	}

	// casts the numOfDifferences to a string
	t := strconv.Itoa(numOfDifferences)

	// prints out the result
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Found " + t + " difference(s) in total")
	fmt.Println("-----------------------------------------------------------")
	// and returns the number as well
	generateHTML(w, r)
	return numOfDifferences
}

// compareViewController takes to ViewController and compares their UIComponents against each other to determine differences
func compareViewController(newVC ViewController, oldVC ViewController) int {

	// numOfDifferencesFound saves the number of Differences that were detected on the ViewControllers
	numOfDifferencesFound := 0

	// marks the beginning of the comparison
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("Comparing " + newVC.Name + " with Build-Version " + oldVC.BuildVersion + " and " + newVC.BuildVersion)
	fmt.Println("-----------------------------------------------------------")
	fmt.Println()

	// iterate over the UIComponents of the newVC
	for _, uicompNew := range newVC.UIComponents {

		// allows only components with a name to be compared
		if uicompNew.Name != "" {
			// prints the current UIComponent name
			fmt.Println(uicompNew.Name)

			// saves the numOfDifferencesFound before comparing in order to later determine if errors where found
			//rememberNumDiffs := numOfDifferencesFound

			// iterate over the UIComponents of the oldVC
			for _, uicompOld := range oldVC.UIComponents {

				// only compares them, when the name is equal
				if uicompOld.Name == uicompNew.Name {
					// compares the height
					if uicompOld.Height != uicompNew.Height {

						// logs the difference and increments the numOfDifferencesFound
						fmt.Println("-----------------------------------------------------------")
						fmt.Println("  The height is different to the previous build-version!")
						var result = fmt.Sprintf("  Was: %.2f now: %.2f", uicompOld.Height, uicompNew.Height)
						fmt.Println(result)

						numOfDifferencesFound++
					}

					// compares the width
					if uicompOld.Width != uicompNew.Width {

						// logs the difference and increments the numOfDifferencesFound
						fmt.Println("-----------------------------------------------------------")
						fmt.Println("  The width is different to the previous build-version!")
						var result = fmt.Sprintf("  Was: %.2f now: %.2f", uicompOld.Width, uicompNew.Width)
						fmt.Println(result)

						numOfDifferencesFound++
					}

					// compares the XCoor
					if uicompOld.XCoor != uicompNew.XCoor {

						// logs the difference and increments the numOfDifferencesFound
						fmt.Println("-----------------------------------------------------------")
						fmt.Println("  The x-Coordinate is different to the previous build-version!")
						var result = fmt.Sprintf("  Was: %.2f now: %.2f", uicompOld.XCoor, uicompNew.XCoor)
						fmt.Println(result)

						numOfDifferencesFound++
					}

					// compares the YCoor
					if uicompOld.YCoor != uicompNew.YCoor {

						// logs the difference and increments the numOfDifferencesFound
						fmt.Println("-----------------------------------------------------------")
						fmt.Println("  The y-Coordinate is different to the previous build-version!")
						var result = fmt.Sprintf("  Was: %.2f now: %.2f", uicompOld.YCoor, uicompNew.YCoor)
						fmt.Println(result)

						numOfDifferencesFound++
					}

					// subview comparison
					fmt.Println("    -----------------------------------------------------------")
					fmt.Println("    SubviewComponents comparison beginning")

					numOfDifferencesFound += compareAllSubViews(uicompNew.SubviewComponents, uicompOld.SubviewComponents)
				}
			}
			// when no errors where detected log this message
			//if rememberNumDiffs == numOfDifferencesFound {
			//	fmt.Println("-----------------------------------------------------------")
			//	fmt.Println("  No differences found for " + uicompNew.Name)
			//}
			fmt.Println()
		}
	}

	fmt.Println("-----------------------------------------------------------")

	// when no differences where found, sets hasDifferences to false
	// if there were differences log the number of found differences
	if numOfDifferencesFound == 0 {
		fmt.Println("FINISHED: There were no differences detected in " + newVC.Name)
	} else {
		t := strconv.Itoa(numOfDifferencesFound)
		fmt.Println("FINISHED: " + t + " difference(s) found!")
	}

	fmt.Println("-----------------------------------------------------------")

	// returns the number of differences that were found
	return numOfDifferencesFound
}

// compareAllSubViews will check two UIComponent arrays on UILabel and UIImageView components and send them to the according comparison method
func compareAllSubViews(newUIComponents []UIComponent, oldUIComponents []UIComponent) int {

	// saves the number of differences that were found
	numOfDifferencesFound := 0

	// iterates over all the newUIComponents
	for _, newUIComponent := range newUIComponents {
		// iterates over all the oldUIComponents
		for _, oldUIComponent := range oldUIComponents {
			// when two UIComponent with the same name are found, they will be compared in one of three ways
			if newUIComponent.Name == oldUIComponent.Name {
				// Way 1: it is a label component
				// Therefore it needs to have no further SubviewComponents and contain Label in its ComponentType
				// compareLabelSubview will be called
				// Way 2: it is a UIImageView component
				// Therefore it needs to have no further SubviewComponents and contain UIImageView in its ComponentType
				// compareUIImageViewSubview will be called
				// Way 3: it is none of both and thus needs to be checked recursively for UILabel and UIImageView components in its SubviewComponents
				// Therefore it will call this method with the SubviewComponents array
				if len(newUIComponent.SubviewComponents) == 0 && strings.Contains(newUIComponent.ComponentType, "Label") {
					numOfDifferencesFound += compareLabelSubview(newUIComponent, oldUIComponent)
				} else if len(newUIComponent.SubviewComponents) == 0 && strings.Contains(newUIComponent.ComponentType, "UIImageView") {
					numOfDifferencesFound += compareUIImageViewSubview(newUIComponent, oldUIComponent)
				} else {
					numOfDifferencesFound += compareAllSubViews(newUIComponent.SubviewComponents, oldUIComponent.SubviewComponents)
				}
			}
		}
	}
	// at the end, the numOfDifferencesFound will be returned
	return numOfDifferencesFound
}

// compareLabelSubview takes two UIComponents that were earlier determined as UILabel objects and compares them on the UILabel properties
func compareLabelSubview(newUIComponent UIComponent, oldUIComponent UIComponent) int {

	// saves the number of differences that were found
	numOfDifferencesFound := 0

	// marks the beginning of the comparison
	//fmt.Println("    -----------------------------------------------------------")
	//fmt.Println("    UILabel: Comparison of " + newUIComponent.Name)
	//fmt.Println("    -----------------------------------------------------------")
	//fmt.Println()

	// comparison of the UILabel_Text
	if newUIComponent.UILabelText != oldUIComponent.UILabelText {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      The UILabel_Text of " + newUIComponent.ComponentType + " : \"" + newUIComponent.UILabelText + "\" is different to the previous build-version!")
		fmt.Println("      Was: " + oldUIComponent.UILabelText + " now: " + newUIComponent.UILabelText)

		numOfDifferencesFound++
	}

	// comparison of the UILabel_Font
	if newUIComponent.UILabelFont != oldUIComponent.UILabelFont {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      The UILabel_Font of " + newUIComponent.ComponentType + " : \"" + newUIComponent.UILabelText + "\" is different to the previous build-version!")
		fmt.Println("      Was: " + oldUIComponent.UILabelFont + " now: " + newUIComponent.UILabelFont)

		numOfDifferencesFound++
	}

	// comparison of the UILabel_Size
	if newUIComponent.UILabelFontSize != oldUIComponent.UILabelFontSize {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      The UILabel_Size of " + newUIComponent.ComponentType + " : \"" + newUIComponent.UILabelText + "\" is different to the previous build-version!")
		var result = fmt.Sprintf("      Was: %.2f now: %.2f", oldUIComponent.UILabelFontSize, newUIComponent.UILabelFontSize)
		fmt.Println(result)

		numOfDifferencesFound++
	}

	// warning for truncated text
	//TODO: acurate value to substitute
	if (newUIComponent.UILabelTextWidth - 5) > newUIComponent.Width {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      Warning: The UILabel_Text of " + newUIComponent.ComponentType + " : \"" + newUIComponent.UILabelText + "\" might be truncated!")
		var result = fmt.Sprintf("      As the label width is: %.2f, while the text width is: %.2f", newUIComponent.Width, newUIComponent.UILabelTextWidth)
		fmt.Println(result)

		numOfDifferencesFound++
	}

	// when no errors where detected log this message
	//if numOfDifferencesFound == 0 {
	//	fmt.Println("    -----------------------------------------------------------")
	//	fmt.Println("      No differences found for " + newUIComponent.Name)
	//}
	fmt.Println()

	// return the number of differences that were found
	return numOfDifferencesFound
}

// compareUIImageViewSubview takes two UIComponents that were earlier determined as UIImageView objects and compares their properties with each other
func compareUIImageViewSubview(newUIComponent UIComponent, oldUIComponent UIComponent) int {

	// saves the number of differences found
	numOfDifferencesFound := 0

	// marks the beginning of the comparison
	//fmt.Println("    -----------------------------------------------------------")
	//fmt.Println("    UIImageView: Comparison of " + newUIComponent.Name)
	//fmt.Println("    -----------------------------------------------------------")
	//fmt.Println()

	// comparison of the image height
	if newUIComponent.UIImageViewImageHeight != oldUIComponent.UIImageViewImageHeight {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      The UIImageView_ImageHeight is different to the previous build-version!")
		var result = fmt.Sprintf("      Was: %.2f now: %.2f", oldUIComponent.UIImageViewImageHeight, newUIComponent.UIImageViewImageHeight)
		fmt.Println(result)

		numOfDifferencesFound++
	}

	// comparison of the image width
	if newUIComponent.UIImageViewImageWidth != oldUIComponent.UIImageViewImageWidth {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      The UIImageView_ImageWidth is different to the previous build-version!")
		var result = fmt.Sprintf("      Was: %.2f now: %.2f", oldUIComponent.UIImageViewImageWidth, newUIComponent.UIImageViewImageWidth)
		fmt.Println(result)

		numOfDifferencesFound++
	}

	// comparison of the image scale
	if newUIComponent.UIImageViewImageScale != oldUIComponent.UIImageViewImageScale {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      The UIImageView_ImageScale is different to the previous build-version!")
		var result = fmt.Sprintf("      Was: %.2f now: %.2f", oldUIComponent.UIImageViewImageScale, newUIComponent.UIImageViewImageScale)
		fmt.Println(result)

		numOfDifferencesFound++
	}

	// comparison of the image orientation
	if newUIComponent.UIImageViewImageOrientation != oldUIComponent.UIImageViewImageOrientation {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      The UIImageView_ImageOrientation is different to the previous build-version!")
		var result = fmt.Sprintf("      Was: %.2f now: %.2f", oldUIComponent.UIImageViewImageOrientation, newUIComponent.UIImageViewImageOrientation)
		fmt.Println(result)

		numOfDifferencesFound++
	}

	// comparison of the UIImageView size and its image size
	if newUIComponent.Height != newUIComponent.UIImageViewImageHeight || newUIComponent.Width != oldUIComponent.UIImageViewImageWidth {
		// logs the difference and increments the numOfDifferencesFound
		fmt.Println("    -----------------------------------------------------------")
		fmt.Println("      Warning: The UIImageView_Image size is different to the UIImageView size!")
		var result = fmt.Sprintf("      While UIImageView is %.2fx%.2f, UIImageView_Image is %.2fx%.2f", newUIComponent.Height, newUIComponent.Width, newUIComponent.UIImageViewImageHeight, newUIComponent.UIImageViewImageWidth)
		fmt.Println(result)

		numOfDifferencesFound++
	}

	// when no errors where detected log this message
	//if numOfDifferencesFound == 0 {
	//	fmt.Println("    -----------------------------------------------------------")
	//	fmt.Println("      No differences found for " + newUIComponent.Name)
	//}
	fmt.Println()

	// returns the number of differences found
	return numOfDifferencesFound
}
