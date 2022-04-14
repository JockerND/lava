package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/utils"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"github.com/lavanet/lava/x/pairing/types"
)

func (k Keeper) StakeNewEntry(ctx sdk.Context, provider bool, creator string, chainID string, amount sdk.Coin, endpoints []epochstoragetypes.Endpoint, geolocation uint64) error {
	logger := k.Logger(ctx)
	stake_type := func() string {
		if provider {
			return epochstoragetypes.ProviderKey
		}
		return epochstoragetypes.ClientKey
	}
	//TODO: basic validation for chain ID
	specChainID := chainID

	foundAndActive, _, _ := k.specKeeper.IsSpecFoundAndActive(ctx, specChainID)
	if !foundAndActive {
		details := map[string]string{"spec": specChainID}
		return utils.LavaError(ctx, logger, "stake_"+stake_type()+"_spec", details, "spec not found or not active")
	}
	var minStake sdk.Coin
	if provider {
		minStake = k.GetMinStakeProvider(ctx)
	} else {
		minStake = k.GetMinStakeClient(ctx)
	}
	//if we get here, the spec is active and supported
	if amount.IsLT(minStake) {
		details := map[string]string{"spec": specChainID, stake_type(): creator, "stake": amount.String(), "minStake": minStake.String()}
		return utils.LavaError(ctx, logger, "stake_"+stake_type()+"_amount", details, "insufficient "+stake_type()+" stake amount")
	}
	senderAddr, err := sdk.AccAddressFromBech32(creator)
	if err != nil {
		details := map[string]string{stake_type(): creator, "error": err.Error()}
		return utils.LavaError(ctx, logger, "stake_"+stake_type()+"_addr", details, "invalid "+stake_type()+" address")
	}
	//define the function here for later use
	verifySufficientAmountAndSendToModule := func(ctx sdk.Context, k Keeper, addr sdk.AccAddress, neededAmount sdk.Coin) (bool, error) {
		if k.bankKeeper.GetBalance(ctx, addr, "stake").IsLT(neededAmount) {
			return false, fmt.Errorf("insufficient balance for staking %s current balance: %s", neededAmount, k.bankKeeper.GetBalance(ctx, addr, "stake"))
		}
		err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, []sdk.Coin{neededAmount})
		if err != nil {
			return false, fmt.Errorf("invalid transfer coins to module, %s", err)
		}
		return true, nil
	}
	if provider {
		for _, endpoint := range endpoints {
			//TODO: validate endpoints for spec types and geolocation
			_ = endpoint
		}
	}

	// new staking takes effect from the next block
	blockDeadline := uint64(ctx.BlockHeight()) + 1

	existingEntry, entryExists, indexInStakeStorage := k.epochStorageKeeper.StakeEntryByAddress(ctx, stake_type(), chainID, senderAddr)
	if entryExists {
		//modify the entry
		if existingEntry.Address != creator {
			details := map[string]string{"spec": specChainID, stake_type(): senderAddr.String(), stake_type(): creator}
			utils.LavaError(ctx, logger, "stake_"+stake_type()+"_panic", details, "returned stake entry by address doesn't match sender address!")
		}
		details := map[string]string{"spec": specChainID, stake_type(): senderAddr.String(), "deadline": strconv.FormatUint(blockDeadline, 10), "stake": amount.String()}
		if existingEntry.Stake.IsLT(amount) {
			// increasing stake is allowed
			valid, err := verifySufficientAmountAndSendToModule(ctx, k, senderAddr, amount.Sub(existingEntry.Stake))
			if !valid {
				details["error"] = err.Error()
				details["neededStake"] = amount.Sub(existingEntry.Stake).String()
				return utils.LavaError(ctx, logger, "stake_"+stake_type()+"_update_amount", details, "insufficient funds to pay for difference in stake")
			}
			//paid the difference to module
			existingEntry.Stake = amount
			//we dont change deadlines and chain once they are set, if they need to change, unstake first
			existingEntry.Geolocation = geolocation
			existingEntry.Endpoints = endpoints
			k.epochStorageKeeper.ModifyStakeEntry(ctx, stake_type(), chainID, existingEntry, indexInStakeStorage)
			utils.LogLavaEvent(ctx, logger, stake_type()+"_stake_update", details, "Changing Staked "+stake_type())
			return nil
		}
		details["existingStake"] = existingEntry.Stake.String()
		return utils.LavaError(ctx, logger, "stake_"+stake_type()+"_stake", details, "can't decrease stake for existing "+stake_type())
	}

	// entry isn't staked so add him
	details := map[string]string{"spec": specChainID, stake_type(): senderAddr.String(), "deadline": strconv.FormatUint(blockDeadline, 10), "stake": amount.String()}
	valid, err := verifySufficientAmountAndSendToModule(ctx, k, senderAddr, amount)
	if !valid {
		details["error"] = err.Error()
		return utils.LavaError(ctx, logger, "stake_"+stake_type()+"_new_amount", details, "insufficient amount to pay for stake")
	}

	stakeEntry := epochstoragetypes.StakeEntry{Stake: amount, Address: creator, Deadline: blockDeadline, Endpoints: endpoints, Geolocation: geolocation, Chain: chainID}
	k.epochStorageKeeper.AppendStakeEntry(ctx, stake_type(), chainID, stakeEntry)
	utils.LogLavaEvent(ctx, logger, stake_type()+"_stake_new", details, "Adding Staked "+stake_type())
	return nil
}
