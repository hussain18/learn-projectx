package network

import (
	"fmt"
	"time"

	"github.com/hussain18/learn-projectx/core"
	"github.com/hussain18/learn-projectx/crypto"
	"github.com/sirupsen/logrus"
)

var defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	RPCHandler RPCHandler
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
	if opts.BlockTime == time.Duration(0) {
		opts.BlockTime = defaultBlockTime
	}

	s := &Server{
		ServerOpts:  opts,
		isValidator: opts.PrivateKey != nil,
		mempool:     *NewTxPool(),
		blockTime:   opts.BlockTime,
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}, 1),
	}

	if opts.RPCHandler == nil {
		s.RPCHandler = NewDefaultRPCHandler(s)
	}

	return s
}

func (s *Server) Start() {
	s.initTransports()
	ticker := time.NewTicker(s.blockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			if err := s.RPCHandler.HandleRPC(rpc); err != nil {
				logrus.Error(err)
			}
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

func (s *Server) ProcessTransaction(from NetAddr, tx *core.Transaction) error {
	txHash := tx.Hash(core.TxHasher{})

	if s.mempool.Has(txHash) {

		logrus.WithField("NEW_TX", logrus.Fields{
			"hash": txHash,
		}).Info("transaction already in mempool")

		return nil
	}

	if err := tx.Verify(); err != nil {
		return err
	}

	logrus.WithField("NEW_TX", logrus.Fields{
		"hash":           txHash,
		"mempool length": s.mempool.Len(),
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
