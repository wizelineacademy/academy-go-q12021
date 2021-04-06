package vo

type Paged struct {
	Type            string `query:"type" param:"type" validate:"required,oneof=odd even"`
	Items           int    `query:"items" param:"items" validate:"required"`
	ItemsPerWorkers int    `query:"items_per_workers" param:"items_per_workers" validate:"required"`
}
