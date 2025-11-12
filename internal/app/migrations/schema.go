package migrations

import "embed"

//go:embed schema/*.sql
var Schema embed.FS

var SchemaPath = "schema"
