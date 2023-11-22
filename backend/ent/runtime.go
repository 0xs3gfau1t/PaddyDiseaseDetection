// Code generated by ent, DO NOT EDIT.

package ent

import (
	"segFault/PaddyDiseaseDetection/ent/diseaseidentified"
	"segFault/PaddyDiseaseDetection/ent/image"
	"segFault/PaddyDiseaseDetection/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	diseaseidentifiedFields := schema.DiseaseIdentified{}.Fields()
	_ = diseaseidentifiedFields
	// diseaseidentifiedDescSeverity is the schema descriptor for severity field.
	diseaseidentifiedDescSeverity := diseaseidentifiedFields[2].Descriptor()
	// diseaseidentified.SeverityValidator is a validator for the "severity" field. It is called by the builders before save.
	diseaseidentified.SeverityValidator = diseaseidentifiedDescSeverity.Validators[0].(func(int) error)
	// diseaseidentifiedDescCreatedAt is the schema descriptor for created_at field.
	diseaseidentifiedDescCreatedAt := diseaseidentifiedFields[3].Descriptor()
	// diseaseidentified.DefaultCreatedAt holds the default value on creation for the created_at field.
	diseaseidentified.DefaultCreatedAt = diseaseidentifiedDescCreatedAt.Default.(func() time.Time)
	imageFields := schema.Image{}.Fields()
	_ = imageFields
	// imageDescCreatedAt is the schema descriptor for created_at field.
	imageDescCreatedAt := imageFields[2].Descriptor()
	// image.DefaultCreatedAt holds the default value on creation for the created_at field.
	image.DefaultCreatedAt = imageDescCreatedAt.Default.(func() time.Time)
}
