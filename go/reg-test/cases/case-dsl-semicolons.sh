run_mlr --from $indir/s.dkvp --idkvp --opprint put '$x=1; $y=2'
run_mlr --from $indir/s.dkvp --idkvp --opprint put '$x=1; $y=2;'
run_mlr --from $indir/s.dkvp --idkvp --opprint put '; $x=1'
run_mlr --from $indir/s.dkvp --idkvp --opprint put ';;;;;'
run_mlr --from $indir/s.dkvp --idkvp --opprint put 'begin{} $x=1; end{}'
run_mlr --from $indir/s.dkvp --idkvp --opprint put 'begin{}; $x=1; end{}'
run_mlr --from $indir/s.dkvp --idkvp --opprint put 'begin{} $x=1;;; end{}'
run_mlr --from $indir/s.dkvp --idkvp --opprint put ';;;begin{} ;;; end{};;'
