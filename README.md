# smart-contracts-with-go
A simple example of how to deploy and interact with ETH smart contracts using Go on a simulated Blockchain.

## Prerequisites

* [solc](http://solidity.readthedocs.io/en/develop/installing-solidity.html)
* geth (go-ethereum)

```bash
go get github.com/ethereum/go-ethereum
cd $GOPATH/src/github.com/ethereum/go-ethereum/
make
make devtools
```

## Generating contract.go

```bash
abigen --sol=Contract.sol --pkg=main --out=contract.go
```

abigen --sol=FomoETC.sol --pkg=FomoETC --out=FomoETC.go

abigen --sol=erc721.sol --pkg=erc721 --out=erc721.go
abigen --sol=/root/gosmartcontract/erc721/tokens/nf-token.sol --pkg=erc721 --out=erc721.go
solc -o . --bin --ast --asm --allow-paths /root/gosmartcontract/erc721,/root/gosmartcontract/erc721/math,/root/gosmartcontract/erc721/ownership,/root/gosmartcontract/erc721/tokens,/root/gosmartcontract/erc721/utils nf-token.sol 


solc -o . --bin --ast --asm --allow-paths /root/gosmartcontract/erc721 /root/gosmartcontract/erc721/tokens/nf-token.sol   /////ok
solc --bin --abi --optimize -o . --allow-paths /root/gosmartcontract/erc721 /root/gosmartcontract/erc721/mocks/nf-token-mock.sol   /////ok

solc --bin --abi --optimize -o . erc721.sol


abigen -abi NFTokenMock.abi -bin NFTokenMock.bin --pkg=erc721 --out=erc721.go   ///ok

abigen --sol=testinterface.sol --pkg=testinterface --out=testinterface.go
abigen --sol=etclottery.sol --pkg=main --out=etclottery.go
abigen --sol=calletclottery.sol --pkg=main --out=calletclottery.go
## Running

```bash
go build . && ./smart-contracts-with-go
```

sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc

F3DexternalSettingsInterface constant private extSettings = F3DexternalSettingsInterface(0x32967D6c142c2F38AB39235994e2DDF11c37d590);

0x6080604052600436106100535763ffffffff60e060020a600035041663114719c5811461005857806318d0376c1461007f5780631a9be331146100945780634ccbe888146100a9578063c105c5ed146100be575b600080fd5b34801561006457600080fd5b5061006d6100d5565b60408051918252519081900360200190f35b34801561008b57600080fd5b5061006d6100ed565b3480156100a057600080fd5b5061006d610104565b3480156100b557600080fd5b5061006d61011b565b3480156100ca57600080fd5b506100d3610132565b005b6000805460ff16156100e657600080fd5b50610e1090565b6000805460ff16156100fe57600080fd5b50603c90565b6000805460ff161561011557600080fd5b50607890565b6000805460ff161561012c57600080fd5b50600f90565b6000805460ff191660011790555600a165627a7a72305820bef27da50bba901119b988fb734dcd5734d21efbd6a773b0a9b165c2d523e7230029

playerbook
0x45beee899f18e29ca5bbf38206cb28ab1398e883
JIincForwarderBin
0xed02aacb4c1f924da723e9ddac7e14dae07b8cdc
divies
0x81f8380c141fd2feb9855724c22ca55f5178fde3


abigen --bin="0x6080604052600436106100535763ffffffff60e060020a600035041663114719c5811461005857806318d0376c1461007f5780631a9be331146100945780634ccbe888146100a9578063c105c5ed146100be575b600080fd5b34801561006457600080fd5b5061006d6100d5565b60408051918252519081900360200190f35b34801561008b57600080fd5b5061006d6100ed565b3480156100a057600080fd5b5061006d610104565b3480156100b557600080fd5b5061006d61011b565b3480156100ca57600080fd5b506100d3610132565b005b6000805460ff16156100e657600080fd5b50610e1090565b6000805460ff16156100fe57600080fd5b50603c90565b6000805460ff161561011557600080fd5b50607890565b6000805460ff161561012c57600080fd5b50600f90565b6000805460ff191660011790555600a165627a7a72305820bef27da50bba901119b988fb734dcd5734d21efbd6a773b0a9b165c2d523e7230029" --pkg=main --out=test.go


./porosityexe --code 0x6080604052600436106100535763ffffffff60e060020a600035041663114719c5811461005857806318d0376c1461007f5780631a9be331146100945780634ccbe888146100a9578063c105c5ed146100be575b600080fd5b34801561006457600080fd5b5061006d6100d5565b60408051918252519081900360200190f35b34801561008b57600080fd5b5061006d6100ed565b3480156100a057600080fd5b5061006d610104565b3480156100b557600080fd5b5061006d61011b565b3480156100ca57600080fd5b506100d3610132565b005b6000805460ff16156100e657600080fd5b50610e1090565b6000805460ff16156100fe57600080fd5b50603c90565b6000805460ff161561011557600080fd5b50607890565b6000805460ff161561012c57600080fd5b50600f90565b6000805460ff191660011790555600a165627a7a72305820bef27da50bba901119b988fb734dcd5734d21efbd6a773b0a9b165c2d523e7230029 --decompile --verbose 0
Porosity v0.1 (https://www.comae.io)
Matt Suiche, Comae Technologies <support@comae.io>
The Ethereum bytecode commandline decompiler.
Decompiles the given Ethereum input bytecode and outputs the Solidity code.

Hash: 0x114719C5
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: CALLDATASIZE
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: CALLDATASIZE
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: CALLDATASIZE
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: CALLDATASIZE
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
function func_114719c5 {
      if (!msg.value) {
      }
}


LOC: 4
Hash: 0x18D0376C
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
function func_18d0376c {
      if (!msg.value) {
      }
}


LOC: 4
Hash: 0x1A9BE331
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
function func_1a9be331 {
      if (!msg.value) {
      }
}


LOC: 4
Hash: 0x4CCBE888
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
executeInstruction: NOT_IMPLEMENTED: REVERT
function func_4ccbe888 {
      if (!msg.value) {
      }
}


LOC: 4
Hash: 0xC105C5ED
executeInstruction: NOT_IMPLEMENTED: REVERT
function func_c105c5ed {
      if (!msg.value) {
      }
      store[var_f9TrE] = 0x1;
      return;
      return;
}


https://etherscan.io/address/0x4c7b8591c50f4ad308d07d6294f2945e074420f5#code

https://etherscan.io/address/0x4b7ac91b53545ae20a4990f9b5f6a14682deecbc#code namefilter

https://etherscan.io/address/0xd4b5556dad4a0affc0eef0da56a928712e412180#code SafeMath

https://etherscan.io/address/0xbf57726f133e0e57896a52d3baf377d2bf91f5b1#code 
0x6080604052600080fd00a165627a7a72305820777e5f1e0da4d9cc37b745b36a9f7081550f3b07c869709c3ff373b0e8832e640029 未知

https://etherscan.io/address/0x1fb5464720ac6610da724fd9b65a08837bea51d3#code
0x6080604052600080fd00a165627a7a72305820cb6e27df00eea697bb2571ddf15ab37f12bec6c6aa6934253dd6fa1bd7eadc1b0029 未知
https://etherscan.io/address/0x27afcbe78ba41543c8e6ede1ec0560cd128adccb#code
0x6080604052600436106100535763ffffffff60e060020a600035041663114719c5811461005857806318d0376c1461007f5780631a9be331146100945780634ccbe888146100a9578063c105c5ed146100be575b600080fd5b34801561006457600080fd5b5061006d6100d5565b60408051918252519081900360200190f35b34801561008b57600080fd5b5061006d6100ed565b3480156100a057600080fd5b5061006d610104565b3480156100b557600080fd5b5061006d61011b565b3480156100ca57600080fd5b506100d3610132565b005b6000805460ff16156100e657600080fd5b50610e1090565b6000805460ff16156100fe57600080fd5b50603c90565b6000805460ff161561011557600080fd5b50607890565b6000805460ff161561012c57600080fd5b50600f90565b6000805460ff191660011790555600a165627a7a72305820bef27da50bba901119b988fb734dcd5734d21efbd6a773b0a9b165c2d523e7230029 未知

https://etherscan.io/address/0xd14f5d11fde8f2baf394d3334df13ee6aa58c708#code
0x73d14f5d11fde8f2baf394d3334df13ee6aa58c70830146080604052600080fd00a165627a7a7230582064d9e362d23c96e5fef8bd6634582f59eb7da13a151306cfe872808e68990ad80029 未知