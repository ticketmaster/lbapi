package backup

// DbRecord - resource configuration.
type DbRecord struct {
	Data interface{} `json:"data,omitempty"`
	ID   string      `json:"id,omitempty"`
}

// DbRecordCollection - resource configuration.
type DbRecordCollection struct {
	DbRecords []DbRecord `json:"db_records,omitempty"`
}
