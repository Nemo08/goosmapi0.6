package goosmapi

import (
	"encoding/xml"
	"time"
)

//In this format _list_ of changesets return from osm.org
// Get from http://api.openstreetmap.org/api/0.6/changesets?
type ChangesetList struct {
	XMLName    xml.Name        `xml:"osm"`
	Changesets []ChangesetInfo `xml:"changeset"`
	AttributionBasicInfo
}

// Part of _list_ OSM changesets response structure: changesets
// Get from http://www.openstreetmap.org/api/0.6/changeset/_changeset_id_/download
type ChangesetInfo struct {
	ChangesetId   int64      `xml:"id,attr"`
	CreatedAt     *time.Time `xml:"created_at,attr"`
	ClosedAt      *time.Time `xml:"closed_at,attr"`
	Open          bool       `xml:"open,attr"`
	Maxlat        float64    `xml:"max_lat,attr"`
	Maxlon        float64    `xml:"max_lon,attr"`
	Minlat        float64    `xml:"min_lat,attr"`
	Minlon        float64    `xml:"min_lon,attr"`
	CommentsCount int32      `xml:"comments_count,attr"`
	Tags          []Tag      `xml:"tag"`
	UserBasicInfo
}

//Part of _list_ OSM changesets response structure: tags
type Tag struct {
	Key   string `xml:"k,attr"`
	Value string `xml:"v,attr"`
}

//Bound box: minlat, minlon, maxlat, maxlon,
type BoundsBox struct {
	Minlat float64 `xml:"min_lat,attr" json:"minlat"`
	Minlon float64 `xml:"min_lon,attr" json:"minlon"`
	Maxlat float64 `xml:"max_lat,attr" json:"maxlat"`
	Maxlon float64 `xml:"max_lat,attr" json:"maxlon"`
}

// Osmchange.xml, _one_ changeset information
// Get from http://www.openstreetmap.org/api/0.6/changeset/_changeset_ID_/download
type Changeset struct {
	XMLName           xml.Name   `xml:"osmChange"`
	CreatedNodes      []Node     `xml:"create>node"`
	CreatedWays       []Way      `xml:"create>way"`
	CreatedRelations  []Relation `xml:"create>relation"`
	ModifiedNodes     []Node     `xml:"modify>node"`
	ModifiedWays      []Way      `xml:"modify>way"`
	ModifiedRelations []Relation `xml:"modify>relation"`
	DeletedNodes      []Node     `xml:"delete>node"`
	DeletedWays       []Way      `xml:"delete>way"`
	DeletedRelations  []Relation `xml:"delete>relation"`
}

// History of _one_ node, way or relation
// Get from http://www.openstreetmap.org/api/0.6/[node|way|relation]/_id_/history
type History struct {
	XMLName   xml.Name   `xml:"osm"`
	Nodes     []Node     `xml:"node"`
	Ways      []Way      `xml:"way"`
	Relations []Relation `xml:"relation"`
	AttributionBasicInfo
}

// Common data for osm.org responces
type AttributionBasicInfo struct {
	Version     string `xml:"version,attr"`
	Generator   string `xml:"generator,attr"`
	Copyright   string `xml:"copyright,attr"`
	Attribution string `xml:"attribution,attr"`
	License     string `xml:"license,attr"`
}

// Common data for nodes, ways, relations
type ElementBasicInfo struct {
	Id          int64      `xml:"id,attr"`
	ChangesetId int64      `xml:"changeset,attr"`
	TimeStamp   *time.Time `xml:"timestamp,attr"`
	Version     int64      `xml:"version,attr"`
	Visible     bool       `xml:"visible,attr"`
}

// Common data for nodes, ways, relations : user info
type UserBasicInfo struct {
	User   string `xml:"user,attr"`
	UserId int64  `xml:"uid,attr"`
}

// Part of OSM changeset response structure (osmchange.xml): node
// Or taken from http://www.openstreetmap.org/api/0.6/node/_node_id_
type Node struct {
	XMLName xml.Name `xml:"node"`
	Lat     float64  `xml:"lat,attr"`
	Lon     float64  `xml:"lon,attr"`
	Tags    []Tag    `xml:"tag"`
	ElementBasicInfo
	UserBasicInfo
}

// Part of OSM changeset response structure (osmchange.xml): way
// Or taken from http://www.openstreetmap.org/api/0.6/way/_way_id_
type Way struct {
	XMLName  xml.Name  `xml:"way"`
	NodeRefs []NodeRef `xml:"nd"`
	Tags     []Tag     `xml:"tag"`
	ElementBasicInfo
	UserBasicInfo
}

//Part of OSM changeset response structure (osmchange.xml): relation
// Or taken from http://www.openstreetmap.org/api/0.6/relation/_relation_id_
type Relation struct {
	XMLName xml.Name         `xml:"relation"`
	Members []RelationMember `xml:"member"`
	Tags    []Tag            `xml:"tag"`
	ElementBasicInfo
	UserBasicInfo
}

//Part of OSM changeset response structure (osmchange.xml): relation member, reference by id
type RelationMember struct {
	XMLName xml.Name `xml:"member"`
	Type    string   `xml:"type,attr"`
	Ref     int64    `xml:"ref,attr"`
	Role    string   `xml:"role,attr"`
}

//Part of OSM changeset response structure: node, reference by node id
type NodeRef struct {
	XMLName xml.Name `xml:"nd"`
	Ref     int64    `xml:"ref,attr"`
}
