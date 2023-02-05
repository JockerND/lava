package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	common "github.com/lavanet/lava/common"
	commontypes "github.com/lavanet/lava/common/types"
	"github.com/lavanet/lava/utils"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"github.com/lavanet/lava/x/packagemanager/types"
)

// SetPackageVersionsStorage set a specific packageVersionsStorage in the store from its index
func (k Keeper) SetPackageVersionsStorage(ctx sdk.Context, packageVersionsStorage types.PackageVersionsStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PackageVersionsStorageKeyPrefix))
	b := k.cdc.MustMarshal(&packageVersionsStorage)
	store.Set(types.PackageVersionsStorageKey(
		packageVersionsStorage.PackageIndex,
	), b)
}

// GetPackageVersionsStorage returns a packageVersionsStorage from its index
func (k Keeper) GetPackageVersionsStorage(
	ctx sdk.Context,
	packageIndex string,

) (val types.PackageVersionsStorage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PackageVersionsStorageKeyPrefix))

	b := store.Get(types.PackageVersionsStorageKey(
		packageIndex,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePackageVersionsStorage removes a packageVersionsStorage from the store
func (k Keeper) RemovePackageVersionsStorage(
	ctx sdk.Context,
	packageIndex string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PackageVersionsStorageKeyPrefix))
	store.Delete(types.PackageVersionsStorageKey(
		packageIndex,
	))
}

// GetAllPackageVersionsStorage returns all packageVersionsStorage
func (k Keeper) GetAllPackageVersionsStorage(ctx sdk.Context) (list []types.PackageVersionsStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PackageVersionsStorageKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.PackageVersionsStorage
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// Function to get a package by package index + epoch
func (k Keeper) getPackageByEpoch(ctx sdk.Context, packageIndex string, epoch uint64) (*types.Package, error) {
	// get packageVersionsStorage
	packageVersionsStorage, found := k.GetPackageVersionsStorage(ctx, packageIndex)
	if !found {
		return nil, utils.LavaError(ctx, k.Logger(ctx), "get_package_versions_storage", map[string]string{"packageIndex": packageIndex}, "could not get packageVersionsStorage with package index")
	}

	// get packageStorage
	packageStorage := packageVersionsStorage.GetPackageStorage()

	// get packageFixationEntry
	packageFixationEntry, _, found := common.GetFixatedEntry(packageStorage.GetEntryList(), epoch)
	if !found {
		return nil, utils.LavaError(ctx, k.Logger(ctx), "get_fixated_entry", map[string]string{"packageIndex": packageIndex, "epoch": strconv.FormatUint(epoch, 10)}, "could not get packageFixationEntry from packageVersionsStorage's entry list")
	}

	unmarshaledPackage, err := k.getPackageFromPackageFixationEntry(ctx, packageFixationEntry)
	if err != nil {
		return nil, utils.LavaError(ctx, k.Logger(ctx), "get_package_from_package_entry", map[string]string{"err": err.Error()}, "could not get unmarshal data from packageFixationEntry to get package object")
	}

	return unmarshaledPackage, nil
}

// Function to get a package's latest version
func (k Keeper) getPackageLatestVersion(ctx sdk.Context, packageIndex string) (*types.Package, error) {
	// get packageVersionsStorage
	packageVersionsStorage, found := k.GetPackageVersionsStorage(ctx, packageIndex)
	if !found {
		return nil, utils.LavaError(ctx, k.Logger(ctx), "get_package_versions_storage", map[string]string{"packageIndex": packageIndex}, "could not get packageVersionsStorage with package index")
	}

	// get the package fixation entry's latest version
	latestVersionPackageFixationEntry := common.GetLatestFixatedEntry(packageVersionsStorage.GetPackageStorage().GetEntryList())

	// get the package from the package fixation entry
	latestVersionPackage, err := k.getPackageFromPackageFixationEntry(ctx, latestVersionPackageFixationEntry)
	if err != nil {
		return nil, utils.LavaError(ctx, k.Logger(ctx), "get_package_from_package_entry", map[string]string{"err": err.Error()}, "could not get unmarshal data from packageFixationEntry to get package object")
	}

	return latestVersionPackage, nil
}

// Function to extract and unmarshal the package object from a package fixation entry
func (k Keeper) getPackageFromPackageFixationEntry(ctx sdk.Context, packageFixationEntry *commontypes.Entry) (*types.Package, error) {
	// get the package object's marshalled data from the package fixation entry
	marshaledPackage, err := common.GetMarshaledDataFromEntry(ctx, packageFixationEntry)
	if err != nil {
		return nil, utils.LavaError(ctx, k.Logger(ctx), "get_marshaled_data_from_entry", map[string]string{"err": err.Error()}, "could not get marshaled data from packageFixationEntry")
	}

	// unmarshal the marshaled data to get the package object
	var unmarshaledPackage *types.Package
	err = k.cdc.Unmarshal(marshaledPackage, unmarshaledPackage)
	if err != nil {
		return nil, utils.LavaError(ctx, k.Logger(ctx), "unmarshal_package", map[string]string{"err": err.Error()}, "could not unmarshal package")
	}

	return unmarshaledPackage, nil
}

// Function to add a new package to the packageVersionsStorage. It supports addition of packages with new index and packages with existing index (index that is already saved in the storage)
func (k Keeper) addNewPackageToStorage(ctx sdk.Context, packageToAdd *types.Package) error {
	// get current epoch
	currentEpoch := k.epochStorageKeeper.GetEpochStart(ctx)

	// verify package's fields
	err := k.verifyNewPackage(ctx, packageToAdd, currentEpoch)
	if err != nil {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_new_package", map[string]string{"err": err.Error()}, "one or more fields of the new package hold invalid values")
	}

	// marshal the package
	marshaledPackage, err := k.cdc.Marshal(packageToAdd)
	if err != nil {
		return utils.LavaError(ctx, k.Logger(ctx), "marshal_new_package", map[string]string{"err": err.Error()}, "could not marshal package")
	}

	// create a new packageVersionEntry
	packageFixationEntryToAdd, err := common.CreateNewFixatedEntry(ctx, currentEpoch, marshaledPackage)
	if err != nil {
		return utils.LavaError(ctx, k.Logger(ctx), "create_new_fixated_entry", map[string]string{"err": err.Error()}, "could not create fixated entry")
	}

	// check if there is a package storage for the new package's index
	packageVersionsStorage, found := k.GetPackageVersionsStorage(ctx, packageToAdd.GetIndex())
	if !found {
		// packageVersionsStorage not found -> create a new entryStorage which consists the new packageFixationEntryToAdd
		entryStorage, err := common.CreateNewFixatedEntryStorage(ctx, packageFixationEntryToAdd)
		if err != nil {
			return utils.LavaError(ctx, k.Logger(ctx), "create_new_fixated_entry_storage", map[string]string{"err": err.Error()}, "could not create fixated entry storage")
		}

		// create a new packageVersionsStorage
		packageVersionsStorage = types.PackageVersionsStorage{PackageIndex: packageToAdd.GetIndex(), PackageStorage: entryStorage}
	} else {
		// packageVersionsStorage found -> get a new entry list with the new entry preprended to it
		updatedEntryList, err := common.SetFixatedEntry(ctx, packageVersionsStorage.GetPackageStorage().GetEntryList(), packageFixationEntryToAdd, true)
		if err != nil {
			return utils.LavaError(ctx, k.Logger(ctx), "set_fixated_entry", map[string]string{"err": err.Error()}, "could not set package storage")
		}

		// update the packageVersionsStorage's entry list
		packageVersionsStorage.PackageStorage.EntryList = updatedEntryList
	}

	// update the KVStore with the new packageVersionsStorage
	k.SetPackageVersionsStorage(ctx, packageVersionsStorage)

	return nil
}

// Function to verify a package object fields
func (k Keeper) verifyNewPackage(ctx sdk.Context, packageToCheck *types.Package, epochToAddPackage uint64) error {
	// check that the epoch field inside the package object matches the epoch that the package is going to be added on
	if packageToCheck.GetEpoch() != epochToAddPackage {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_package_fields", map[string]string{"packageEpoch": strconv.FormatUint(packageToCheck.GetEpoch(), 10), "epochToAddPackage": strconv.FormatUint(epochToAddPackage, 10)}, "the package's epoch field does not match the epoch it's supposed to be added on")
	}

	// check that the package's duration is non-zero
	if packageToCheck.GetDuration() == 0 {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_package_fields", map[string]string{"packageDuration": strconv.FormatUint(packageToCheck.GetDuration(), 10)}, "the package's duration is zero")
	}

	// check that the package's price is non-zero
	if packageToCheck.GetPrice() == sdk.NewCoin(epochstoragetypes.TokenDenom, sdk.ZeroInt()) {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_package_fields", map[string]string{"packagePrice": strconv.FormatUint(packageToCheck.GetPrice().Amount.Uint64(), 10)}, "the package's price is zero")
	}

	// check that if overuse is allowed then the overuse rate is non-zero
	if packageToCheck.GetAllowOveruse() && packageToCheck.GetOveruseRate() == 0 {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_package_fields", map[string]string{"packageAllowOveruse": strconv.FormatBool(packageToCheck.GetAllowOveruse()), "overuseRate": strconv.FormatUint(packageToCheck.GetOveruseRate(), 10)}, "package allows overuse but overuse rate is 0")
	}

	// check compute units fields
	if packageToCheck.GetComputeUnits() == 0 || packageToCheck.GetComputeUnitsPerEpoch() == 0 {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_package_fields", map[string]string{"packageComputeUnits": strconv.FormatUint(packageToCheck.GetComputeUnits(), 10), "packageComputeUnitsPerEpoch": strconv.FormatUint(packageToCheck.GetComputeUnitsPerEpoch(), 10)}, "one or more of the package's compute unit fields is zero")
	}

	// check that the package's servicersToPair is larger than 1
	if packageToCheck.GetServicersToPair() <= 1 {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_package_fields", map[string]string{"packageServicersToPair": strconv.FormatUint(packageToCheck.GetServicersToPair(), 10)}, "package's servicersToPair field is less than or equal to zero")
	}

	// check that the subscriptions field is zero (this field counts the number of users subscribed to this package. If it's a new package, it must be zero)
	if packageToCheck.GetSubscriptions() != 0 {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_package_fields", map[string]string{"packageSubscriptions": strconv.FormatUint(packageToCheck.GetSubscriptions(), 10)}, "package's subscriptions field is not equal to zero")
	}

	// check that the package's name length is below the max length
	if len(packageToCheck.GetName()) > types.MAX_LEN_PACKAGE_NAME {
		return utils.LavaError(ctx, k.Logger(ctx), "verify_package_fields", map[string]string{"packageName": packageToCheck.GetName()}, "package's name is too long")
	}

	return nil
}

// Function to delete a package from storage
func (k Keeper) deletePackage(ctx sdk.Context, packageIndex string, epoch uint64) error {
	// get the package storage of the input package ID
	packageVersionsStorage, found := k.GetPackageVersionsStorage(ctx, packageIndex)
	if !found {
		return utils.LavaError(ctx, k.Logger(ctx), "get_package_versions_storage", map[string]string{"packageIndex": packageIndex}, "could not get packageVersionsStorage with package index")
	}

	// get the fixation entry list
	entryList := packageVersionsStorage.GetPackageStorage().GetEntryList()

	// get the package fixation entry
	packageFixationEntryToDelete, _, found := common.GetFixatedEntry(entryList, epoch)
	if !found {
		return utils.LavaError(ctx, k.Logger(ctx), "get_fixated_entry", map[string]string{"packageIndex": packageIndex, "epoch": strconv.FormatUint(epoch, 10)}, "could not get package fixation entry with package index and epoch")
	}

	// get an edited entryList that doesn't include the packageFixationEntryToDelete
	updatedEntryList, err := common.SetFixatedEntry(ctx, entryList, packageFixationEntryToDelete, false)
	if err != nil {
		return utils.LavaError(ctx, k.Logger(ctx), "set_fixated_entry", map[string]string{"err": err.Error()}, "could not set package storage (delete package)")
	}

	// update the packageVersionsStorage's entry list
	packageVersionsStorage.PackageStorage.EntryList = updatedEntryList

	// update the KVStore with the new packageVersionsStorage
	k.SetPackageVersionsStorage(ctx, packageVersionsStorage)

	return nil
}
