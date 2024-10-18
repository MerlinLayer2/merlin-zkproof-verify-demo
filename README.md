# merlin-zkproof-verify-demo

### 1. This demo is for call the merlin_getZkProof and merlin_verifyZkProof method

### 2. The merlin_getZkProof method 
 * param1(type: string): block number or "latest" 
 * param2(type: bool): true( need with raw PubSignals data); false(no need with raw PubSignals data)
 * curl -s http://127.0.0.1:8123 -X POST -H "Content-Type: application/json" --data '{"jsonrpc": "2.0", "method": "merlin_getZkProof", "params":  ["0xcb43d2", true], "id": 0 }'
```
{
    "forkID": 8,
    "proof": [
        "0x2df047ac570cad92ed2826eff6a2ccd87e4dfce8ff474631f15f04a1e646ed09",
        "0x0c25b53181477369f5371991cc334febb1a3f331fce84e53c269359c14c5a682",
        "0x22b1ae66ff60e762406c77a8af919bb72a3febce6331637d2e5d456f63a27494",
        "0x0299b856937ee9d7ca42140564aea40de1b162740896f8733d55c8dfb6466bc4",
        "0x0e4f333afbd4e4232ac2086f66d3e9459d84d3935df5f9f4bdd7a3bcb257fc62",
        "0x0d8fd4affa1978906950ba78a100bcdb347617eb891d559accc239e058412aee",
        "0x0faa3bffff2ed911a9c6268890168ee620588c968c1f95d65db5f87901e4e706",
        "0x114137daed48c6783c879cf8ea68b82e56cf574e7d514c3816d7d327b1ebf166",
        "0x2a09abcd6ce09923731f3d630ce9739df49f59816ed7a5ddd4229e9bf06e3d64",
        "0x18557d9ecd0075c8ddc58581828e63d17540bba3f8d7fd48e9beeaaeb3b8db86",
        "0x0f775e8a8840ea9043c165885a417c0d867a9ea01b249372dc7011c25f3cfaaa",
        "0x23472c3aaf4c5ff8c3bd74e82957e47ef403b94c0a940dc1ead45c97b675740c",
        "0x04e4d43f25d35e40986e8d5bf64d350ac75b8fcc4be56d178b2bb0f1a03a2895",
        "0x2deea19a2f70a854a2feb0741e5b7316c38e13a20c501eb333a1ca947daec1c1",
        "0x10e9bffa78b3b1d8b7ca454ad8101730196eca983641c35613b581e08c8d9ad9",
        "0x0b61fa7e109edd3320a0e4a890fbe0c9db220f84d9cc146e8be89640d2a3e493",
        "0x0b2cc5241aed152b66c50a194f2bc114111f330b33434df167d4440e14cb7292",
        "0x097bc1ae33c486a880b2a6218bb8b77cb9f3da5909347bb16789b02fcb0cd32e",
        "0x0017a5ed0bbd778a892e241c19370d3a01fb590d94c30c2f3ef2df7fd6c1d3de",
        "0x0263a34b3fd611b307cc7bc3264910c2370189f1ab375d42af88ddaee8a3ed13",
        "0x2d3a4c1361af57154d56a89a0be974c06e88cf65212da01f4663261da062d2fe",
        "0x062b01e9d0f35cb048d58e2205b80129b2aad0c81ef21d5ae45fd8b318b6e2c8",
        "0x02cb0662fbaa046b3a6d3eee5badc56916c2fbff8c3784b0895c969e3cd242e5",
        "0x217ffd40454cad9215101b384f86559fdb1bb83497d60e1033486c852c8fd583"
    ],
    "pubSignals": [
        6603510186853169972152213777899593011637911202509582519799533839979742083878
    ],
    "rawPubSignals": {
        "snark": {
            "sender": "0xe76cc099094d484e67cd7b777d22a93afc2920cc",
            "oldStateRoot": "0xbc26b56bbd4fa7c91c97a0e0fea120b7d26eba75daa2cc3035b5edcc2b5c6630",
            "oldAccInputHash": "0xab07cc71710e24d280bcd070abf25eb01b99788c985c9cd3ede196a5e9586672",
            "initNumBatch": 1774053,
            "chainId": 4200,
            "forkID": 8,
            "newStateRoot": "0x97b2f0666edfff8c6eb8315c0161db5a10ae11342ba7f34da46d581bcb70e376",
            "newAccInputHash": "0x0db4014d73587d6ef5f9dfabdc9a14ebafddeee91f6da5fba029f9f84bfd1631",
            "newLocalExitRoot": "0x0000000000000000000000000000000000000000000000000000000000000000",
            "finalNewBatch": 1774057
        },
        "rfield": "21888242871839275222246405745257275088548364400416034343698204186575808495617"
    },
    "startBlockNum": 13321151,
    "endBlockNum": 13321188
}
```
 * curl -s http://127.0.0.1:8123 -X POST -H "Content-Type: application/json" --data '{"jsonrpc": "2.0", "method": "merlin_getZkProof", "params":  ["0xcb43d2", false], "id": 0 }'
```
{
    "forkID": 8,
    "proof": [
        "0x2df047ac570cad92ed2826eff6a2ccd87e4dfce8ff474631f15f04a1e646ed09",
        "0x0c25b53181477369f5371991cc334febb1a3f331fce84e53c269359c14c5a682",
        "0x22b1ae66ff60e762406c77a8af919bb72a3febce6331637d2e5d456f63a27494",
        "0x0299b856937ee9d7ca42140564aea40de1b162740896f8733d55c8dfb6466bc4",
        "0x0e4f333afbd4e4232ac2086f66d3e9459d84d3935df5f9f4bdd7a3bcb257fc62",
        "0x0d8fd4affa1978906950ba78a100bcdb347617eb891d559accc239e058412aee",
        "0x0faa3bffff2ed911a9c6268890168ee620588c968c1f95d65db5f87901e4e706",
        "0x114137daed48c6783c879cf8ea68b82e56cf574e7d514c3816d7d327b1ebf166",
        "0x2a09abcd6ce09923731f3d630ce9739df49f59816ed7a5ddd4229e9bf06e3d64",
        "0x18557d9ecd0075c8ddc58581828e63d17540bba3f8d7fd48e9beeaaeb3b8db86",
        "0x0f775e8a8840ea9043c165885a417c0d867a9ea01b249372dc7011c25f3cfaaa",
        "0x23472c3aaf4c5ff8c3bd74e82957e47ef403b94c0a940dc1ead45c97b675740c",
        "0x04e4d43f25d35e40986e8d5bf64d350ac75b8fcc4be56d178b2bb0f1a03a2895",
        "0x2deea19a2f70a854a2feb0741e5b7316c38e13a20c501eb333a1ca947daec1c1",
        "0x10e9bffa78b3b1d8b7ca454ad8101730196eca983641c35613b581e08c8d9ad9",
        "0x0b61fa7e109edd3320a0e4a890fbe0c9db220f84d9cc146e8be89640d2a3e493",
        "0x0b2cc5241aed152b66c50a194f2bc114111f330b33434df167d4440e14cb7292",
        "0x097bc1ae33c486a880b2a6218bb8b77cb9f3da5909347bb16789b02fcb0cd32e",
        "0x0017a5ed0bbd778a892e241c19370d3a01fb590d94c30c2f3ef2df7fd6c1d3de",
        "0x0263a34b3fd611b307cc7bc3264910c2370189f1ab375d42af88ddaee8a3ed13",
        "0x2d3a4c1361af57154d56a89a0be974c06e88cf65212da01f4663261da062d2fe",
        "0x062b01e9d0f35cb048d58e2205b80129b2aad0c81ef21d5ae45fd8b318b6e2c8",
        "0x02cb0662fbaa046b3a6d3eee5badc56916c2fbff8c3784b0895c969e3cd242e5",
        "0x217ffd40454cad9215101b384f86559fdb1bb83497d60e1033486c852c8fd583"
    ],
    "pubSignals": [
        6603510186853169972152213777899593011637911202509582519799533839979742083878
    ],
    "startBlockNum": 13321151,
    "endBlockNum": 13321188
}
```

### 3. The merlin_verifyZkProof method
 * param1(type: uint64): forkID 
 * param2(type: [24]common.Hash): proof
 * param3(type: [1]*big.Int): pubSignals
 * curl -s http://127.0.0.1:8123 -X POST -H "Content-Type: application/json" --data '{"jsonrpc": "2.0", "method": "merlin_verifyZkProof", "params":[8,["0x20227cbcef731b6cbdc0edd5850c63dc7fbc27fb58d12cd4d08298799cf66a05","0x12c230867d3375a1f4669e7267dad2c31ebcddbaccea6abd67798ceae35ae761","0x1c665b6069339e6812d015e239594aa71c4e217288e374448c358f6459e057c9","0x1ad2ef514570b5dea21508e214430daadabdd23433820000fe98b1c6fa81d5c5","0x12b86fbf87bd7102775f8ef1da7e8014dc7aab225503237c7927c032e589e9a0","0x1a0eab9fda82ffe834c2a4977f36cc9bcb1f2327bdac5fb48ffbeb9656efcdf7","0x0d2656c328903e9fb96e4e3f470c447b3053cc68d68cf0ad317fe10aa7f25422","0x2e47ea07f3c1c3aacb74e5926a67262f261c1ed3120576ab877b49a81fb8aac5","0x1431858662af6b1a8138a44e9d0812d032340369459ccc98b109347cc874c720","0x2dceecc3dbb09d7f9e5658f1ca3a92d22be1fa28f9945205d853e2c866d9b649","0x301ac9857b07b92e4865283d3d5e2b711ea5f85cb2da71965382ece050508d3d","0x008bbe4df5458f70bd3e1bfcc50b34222b43cd28cbe39a3bab6e464664a74216","0x1df99c607638e415ced49d0cd719518539ed5f561f81d07fe40d3ce85508e033","0x2465313e60ad9ae271d580022ffca4fbe4d72d38d18e7a6e20d020a1d1e5a8f4","0x11291ab95521386fa538ddfe6a391d4a3669cc64c40f07895f031550b32f7d73","0x205a69c214a8ef3cdf996c495e3fd24c00873f30ea6b2bfabfd38de1c3da357d","0x1fefe203573fdad22f675cb5cfabbec0a041b1b31274f70193da8e90cfc4d6dc","0x054c7cd26d09c1dadd064ec52b6ddcfa0cb144d65d9e131c0c88f8004f90d363","0x034d839aa7760167b5302c36d2c2f6714b41782070b10c51c178bd923182d285","0x02f36e19b079b190008c46d19c399331fd60b6b6bde898bd1dd0a71ee7ec7ff7","0x124cc3d374846614389e7b5975b77c4059bc42b810673dbb6f8b951e5b636bdf","0x24afd2a3cbe96ce8600e8a79731b4a56c697596e0bff7b73f413bdbc75069b00","0x2b00d713fae8d6450428246f1b794d56717050fdb77bbe094ac2ee6af54a153e","0x2fb8ce1d31a86c4fdd523783b910bedf7db58a46ba6ce48ac3ca194f3cf2275e"],[20198128279241534280278478823280144292885007386329693578864221874880227925160]], "id": 0 }'

### 4. You can verify the zkproof by yourself by delpoyed a verify contract
 * You should delpoy a verifier contract, different forkid have a different verifier contract, forkID5 contract is [here](https://github.com/MerlinLayer2/merlin-cdk-validium-contracts/blob/fork5/contracts/verifiers/FflonkVerifier.sol), forkID8 contrat is [here](https://github.com/MerlinLayer2/merlin-cdk-validium-contracts/blob/fork8/contracts/verifiers/FflonkVerifier.sol)
 * After deploy the verifier contract you can verify the proof

### 5. You can use the `rawPubSignals` data  calculated the `pubSignals` like `func TestGetPubSignals(t *testing.T)` in the file of `pubsignals_test.go`

### 6. For more used detail to see the `main.go`
