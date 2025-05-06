package menu_elements

type ViewElements struct {
	elements []IViewElement
}

func NewViewElements() *ViewElements {
	return &ViewElements{elements: make([]IViewElement, 0, 20)}
}

func (this *ViewElements) GetElements() []IViewElement {
	return this.elements
}

func (this *ViewElements) SetElements(elements []IViewElement) {
	this.elements = elements
}

func (this *ViewElements) AddElement(el IViewElement) {
	this.elements = append(this.elements, el)
}

func (this *ViewElements) RemoveElement(el IViewElement) {
	var filteredElements []IViewElement
	for _, element := range this.elements {
		if element.Id() == el.Id() {
			continue
		}
		filteredElements = append(filteredElements, element)
	}

	this.elements = filteredElements
}
