# merlin-zkproof-verify-demo

### 1. This demo is for call the merlin_getZkProof and merlin_verifyZkProof method

### 2. merlin_getZkProof method 
 * the params is the block number 
 * curl -s http://127.0.0.1:8123 -X POST -H "Content-Type: application/json" --data '{"jsonrpc": "2.0", "method": "merlin_getZkProof", "params":  ["0xcb43d2"], "id": 0 }'

### 3. merlin_verifyZkProof method
 * param1(type: uint64): forkID 
 * param2(type: [24]common.Hash): proof
 * param3(type: [1]*big.Int): pubSignals
 * curl -s http://127.0.0.1:8123 -X POST -H "Content-Type: application/json" --data '{"jsonrpc": "2.0", "method": "merlin_verifyZkProof", "params":[8,["0x20227cbcef731b6cbdc0edd5850c63dc7fbc27fb58d12cd4d08298799cf66a05","0x12c230867d3375a1f4669e7267dad2c31ebcddbaccea6abd67798ceae35ae761","0x1c665b6069339e6812d015e239594aa71c4e217288e374448c358f6459e057c9","0x1ad2ef514570b5dea21508e214430daadabdd23433820000fe98b1c6fa81d5c5","0x12b86fbf87bd7102775f8ef1da7e8014dc7aab225503237c7927c032e589e9a0","0x1a0eab9fda82ffe834c2a4977f36cc9bcb1f2327bdac5fb48ffbeb9656efcdf7","0x0d2656c328903e9fb96e4e3f470c447b3053cc68d68cf0ad317fe10aa7f25422","0x2e47ea07f3c1c3aacb74e5926a67262f261c1ed3120576ab877b49a81fb8aac5","0x1431858662af6b1a8138a44e9d0812d032340369459ccc98b109347cc874c720","0x2dceecc3dbb09d7f9e5658f1ca3a92d22be1fa28f9945205d853e2c866d9b649","0x301ac9857b07b92e4865283d3d5e2b711ea5f85cb2da71965382ece050508d3d","0x008bbe4df5458f70bd3e1bfcc50b34222b43cd28cbe39a3bab6e464664a74216","0x1df99c607638e415ced49d0cd719518539ed5f561f81d07fe40d3ce85508e033","0x2465313e60ad9ae271d580022ffca4fbe4d72d38d18e7a6e20d020a1d1e5a8f4","0x11291ab95521386fa538ddfe6a391d4a3669cc64c40f07895f031550b32f7d73","0x205a69c214a8ef3cdf996c495e3fd24c00873f30ea6b2bfabfd38de1c3da357d","0x1fefe203573fdad22f675cb5cfabbec0a041b1b31274f70193da8e90cfc4d6dc","0x054c7cd26d09c1dadd064ec52b6ddcfa0cb144d65d9e131c0c88f8004f90d363","0x034d839aa7760167b5302c36d2c2f6714b41782070b10c51c178bd923182d285","0x02f36e19b079b190008c46d19c399331fd60b6b6bde898bd1dd0a71ee7ec7ff7","0x124cc3d374846614389e7b5975b77c4059bc42b810673dbb6f8b951e5b636bdf","0x24afd2a3cbe96ce8600e8a79731b4a56c697596e0bff7b73f413bdbc75069b00","0x2b00d713fae8d6450428246f1b794d56717050fdb77bbe094ac2ee6af54a153e","0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e"],[20198128279241534280278478823280144292885007386329693578864221874880227925160]], "id": 0 }'
### 4. Or you can verify the zkproof by yourself, for more details you can see the main.go
