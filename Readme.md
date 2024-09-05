# Tests
- To test this run the following commands:
    - `go run main.go print` -> there should be no output since the block chain has not been created yet.
    - `go run main.go add -block "first block"` -> add the first block to the chain. Since it's the first block, by default the genesis block will be added.
    - `go run main.go print` -> running again the print command will output the `genesis` block plus the one created in the last step `first block`
    - `go run main.go add -block "send 10 BTC to Jackie"` -> adds a third block.
    - `go run main.go print` -> please see the output below to see how these 3 blocks would look like when calling the print command at this point:
        macbookpro@MBP-de-MacBook blockchain_poc % go run main.go print                             
        badger 2024/09/05 13:41:02 INFO: All 1 tables opened in 1ms
        badger 2024/09/05 13:41:02 INFO: Replaying file id: 0 at offset: 929
        badger 2024/09/05 13:41:02 INFO: Replay took: 5.027µs
        badger 2024/09/05 13:41:02 DEBUG: Value log discard stats empty
        Previous Hash: 00002c3afdd8f7ee21f4dc366c79e3408da23f583847e19bbdc5a25b7e151560 
        Data in block: send 10 BTC to Jackie 
        Hash: 0000079a8366d863c87289fd58f858585a4040b9b512bf2ccb6198685852fda4 
        PoW: true 

        Previous Hash: 0000342dc11a9fd1833ed9fe18ca5627cedc56507de6698acfcafd301398cb35 
        Data in block: first block 
        Hash: 00002c3afdd8f7ee21f4dc366c79e3408da23f583847e19bbdc5a25b7e151560 
        PoW: true 

        Previous Hash:  
        Data in block: Genesis 
        Hash: 0000342dc11a9fd1833ed9fe18ca5627cedc56507de6698acfcafd301398cb35 
        PoW: true 

        badger 2024/09/05 13:41:02 INFO: Got compaction priority: {level:0 score:1.73 dropPrefixes:[]}