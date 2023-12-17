	// withdraw all delegator rewards
dels := app.StakingKeeper.GetAllDelegations(ctx)
for _, delegation := range dels {
err := app.DistrKeeper.WithdrawDelegationRewards(ctx, delegation.GetDelegatorAddr(), delegation.GetValidatorAddr())
if err != nil {
panic(err)
}
}

	// clear validator slash events
	// app.DistrKeeper.DeleteAllValidatorSlashEvents(ctx)

	// clear validator historical rewards
	// app.DistrKeeper.DeleteAllValidatorHistoricalRewards(ctx)

	// set context height to zero
	height := ctx.BlockHeight()
