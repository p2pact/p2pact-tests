package bench

import p2pact "github.com/p2pact/p2pact-go"

// UseDict gives every P2PACT node the common dictionary and lets it train its own, as a device
// does; off, P2PACT compresses with plain zstd.
var UseDict = true

func newNode(dir, name string) (*p2pact.Node, error) {
	n, err := p2pact.NewNode(dir, name)
	if err != nil {
		return nil, err
	}
	if !UseDict {
		n.NoLearning = true
		return n, nil
	}
	if err := n.AddDict(p2pact.CommonDict); err != nil {
		n.Close()
		return nil, err
	}
	return n, nil
}
