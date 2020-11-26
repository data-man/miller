run_mlr --from $indir/s.dkvp --idkvp --opprint put '$z = [$a,$b,$i,$x,$y][1]'
run_mlr --from $indir/s.dkvp --idkvp --opprint put '$z = [$a,$b,$i,$x,$y][-1]'
run_mlr --from $indir/s.dkvp --idkvp --opprint put '$z = [$a,$b,$i,$x,$y][NR]'

run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$up   = $[NR]'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$down = $[-NR]'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$up   = $*[NR]'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$down = $*[-NR]'

run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[1] = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[2] = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[5] = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[-1] = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[-2] = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[-5] = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '@idx = NR % 5; @idx = @idx == 0 ? 5 : @idx; $[@idx] = "NEW"'

run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[1]       = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[2]       = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[5]       = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[-1]      = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[-2]      = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[-5]      = "new"'

run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[1]]     = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[2]]     = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[5]]     = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[-1]]    = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[-2]]    = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[-5]]    = "new"'

run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[[1]]]   = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[[2]]]   = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[[5]]]   = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[[-1]]]  = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[[-2]]]  = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$[[[-5]]]  = "new"'

run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[1]      = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[2]      = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[5]      = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[-1]     = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[-2]     = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[-5]     = "new"'

run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[1]]    = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[2]]    = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[5]]    = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[-1]]   = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[-2]]   = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[-5]]   = "new"'

run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[[1]]]  = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[[2]]]  = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[[5]]]  = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[[-1]]] = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[[-2]]] = "new"'
run_mlr --opprint --from $indir/s.dkvp --from $indir/t.dkvp put '$*[[[-5]]] = "new"'

