package main

import (
	. "github.com/lukechampine/uint128"
)

// Due to the use of a non builtin type (uint128), this all may need to be moved into a var block, which is more flexible anyway, but mutable.
const (
	SLOT_BLOCKS_TARGET             uint32 = 32  // How many blocks to target per sub-slot
	MIN_BLOCKS_PER_CHALLENGE_BLOCK uint8  = 16  // How many blocks must be created per slot (to make challenge sb)
	MAX_SUB_SLOT_BLOCKS            uint32 = 128 /* Max number of blocks that can be infused into a sub-slot.
	Note: this must be less than SUB_EPOCH_BLOCKS/2, and > SLOT_BLOCKS_TARGET */
	NUM_SPS_SUB_SLOT uint32 = 64 // The number of signage points per sub-slot (including the 0th sp at the sub-slot start)

	SUB_SLOT_ITERS_STARTING    uint64  = 1 << 27 // The sub_slot_iters for the first epoch
	DIFFICULTY_CONSTANT_FACTOR Uint128 = 1 << 67 // Multiplied by the difficulty to get iterations
	DIFFICULTY_STARTING        uint64  = 7       // The difficulty for the first epoch

	DIFFICULTY_CHANGE_MAX_FACTOR uint32 = 3    // The maximum factor by which difficulty and sub_slot_iters can change per epoch
	SUB_EPOCH_BLOCKS             uint32 = 384  // The number of blocks per sub-epoch
	EPOCH_BLOCKS                 uint32 = 4608 // The number of blocks per sub-epoch, must be a multiple of SUB_EPOCH_BLOCKS

	SIGNIFICANT_BITS             int = 8    // The number of bits to look at in difficulty and min iters. The rest are zeroed
	DISCRIMINANT_SIZE_BITS       int = 1024 // Max is 1024 (based on ClassGroupElement int size)
	NUMBER_ZERO_BITS_PLOT_FILTER int = 9    // H(plot id + challenge hash + signage point) must start with these many zeroes
	MIN_PLOT_SIZE                int = 32
	MAX_PLOT_SIZE                int = 50
	SUB_SLOT_TIME_TARGET         int = 600    // The target number of seconds per sub-slot
	NUM_SP_INTERVALS_EXTRA       int = 3      // The difference between signage point and infusion point (plus required_iters)
	MAX_FUTURE_TIME              int = 5 * 60 // The next block can have a timestamp of at most these many seconds more
	NUMBER_OF_TIMESTAMPS         int = 11     /* Than the average of the last NUMBER_OF_TIMESTAMPS blocks
	Used as the initial cc rc challenges, as well as first block back pointers, and first SES back pointer
	We override this value based on the chain being run (testnet0, testnet1, mainnet, etc) */
	GENESIS_CHALLENGE byte = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

	AGG_SIG_ME_ADDITIONAL_DATA          byte = "ccd5bb71183532bff220ba46c268991a3ff07eb358e8255a65c30a2dce0e5fbb" // Forks of chia should change this value to provide replay attack protection
	GENESIS_PRE_FARM_POOL_PUZZLE_HASH   byte = "d23da14695a188ae5708dd152263c4db883eb27edeb936178d4d988b8f3ce5fc" // The block at height must pay out to this pool puzzle hash
	GENESIS_PRE_FARM_FARMER_PUZZLE_HASH byte = "3d8765d3a597ec1d99663f6c9816d915b9f68613ac94009884c4addaefcce6af" // The block at height must pay out to this farmer puzzle hash
	MAX_VDF_WITNESS_SIZE                int  = 64                                                                 // The maximum number of classgroup elements within an n-wesolowski proof

	MEMPOOL_BLOCK_BUFFER int = 50 // Size of mempool = 10x the size of block

	MAX_COIN_AMOUNT int = (1 << 64) - 1 // Max coin amount uint(1 << 64). This allows coin amounts to fit in 64 bits. This is around 18M chia.

	MAX_BLOCK_COST_CLVM int = 11000000000 // Max block cost in clvm cost units

	COST_PER_BYTE int = 12000 // Cost per byte of generator program

	WEIGHT_PROOF_THRESHOLD       uint8  = 2
	WEIGHT_PROOF_RECENT_BLOCKS   uint32 = 1000
	MAX_BLOCK_COUNT_PER_REQUESTS uint32 = 32
	BLOCKS_CACHE_SIZE            uint32 = 4608 + 128*4
	NETWORK_TYPE                 int    = 0
	MAX_GENERATOR_SIZE           uint32 = 1000000
	MAX_GENERATOR_REF_LIST_SIZE  uint32 = 512
	POOL_SUB_SLOT_ITERS          uint64 = 37600000000
)
