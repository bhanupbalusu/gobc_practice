package blockchain

type Block struct {
	PrevHash []byte
	Data     []byte
	Hash     []byte
	Nonce    int
}

type Blockchain struct {
	Blocks []*Block
}

func NewBlock(data string, prevHash []byte) *Block {
	// Initialize the block with data and prevHash
	block := &Block{prevHash, []byte(data), []byte{}, 0}

	// Create New Proof Of Work for the block which sets the Target field
	pow := NewProofOfWork(block)

	// Run the Proof of Work to get nonce and new Hash
	nonce, hash := pow.RunProofOfWork()

	// Assign hash and nonce to the block
	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

func GenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func (chain *Blockchain) AddBlock(data string) {
	// Get previous block
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	// Create new block
	block := NewBlock(data, prevBlock.Hash)

	// Append new block to the existing blocks
	chain.Blocks = append(chain.Blocks, block)
}

func InitBlockchain() *Blockchain {
	// Create Genesis Block and initiate the List
	block := GenesisBlock()
	return &Blockchain{[]*Block{block}}
}
