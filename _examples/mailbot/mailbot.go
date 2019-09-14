package main

import (
	"os"
	"time"

	json "github.com/matthewhartstonge/graph/_examples/utils"
	log "github.com/sirupsen/logrus"

	"github.com/matthewhartstonge/graph"
)

func init() {
	// Log actions so you can see what is happening internally.
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// todo: load graph from file.
	meta := json.LoadJsonGraph("_examples/mailbot/graph.json")

	start := time.Now()
	V, E := json.ConvertToVE(meta)
	G := graph.New(
		graph.WithVertices(V),
		graph.WithEdges(E),
	)
	log.WithField("took", time.Since(start)).Info()

	G.PrintInfo()
}
