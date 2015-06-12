package models

type Version struct {
	Basic
	Data VersionData `json:"data"`
}

type VersionData struct {
	Version string `json:"version"`
}
