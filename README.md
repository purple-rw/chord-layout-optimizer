# chord-layout-optimizer
Layout Optimizer for chording keyboard layout

The purpose is to find better chording assignment for each symbol on a one-handed keyboard layout, similar to [Artsey](https://artsey.io/), [Ardux](https://inkeys.wiki/en/keymaps/ardux), or [Taipo](https://inkeys.wiki/en/keymaps/taipo).

The evolutionary algorithm used here is based on the one mentioned by Dickens at [Algorithm Part 1](https://mathematicalmulticore.wordpress.com/2009/08/07/optimized-evolutionary-algorithm-for-keyboard-design-part-1/) and [Part 2](https://mathematicalmulticore.wordpress.com/2009/08/11/optimized-evolutionary-algorithm-for-keyboard-design-part-2/).

1. Create a pool containing `p` randomly-generated keyboard layouts.
2. Score each keyboard according to a fitness function and sort the keyboards by score.
3. Randomly delete half of the pool (giving preference to keyboards with lower fitness) and create a copy of each remaining keyboard.
4. Mutate the copies by randomly swapping the positions of two keys `m` times.
5. Repeat steps 2-4 until the best keyboard in the pool has not changed for `b` rounds.
6. Place this best keyboard in pool O and sort each keyboard in O by score.
7. Repeat steps 2-6 until the best keyboard in pool O has not changed for `o` rounds.
8. Repeat steps 2-4 using pool O until the best keyboard in the pool has not changed for `q` rounds.

k = number_of_positions
n = nth_best_layout
c = possible_mutations
1/c = (2*n/k) * (1/(k-1))
1 - ((c-1)/c) ^ b >= confidence