package snowflake

import (
	"strings"

	"gorm.io/gorm/schema"
)

// NamingStrategy for snowflake (always uppercase)
type NamingStrategy struct {
	DefaultNS schema.Namer
}

// NewNamingStrategy create new instance of snowflake naming strat
func NewNamingStrategy() schema.Namer {
	return &NamingStrategy{
		DefaultNS: schema.NamingStrategy{},
	}
}

// ColumnName snowflake edition
func (sns NamingStrategy) ColumnName(table, column string) string {
	return strings.ToUpper(sns.DefaultNS.ColumnName(table, column))
}

// TableName snowflake edition
func (sns NamingStrategy) TableName(table string) string {
	return sns.DefaultNS.TableName(table)
}

// SchemaName snowflake edition
func (sns NamingStrategy) SchemaName(table string) string {
	return sns.DefaultNS.SchemaName(table)
}

// JoinTableName snowflake edition
func (sns NamingStrategy) JoinTableName(joinTable string) string {
	return sns.DefaultNS.JoinTableName(joinTable)
}

// RelationshipFKName snowflake edition
func (sns NamingStrategy) RelationshipFKName(rel schema.Relationship) string {
	return sns.DefaultNS.RelationshipFKName(rel)
}

// CheckerName snowflake edition
func (sns NamingStrategy) CheckerName(table, column string) string {
	return sns.DefaultNS.CheckerName(table, column)
}

// IndexName snowflake edition
func (sns NamingStrategy) IndexName(table, column string) string {
	return sns.DefaultNS.IndexName(table, column)
}
