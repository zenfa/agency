# Background
## Permission Management with Current Internet Applications
People need to work with others and other organizations. Online collaboration tools empower people to work together on projects. People can invite others to work on the same document collaboratively. Project owners manange read/write permissions to control access to it. In other scenarios, people need to grant permissions to organizations to access their resources, e.g. customers can add credit cards to their PayPal account, so that when they make an online payment, Paypal conduct transactions for them.

##  Risks
https://en.wikipedia.org/wiki/Credit_card_fraud

"The credit card security code is a safeguard against potential fraud and theft. If someone was able to get a hold of your credit card number via skimming or other means, they’d be out of luck at most websites that require a security code to make a purchase. Essentially, the code is designed to indicate the card is in your possession." From https://thepointsguy.com/guide/credit-card-security-code/

# Solution
More secure and it is under users' control.

1. Users can authorize an agent to access their resoureces, e.g. tokens. Moreover, they can set a limit. In our example below, user1 authorizes "agent" to transfer less than 10token at a time.
2. Agents send requests to access users' resources. 
3. Blockchain network will verify the request, if authorization exists with allowed limit. And accept or reject the request accordingly.
4. Users can revoke the authorization anytime.

# Console

```bash
agency@agencys-MacBook-Pro agency % agencycli tx agency authorize-agent $(agencycli keys show agent -a) bank 10token --from user1
{base:{options:{keygenFunc:0x48d8d80 deriveFunc:0x48da4a0 supportedAlgos:[secp256k1] supportedAlgosLedger:[secp256k1]}} db:0xc000f53fb0}


before sign: cosmos1k4k3jfj0t530lejpfpxzs2cvjgywspw524mtzr;bank


user1
12345678
PubKeySecp256k1{031FEE010729E090448CED9FE882EA0EACD53AE8CB569EFAA264C3D95F51D4C115}
N���g��#��s�K�^�r*F�n������`X����	��?&Q6F��w��<'I���+�

{
  "chain_id": "agency",
  "account_number": "2",
  "sequence": "1",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "agency/AuthorizeAgent",
      "value": {
        "principal": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
        "agent": "cosmos1k4k3jfj0t530lejpfpxzs2cvjgywspw524mtzr",
        "scope": "bank",
        "limit": [
          {
            "denom": "token",
            "amount": "10"
          }
        ]
      }
    }
  ],
  "memo": ""
}

confirm transaction before signing and broadcasting [y/N]: y
{
  "height": "0",
  "txhash": "D7A3A836B59FF44786B338EF9BDA1FD0E785DBE654A6C61B9AA5B34358F9D4B9",
  "raw_log": "[]"
}
agency@agencys-MacBook-Pro agency % agencycli q agency list-agency
[
  {
    "principal": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
    "agent": "cosmos1k4k3jfj0t530lejpfpxzs2cvjgywspw524mtzr",
    "scope": "bank",
    "limit": [
      {
        "denom": "token",
        "amount": "10"
      }
    ]
  }
]
agency@agencys-MacBook-Pro agency % agencycli q agency get-agency $(agencycli keys show user1 -a)  $(agencycli keys show agent -a)
{
  "principal": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
  "agent": "cosmos1k4k3jfj0t530lejpfpxzs2cvjgywspw524mtzr",
  "scope": "bank",
  "limit": [
    {
      "denom": "token",
      "amount": "10"
    }
  ]
}
agency@agencys-MacBook-Pro agency % agencycli q account $(agencycli keys show user1 -a)
{
  "type": "cosmos-sdk/Account",
  "value": {
    "address": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
    "coins": [
      {
        "denom": "token",
        "amount": "1000"
      }
    ],
    "public_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "Ax/uAQcp4JBEjO2f6ILqDqzVOujLVp76omTD2V9R1MEV"
    },
    "account_number": "2",
    "sequence": "2"
  }
}
agency@agencys-MacBook-Pro agency % agencycli tx agency agent-action $(agencycli keys show user1 -a) transfer 1token --from agent
{
  "chain_id": "agency",
  "account_number": "3",
  "sequence": "0",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "agency/AgentAction",
      "value": {
        "principal": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
        "agent": "cosmos1k4k3jfj0t530lejpfpxzs2cvjgywspw524mtzr",
        "action": "transfer",
        "amount": [
          {
            "denom": "token",
            "amount": "1"
          }
        ],
        "sig": "Tn+cmu+xZ7HJI5i1c5VL67FevHJGCCpG2G6hlpmW8vpgWKHS6OYJ6+8/JgdRNhZGjud3seM8J0m9EeUfwyuCDg=="
      }
    }
  ],
  "memo": ""
}

confirm transaction before signing and broadcasting [y/N]: y
{
  "height": "0",
  "txhash": "5F67C540031CD6CD1FF99B8218D09BB97C9739C0E005FABA8A20E889300C329F",
  "raw_log": "[]"
}
agency@agencys-MacBook-Pro agency % agencycli q account $(agencycli keys show user1 -a)
{
  "type": "cosmos-sdk/Account",
  "value": {
    "address": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
    "coins": [
      {
        "denom": "token",
        "amount": "999"
      }
    ],
    "public_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "Ax/uAQcp4JBEjO2f6ILqDqzVOujLVp76omTD2V9R1MEV"
    },
    "account_number": "2",
    "sequence": "2"
  }
}
agency@agencys-MacBook-Pro agency % agencycli tx agency agent-action $(agencycli keys show user1 -a) transfer 11token --from agent
{
  "chain_id": "agency",
  "account_number": "3",
  "sequence": "1",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "agency/AgentAction",
      "value": {
        "principal": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
        "agent": "cosmos1k4k3jfj0t530lejpfpxzs2cvjgywspw524mtzr",
        "action": "transfer",
        "amount": [
          {
            "denom": "token",
            "amount": "11"
          }
        ],
        "sig": "Tn+cmu+xZ7HJI5i1c5VL67FevHJGCCpG2G6hlpmW8vpgWKHS6OYJ6+8/JgdRNhZGjud3seM8J0m9EeUfwyuCDg=="
      }
    }
  ],
  "memo": ""
}

confirm transaction before signing and broadcasting [y/N]: y
{
  "height": "0",
  "txhash": "C6F13B8150C5F1A21CE3BC843CFF14D15ACE1A0C7729B7E822979533921F8073",
  "raw_log": "[]"
}

agency@agencys-MacBook-Pro agency % agencycli q account $(agencycli keys show user1 -a)
{
  "type": "cosmos-sdk/Account",
  "value": {
    "address": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
    "coins": [
      {
        "denom": "token",
        "amount": "999"
      }
    ],
    "public_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "Ax/uAQcp4JBEjO2f6ILqDqzVOujLVp76omTD2V9R1MEV"
    },
    "account_number": "2",
    "sequence": "2"
  }
}
agency@agencys-MacBook-Pro agency % agencycli tx agency deauthorize-agent $(agencycli keys show agent -a) --from user1
{
  "chain_id": "agency",
  "account_number": "2",
  "sequence": "2",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "agency/DeauthorizeAgent",
      "value": {
        "principal": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
        "agent": "cosmos1k4k3jfj0t530lejpfpxzs2cvjgywspw524mtzr"
      }
    }
  ],
  "memo": ""
}

confirm transaction before signing and broadcasting [y/N]: y
{
  "height": "0",
  "txhash": "906EF91E76E2E080379CAB81B9CD4B3CC8BC929C1CD68BA885A35CC4253AC445",
  "raw_log": "[]"
}
agency@agencys-MacBook-Pro agency % agencycli q agency list-agency
null
agency@agencys-MacBook-Pro agency % agencycli tx agency agent-action $(agencycli keys show user1 -a) transfer 1token --from agent
{
  "chain_id": "agency",
  "account_number": "3",
  "sequence": "2",
  "fee": {
    "amount": [],
    "gas": "200000"
  },
  "msgs": [
    {
      "type": "agency/AgentAction",
      "value": {
        "principal": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
        "agent": "cosmos1k4k3jfj0t530lejpfpxzs2cvjgywspw524mtzr",
        "action": "transfer",
        "amount": [
          {
            "denom": "token",
            "amount": "1"
          }
        ],
        "sig": "Tn+cmu+xZ7HJI5i1c5VL67FevHJGCCpG2G6hlpmW8vpgWKHS6OYJ6+8/JgdRNhZGjud3seM8J0m9EeUfwyuCDg=="
      }
    }
  ],
  "memo": ""
}

confirm transaction before signing and broadcasting [y/N]: y
{
  "height": "0",
  "txhash": "A9F80CD3606043A446546F36377EE901C12627E6DBD767ACC38A2284C861041E",
  "raw_log": "[]"
}
agency@agencys-MacBook-Pro agency % agencycli q account $(agencycli keys show user1 -a)
{
  "type": "cosmos-sdk/Account",
  "value": {
    "address": "cosmos1hurvs9p957w8stznqxczrkzp6fn9397klnly2u",
    "coins": [
      {
        "denom": "token",
        "amount": "999"
      }
    ],
    "public_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "Ax/uAQcp4JBEjO2f6ILqDqzVOujLVp76omTD2V9R1MEV"
    },
    "account_number": "2",
    "sequence": "3"
  }
}
```
