/*
 * Copyright (C) 2019. Matthew Hartstonge <matt@mykro.co.nz>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package utils

import (
	// Standard Library Imports
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	// External Imports
	log "github.com/sirupsen/logrus"

	// Internal Imports
	"github.com/matthewhartstonge/graph/edge"
	"github.com/matthewhartstonge/graph/vertex"
)

type JSONGraph struct {
	Vertices []struct {
		Label string `json:"label"`
	} `json:"vertices"`
	Edges []struct {
		V1       string  `json:"v1"`
		V2       string  `json:"v2"`
		Cost     float64 `json:"cost"`
		Directed bool    `json:"directed"`
	} `json:"edges"`
}

func LoadJsonGraph(fp string) (meta *JSONGraph) {
	gopath := os.Getenv("GOPATH")
	fp = filepath.Clean(fmt.Sprintf("%s/src/github.com/matthewhartstonge/graph/%s", gopath, fp))
	f, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}

	meta = &JSONGraph{}
	err = json.Unmarshal(f, meta)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func JSONGraphToVE(metadata *JSONGraph) (V []vertex.Vertexer, E []edge.Edger) {
	var vertexMap = make(map[string]vertex.Vertexer)
	var vertices []vertex.Vertexer
	for _, jsonVertex := range metadata.Vertices {
		v := vertex.New(jsonVertex.Label)
		vertexMap[v.Label()] = v
		vertices = append(vertices, v)
	}

	var edges []edge.Edger
	for _, jsonEdge := range metadata.Edges {
		v1, ok := vertexMap[jsonEdge.V1]
		if !ok {
			continue
		}
		v2, ok := vertexMap[jsonEdge.V2]
		if !ok {
			continue
		}

		e := edge.New(
			v1, v2,
			edge.WithDirected(jsonEdge.Directed),
			edge.WithCost(jsonEdge.Cost),
		)
		edges = append(edges, e)
	}

	vertexMap = nil
	return vertices, edges
}
