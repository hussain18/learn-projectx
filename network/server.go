package network

import (
	"fmt"
	"time"

	"github.com/hussain18/learn-projectx/core"
	"github.com/hussain18/learn-projectx/crypto"
	"github.com/sirupsen/logrus"
)

type ServerOpts struct {
	Transports []Transport
	PrivateKey *crypto.PrivateKey
	BlockTime  time.Duration
}

type Server struct {
	ServerOpts
	rpcCh       chan RPC
	mempool     TxPool
	blockTime   time.Duration
	quitCh      chan struct{}
	isValidator bool
}

func NewServer(opts ServerOpts) *Server {
	return &Server{
		ServerOpts:  opts,
		isValidator: opts.PrivateKey != nil,
		mempool:     *NewTxPool(),
		blockTime:   opts.BlockTime,
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.blockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			if s.isValidator {
				// exec consensus
				s.createNewBlock()
			} else {
				fmt.Println("Do things every x seconds")
			}
		}
	}

	fmt.Println("Server shutdown")
}

func (s *Server) handleTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}

	txHash := tx.Hash(core.TxHasher{})

	if s.mempool.Has(txHash) {

		logrus.WithField("NEW_TX", logrus.Fields{
			"hash": txHash,
		}).Info("transaction already in mempool")

		return nil
	}

	logrus.WithField("NEW_TX", logrus.Fields{
		"hash": txHash,
	}).Info("adding new transaction to mempool")

	return s.mempool.Add(tx)
}

func (s *Server) createNewBlock() error {
	fmt.Println("creating new block")
	return nil
}

func (s *Server) initTransports() {
	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
