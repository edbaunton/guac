// Code generated by ent, DO NOT EDIT.

package certification

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the certification type in the database.
	Label = "certification"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSourceID holds the string denoting the source_id field in the database.
	FieldSourceID = "source_id"
	// FieldPackageVersionID holds the string denoting the package_version_id field in the database.
	FieldPackageVersionID = "package_version_id"
	// FieldPackageNameID holds the string denoting the package_name_id field in the database.
	FieldPackageNameID = "package_name_id"
	// FieldArtifactID holds the string denoting the artifact_id field in the database.
	FieldArtifactID = "artifact_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldJustification holds the string denoting the justification field in the database.
	FieldJustification = "justification"
	// FieldOrigin holds the string denoting the origin field in the database.
	FieldOrigin = "origin"
	// FieldCollector holds the string denoting the collector field in the database.
	FieldCollector = "collector"
	// FieldKnownSince holds the string denoting the known_since field in the database.
	FieldKnownSince = "known_since"
	// EdgeSource holds the string denoting the source edge name in mutations.
	EdgeSource = "source"
	// EdgePackageVersion holds the string denoting the package_version edge name in mutations.
	EdgePackageVersion = "package_version"
	// EdgeAllVersions holds the string denoting the all_versions edge name in mutations.
	EdgeAllVersions = "all_versions"
	// EdgeArtifact holds the string denoting the artifact edge name in mutations.
	EdgeArtifact = "artifact"
	// Table holds the table name of the certification in the database.
	Table = "certifications"
	// SourceTable is the table that holds the source relation/edge.
	SourceTable = "certifications"
	// SourceInverseTable is the table name for the SourceName entity.
	// It exists in this package in order to avoid circular dependency with the "sourcename" package.
	SourceInverseTable = "source_names"
	// SourceColumn is the table column denoting the source relation/edge.
	SourceColumn = "source_id"
	// PackageVersionTable is the table that holds the package_version relation/edge.
	PackageVersionTable = "certifications"
	// PackageVersionInverseTable is the table name for the PackageVersion entity.
	// It exists in this package in order to avoid circular dependency with the "packageversion" package.
	PackageVersionInverseTable = "package_versions"
	// PackageVersionColumn is the table column denoting the package_version relation/edge.
	PackageVersionColumn = "package_version_id"
	// AllVersionsTable is the table that holds the all_versions relation/edge.
	AllVersionsTable = "certifications"
	// AllVersionsInverseTable is the table name for the PackageName entity.
	// It exists in this package in order to avoid circular dependency with the "packagename" package.
	AllVersionsInverseTable = "package_names"
	// AllVersionsColumn is the table column denoting the all_versions relation/edge.
	AllVersionsColumn = "package_name_id"
	// ArtifactTable is the table that holds the artifact relation/edge.
	ArtifactTable = "certifications"
	// ArtifactInverseTable is the table name for the Artifact entity.
	// It exists in this package in order to avoid circular dependency with the "artifact" package.
	ArtifactInverseTable = "artifacts"
	// ArtifactColumn is the table column denoting the artifact relation/edge.
	ArtifactColumn = "artifact_id"
)

// Columns holds all SQL columns for certification fields.
var Columns = []string{
	FieldID,
	FieldSourceID,
	FieldPackageVersionID,
	FieldPackageNameID,
	FieldArtifactID,
	FieldType,
	FieldJustification,
	FieldOrigin,
	FieldCollector,
	FieldKnownSince,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Type defines the type for the "type" enum field.
type Type string

// TypeGOOD is the default value of the Type enum.
const DefaultType = TypeGOOD

// Type values.
const (
	TypeGOOD Type = "GOOD"
	TypeBAD  Type = "BAD"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeGOOD, TypeBAD:
		return nil
	default:
		return fmt.Errorf("certification: invalid enum value for type field: %q", _type)
	}
}

// OrderOption defines the ordering options for the Certification queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySourceID orders the results by the source_id field.
func BySourceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSourceID, opts...).ToFunc()
}

// ByPackageVersionID orders the results by the package_version_id field.
func ByPackageVersionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPackageVersionID, opts...).ToFunc()
}

// ByPackageNameID orders the results by the package_name_id field.
func ByPackageNameID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPackageNameID, opts...).ToFunc()
}

// ByArtifactID orders the results by the artifact_id field.
func ByArtifactID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldArtifactID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByJustification orders the results by the justification field.
func ByJustification(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJustification, opts...).ToFunc()
}

// ByOrigin orders the results by the origin field.
func ByOrigin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrigin, opts...).ToFunc()
}

// ByCollector orders the results by the collector field.
func ByCollector(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollector, opts...).ToFunc()
}

// ByKnownSince orders the results by the known_since field.
func ByKnownSince(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKnownSince, opts...).ToFunc()
}

// BySourceField orders the results by source field.
func BySourceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSourceStep(), sql.OrderByField(field, opts...))
	}
}

// ByPackageVersionField orders the results by package_version field.
func ByPackageVersionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPackageVersionStep(), sql.OrderByField(field, opts...))
	}
}

// ByAllVersionsField orders the results by all_versions field.
func ByAllVersionsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAllVersionsStep(), sql.OrderByField(field, opts...))
	}
}

// ByArtifactField orders the results by artifact field.
func ByArtifactField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newArtifactStep(), sql.OrderByField(field, opts...))
	}
}
func newSourceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SourceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SourceTable, SourceColumn),
	)
}
func newPackageVersionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PackageVersionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PackageVersionTable, PackageVersionColumn),
	)
}
func newAllVersionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AllVersionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AllVersionsTable, AllVersionsColumn),
	)
}
func newArtifactStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ArtifactInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ArtifactTable, ArtifactColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Type) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Type) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Type(str)
	if err := TypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Type", str)
	}
	return nil
}
