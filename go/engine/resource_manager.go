package engine

type abstractResource struct {
	id        uint
	destroyed bool
}

func (ar *abstractResource) SetId(id uint) {
	ar.id = id
}

type abstractResourceInterface interface {
	SetId(id uint)
	Destroy()
}

type resourceManager struct {
	nextId    uint
	resources map[uint]abstractResourceInterface
}

func newResourceManager() *resourceManager {
	return &resourceManager{resources: map[uint]abstractResourceInterface{}}
}

func (rm *resourceManager) Save(resource abstractResourceInterface) abstractResourceInterface {
	resource.SetId(rm.nextId)
	rm.resources[rm.nextId] = resource
	rm.nextId++
	return resource
}

func (rm *resourceManager) Destroy() {
	for _, resource := range rm.resources {
		resource.Destroy()
	}
}
