package main

import (
	// Standard Library Imports
	"os"
	"time"

	// External Imports
	log "github.com/sirupsen/logrus"

	// Internal Imports
	"github.com/matthewhartstonge/graph"
	"github.com/matthewhartstonge/graph/_examples/utils"
)

func init() {
	// Log actions so you can see what is happening internally.
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// todo: load graph from file.
	meta := utils.LoadJsonGraph("_examples/mailbot/graph.json")

	start := time.Now()
	V, E := utils.JSONGraphToVE(meta)
	G := graph.New(
		graph.WithVertices(V),
		graph.WithEdges(E),
	)
	log.WithField("took", time.Since(start)).Info()

	G.PrintInfo()
}
