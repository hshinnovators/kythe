/*
 * Copyright 2015 Google Inc. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Binary ktags emits ctags-formatted lines for the definitions in the given files.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"kythe.io/kythe/go/services/xrefs"
	"kythe.io/kythe/go/util/flagutil"
	"kythe.io/kythe/go/util/kytheuri"
	"kythe.io/kythe/go/util/schema/edges"
	"kythe.io/kythe/go/util/schema/facts"
	"kythe.io/kythe/go/util/schema/nodes"

	"bitbucket.org/creachadair/stringset"
	"golang.org/x/net/context"

	gpb "kythe.io/kythe/proto/graph_proto"
	xpb "kythe.io/kythe/proto/xref_proto"
)

var (
	ctx = context.Background()

	corpus    = flag.String("corpus", "", "Corpus of the given files")
	remoteAPI = flag.String("api", "https://xrefs-dot-kythe-repo.appspot.com", "Remote api server")
)

func init() {
	flag.Usage = flagutil.SimpleUsage("Emit ctags-formatted lines for the definitions in the given files",
		"[--api address] <file>...")
}

// TODO(schroederc): use cross-language facts to determine a node's tag name.
// Currently, this fact is only emitted by the Java indexer.
const identifierFact = "/kythe/identifier"

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		flagutil.UsageError("not given any files")
	}

	xs := xrefs.WebClient(*remoteAPI)

	for _, file := range flag.Args() {
		ticket := (&kytheuri.URI{Corpus: *corpus, Path: file}).String()
		decor, err := xs.Decorations(ctx, &xpb.DecorationsRequest{
			Location:   &xpb.Location{Ticket: ticket},
			SourceText: true,
			References: true,
		})
		if err != nil {
			log.Fatalf("Failed to get decorations for file %q", file)
		}

		nmap := xrefs.NodesMap(decor.Nodes)
		var emitted stringset.Set

		for _, r := range decor.Reference {
			if r.Kind != edges.DefinesBinding || emitted.Contains(r.TargetTicket) {
				continue
			}

			ident := string(nmap[r.TargetTicket][identifierFact])
			if ident == "" {
				continue
			}

			offset, err := strconv.Atoi(string(nmap[r.SourceTicket][facts.AnchorStart]))
			if err != nil {
				log.Printf("Invalid start offset for anchor %q", r.SourceTicket)
				continue
			}

			fields, err := getTagFields(xs, r.TargetTicket)
			if err != nil {
				log.Printf("Failed to get tagfields for %q: %v", r.TargetTicket, err)
			}

			fmt.Printf("%s\t%s\t%d;\"\t%s\n",
				ident, file, offsetLine(decor.SourceText, offset), strings.Join(fields, "\t"))
			emitted.Add(r.TargetTicket)
		}
	}
}

func getTagFields(xs xrefs.Service, ticket string) ([]string, error) {
	reply, err := xs.Edges(ctx, &gpb.EdgesRequest{
		Ticket: []string{ticket},
		Kind:   []string{edges.ChildOf, edges.Param},
		Filter: []string{facts.NodeKind, facts.Subkind, identifierFact},
	})
	if err != nil || len(reply.EdgeSets) == 0 {
		return nil, err
	}

	var fields []string

	nmap := xrefs.NodesMap(reply.Nodes)
	emap := xrefs.EdgesMap(reply.EdgeSets)

	switch string(nmap[ticket][facts.NodeKind]) + "|" + string(nmap[ticket][facts.Subkind]) {
	case nodes.Function + "|":
		fields = append(fields, "f")
		fields = append(fields, "arity:"+strconv.Itoa(len(emap[ticket][edges.Param])))
	case nodes.Enum + "|" + nodes.EnumClass:
		fields = append(fields, "g")
	case nodes.Package + "|":
		fields = append(fields, "p")
	case nodes.Record + "|" + nodes.Class:
		fields = append(fields, "c")
	case nodes.Variable + "|":
		fields = append(fields, "v")
	}

	for parent := range emap[ticket][edges.ChildOf] {
		parentIdent := string(nmap[parent][identifierFact])
		if parentIdent == "" {
			continue
		}
		switch string(nmap[parent][facts.NodeKind]) + "|" + string(nmap[parent][facts.Subkind]) {
		case nodes.Function + "|":
			fields = append(fields, "function:"+parentIdent)
		case nodes.Record + "|" + nodes.Class:
			fields = append(fields, "class:"+parentIdent)
		case nodes.Enum + "|" + nodes.EnumClass:
			fields = append(fields, "enum:"+parentIdent)
		}
	}

	return fields, nil
}

func offsetLine(text []byte, offset int) int {
	return bytes.Count(text[:offset], []byte("\n")) + 1
}
