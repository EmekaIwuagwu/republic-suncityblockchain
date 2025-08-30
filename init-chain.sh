#!/bin/bash
# init-chain.sh

CHAIN_ID="rsuncitychain"
MONIKER="render-validator"
KEYRING_BACKEND="file"

# Initialize the chain
rsuncitychaind init $MONIKER --chain-id $CHAIN_ID --home /root/.rsuncitychain

# Add validator key
rsuncitychaind keys add validator --keyring-backend $KEYRING_BACKEND --home /root/.rsuncitychain

# Add genesis account
rsuncitychaind genesis add-genesis-account validator 100000000000sunc --keyring-backend $KEYRING_BACKEND --home /root/.rsuncitychain

# Generate genesis transaction
rsuncitychaind genesis gentx validator 1000000sunc --chain-id $CHAIN_ID --keyring-backend $KEYRING_BACKEND --home /root/.rsuncitychain

# Collect genesis transactions
rsuncitychaind genesis collect-gentxs --home /root/.rsuncitychain

# Validate genesis file
rsuncitychaind genesis validate-genesis --home /root/.rsuncitychain