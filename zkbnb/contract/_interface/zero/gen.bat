@echo off

abigen --bin=_MyERC20_sol_ZkBNBRelatedERC20.bin --abi=_MyERC20_sol_ZkBNBRelatedERC20.abi --pkg=erc20 --out=ERC20Gen.go && abigen --bin=_Governance_sol_Governance.bin --abi=_Governance_sol_Governance.abi --pkg=governance --out=GovernanceGen.go && abigen --bin=_Governance_sol_Governance.bin --abi=_AssetGovernance_sol_AssetGovernance.abi --pkg=assetGovernance --out=AssetGovernanceGen.go && abigen --bin=_ZkBNBVerifier_sol_ZkBNBVerifier.bin --abi=_ZkBNBVerifier_sol_ZkBNBVerifier.abi --pkg=verifier --out=ZkBNBVerifierGen.go && abigen --bin=_ZkBNB_sol_ZkBNB.bin --abi=_ZkBNB_sol_ZkBNB.abi --pkg=zkbnb --out=ZkBNBGen.go