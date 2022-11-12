package structures

import "encoding/json"

type List struct {
	Id       string           `json:"id,omitempty"`
	Name     string           `json:"name"`
	Elements []RandomElements `json:"elements"`
}

type RandomElements struct {
	Name        string   `json:"name"`
	Probability int      `json:"probability"`
	Tags        []string `json:"tags"`
}

// ToJSON returns the list as a json string
func (list *List) ToJSON() ([]byte, error) {
	return json.Marshal(list)
}

// FromJSON unmarshal a json into a List struct and correct malformed elements
func (list *List) FromJSON(data []byte) error {
	//unmarshal the data to a list
	if err := json.Unmarshal(data, list); err != nil {
		return err
	}

	//loop through each list element to check that every element is correctly formed
	var correctedElements []RandomElements
	for _, element := range list.Elements {

		//check if the element has a name
		if element.Name == "" {
			element.Name = "Unnamed element"
		}

		// check if the probability is between 0 and 100
		if element.Probability < 0 || element.Probability > 100 {
			element.Probability = 0
		}

		//check if the properties are correctly formed (not nil)
		if element.Tags == nil {
			element.Tags = []string{}
		}

		//add the element to the corrected elements array
		correctedElements = append(correctedElements, element)
	}

	//set the corrected elements to the list
	list.Elements = correctedElements

	//check if elements is nil
	if list.Elements == nil {
		list.Elements = []RandomElements{}
	}

	return nil
}
