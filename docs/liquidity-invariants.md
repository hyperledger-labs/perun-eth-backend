# LiquidityPool — invariants and trust model

Canonical source: **`bindings/contracts/contracts/LiquidityPool.sol`** (the
hardhat project that also generates the Go bindings in
`bindings/liquiditypool/`). Tests: `bindings/contracts/test/LiquidityPool.ts`.
Do not keep design copies of the contract elsewhere — a diverged draft under
`docs/` was removed for exactly that reason; this document records the
*invariants*, the source records the code.

The pool is a **collateralized-operator** lending pool: LPs deposit ETH for
shares; the operator (the hub) borrows free liquidity to fund channels and
must return principal plus a fee at settlement; an operator bond collateralizes
every borrowed wei, and slashing the bond is permissionless once a channel
overstays its settlement window.

## State and definitions

| Term | Definition |
| --- | --- |
| `operatorBond` | Operator collateral, posted with `bondETH()` (operator-only, payable). Lives in the contract balance but is **not** LP money. |
| `totalLockedETH` | Sum of principal currently lent to channels via `fundChannel`. |
| `totalAssets()` | `address(this).balance − operatorBond + totalLockedETH` — LP money, including money currently out on loan. |
| `withdrawableETH()` | `address(this).balance − operatorBond` — free liquidity available for withdrawals and new channel fundings. |
| `channelDeadline[id]` | `block.timestamp + settlementWindow`, set at `fundChannel`. |
| `minFeeBps` | On-chain settlement fee floor, set at construction. |

## Invariants

### I1 — Solvency: ETH backing per share never decreases

Share price is `totalAssets() / totalShares` (with the virtual offset, see I4).
Every state transition preserves or increases it:

- `depositETH` adds `msg.value` to assets and mints shares at the current
  price (rounding pool-favor). `depositFor` is the same transition with the
  same pricing; only the mint recipient differs (see I6).
- `withdrawETH` burns shares and pays out at the current price (rounding
  pool-favor).
- `fundChannel` converts free balance into `totalLockedETH` — assets
  unchanged.
- `settleChannel` returns `principal` (assets unchanged: locked → balance) and
  the fee surplus **increases** assets.
- `expireChannel` slashes `operatorBond −= principal` while clearing
  `totalLockedETH −= principal` — the balance term of `totalAssets()` is
  untouched and the two adjustments cancel: **assets unchanged**. LPs never
  pay for an operator failure.

Test anchors: pricing suite ("redeems at totalAssets while locked"), expiry
suite ("slash leaves totalAssets unchanged"), e2e dust accounting.

### I2 — Coverage: every locked wei is collateralized

`operatorBond ≥ totalLockedETH` at all times:

- `fundChannel` requires `operatorBond ≥ totalLockedETH + amount` (and
  `amount ≤ withdrawableETH()`).
- `withdrawBond` requires post-withdraw `operatorBond − amount ≥
  totalLockedETH`.
- `expireChannel` reduces both sides by the same `principal`.

Consequence: the I1 slash in `expireChannel` can always be paid — the bond
cannot be withdrawn out from under an open channel.

### I3 — Fee floor

`settleChannel` requires `msg.value ≥ principal + principal · minFeeBps /
10000`; `minSettlementValue(channelId)` exposes the exact bound. The operator
cannot settle at a loss to LPs by returning bare principal. (The hub's
configured fee must be ≥ the floor or its settlements revert; devnet: 30 bps
hub vs 10 bps floor.)

### I4 — Pricing symmetry and the virtual offset

Deposit and withdraw both price at `totalAssets()`:

```text
shares  = eth    · (totalShares + 1) / (totalAssets + 1)   // deposit
ethOut  = shares · (totalAssets + 1) / (totalShares + 1)   // withdraw
```

The OZ-style +1 virtual share/asset keeps the first mint 1:1, makes rounding
dust accrue to the pool, and renders the first-depositor inflation attack
unprofitable (regression test: attacker's redemption < donation). There is
**no withdraw lockout**: a withdrawal reverts only when the entitled ETH
exceeds `withdrawableETH()` ("Insufficient free liquidity") — it never pays
out at a degraded price.

### I5 — No untracked value paths

Plain transfers revert (`receive()` reverts); ETH enters only through
`depositETH`, `depositFor`, `bondETH`, and `settleChannel`, each with explicit
accounting.
State-mutating externals are `nonReentrant` (reentrancy regression test:
re-entering `withdrawETH` from a receiver contract fails).

### I6 — Deposit attribution

`deposit` credits `msg.sender`; `depositFor(beneficiary)` credits
`beneficiary` and reverts on the zero address. Both mint at the same price
(I4) — the recipient is the only difference, and the payer keeps no claim on
the deposited ETH.

This backs the cross-chain conversion path: when an ETH→CKB swap drains a CKB
LP cell, the operator deposits the traded countervalue via `depositFor` to the
beneficiary recorded **in that cell's on-chain data**. The beneficiary is
owner-designated (the CKB verifier requires the owner's signature at cell
creation) and immutable for the cell's life, so the operator chooses neither
the recipient nor the ability to redirect it; the `Deposited` event ties each
conversion to its beneficiary and amount on-chain.

The contract itself does not verify the CKB-side binding — it cannot read CKB
state. An operator that deposits to the wrong address is committing a
detectable, auditable error, not an undetectable theft: the mint is public and
the cell's beneficiary is public. Attribution correctness across the two
chains is the operator's responsibility (see the hub's conversion records).

## Trust model

What an LP must trust, and what they need not:

- **Need not trust the operator with principal**: borrowed liquidity is
  covered by the bond (I2), and recovery via `expireChannel` is callable by
  *anyone* after the deadline — no owner, no multisig, no oracle.
- **Must trust the operator for liveness and yield**: a dead or malicious
  operator means no new settlements (no fees) and withdrawals capped by free
  liquidity until channels expire and are slashed. Funds are safe; yield and
  immediate exit are not guaranteed.
- **Must trust the owner for role management**: `Ownable.setOperator` can
  rotate the operator. The owner cannot touch LP assets or the bond directly.
- **Off-chain hub accounting** (reservations, matching, PMM pricing) is
  bookkeeping above the contract; nothing the hub records can move contract
  funds outside the invariants above.

## Residual risks

- **Operator griefing via settlement window**: the operator can let channels
  expire instead of settling. LPs are made whole (I1), but capital sits locked
  for up to `settlementWindow` per channel. Size the window accordingly.
- **Bond sizing is operational**: I2 caps *concurrent* locked principal at the
  bond, but nothing on-chain forces the operator to keep the bond topped up
  after slashes — an underfunded bond degrades to "no new channel fundings".
- **Owner key compromise** rotates the operator to an attacker: they still
  cannot take principal (I2 holds for the new operator), but they can stall
  the pool and burn the *previous* operator's remaining bond via deliberate
  expiries.
- **Marginal-rate quotes** (hub-side): the quote protocol carries no amount;
  see the hub's `docs/liquidity-hub-setup.md` §4 for the guard flags.
